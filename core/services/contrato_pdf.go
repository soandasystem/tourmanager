package services

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-pdf/fpdf"

	"tourmanager/core/models"
)

// FirmarContrato lee los datos del contrato guardados en Fase 1,
// genera un PDF completo con todos los campos y la firma incrustada.
func (s *contratoService) FirmarContrato(ctx context.Context, req models.ContratoFirmaReq) (models.ContratoPDFResp, error) {
	// 1. Verificar que el directorio de sesión existe
	tempDir := filepath.Join(os.TempDir(), "contratos", req.SessionID)
	dataPath := filepath.Join(tempDir, "data.json")

	dataBytes, err := os.ReadFile(dataPath)
	if err != nil {
		return models.ContratoPDFResp{}, fmt.Errorf("sesión no encontrada o expirada (%s): %w", req.SessionID, err)
	}

	// 2. Deserializar los datos del contrato
	var contratoData models.ContratoReq
	if err := json.Unmarshal(dataBytes, &contratoData); err != nil {
		return models.ContratoPDFResp{}, fmt.Errorf("error leyendo datos del contrato: %w", err)
	}

	// 3. Decodificar la firma base64 y guardarla como imagen temporal
	firmaPath, err := saveFirmaImage(tempDir, req.FirmaBase64)
	if err != nil {
		return models.ContratoPDFResp{}, fmt.Errorf("error procesando firma: %w", err)
	}
	defer os.Remove(firmaPath)

	// 4. Generar el PDF
	pdfName := "contrato.pdf"
	if req.FileNameFirma != "" {
		pdfName = req.FileNameFirma
		if !strings.HasSuffix(strings.ToLower(pdfName), ".pdf") {
			pdfName += ".pdf"
		}
	}
	pdfPath := filepath.Join(tempDir, pdfName)
	if err := generateContratoPDF(pdfPath, contratoData, firmaPath); err != nil {
		return models.ContratoPDFResp{}, fmt.Errorf("error generando PDF: %w", err)
	}

	var finalPDFUrl string
	if s.storage != nil {
		f, err := os.Open(pdfPath)
		if err != nil {
			return models.ContratoPDFResp{}, fmt.Errorf("error abriendo pdf para subir: %w", err)
		}
		defer f.Close()

		objectKey := fmt.Sprintf("contratos_firmados/%s/%s", req.SessionID, pdfName)
		url, err := s.storage.Upload(ctx, f, objectKey, "application/pdf")
		if err != nil {
			return models.ContratoPDFResp{}, fmt.Errorf("error subiendo PDF a B2: %w", err)
		}
		finalPDFUrl = url
	} else {
		finalPDFUrl = pdfPath
	}

	return models.ContratoPDFResp{
		PDFFile: finalPDFUrl,
		Message: "Contrato PDF generado y firmado correctamente",
	}, nil
}

// saveFirmaImage decodifica el base64 de la firma y lo guarda en disco.
// Soporta PNG y JPEG. Retorna la ruta del archivo temporal.
func saveFirmaImage(dir, firmaBase64 string) (string, error) {
	// Eliminar prefijo data URI si existe (ej: "data:image/png;base64,...")
	raw := firmaBase64
	if idx := indexOf(raw, ","); idx >= 0 {
		raw = raw[idx+1:]
	}

	imgBytes, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		// Intentar con RawStdEncoding (sin padding)
		imgBytes, err = base64.RawStdEncoding.DecodeString(raw)
		if err != nil {
			return "", fmt.Errorf("base64 inválido: %w", err)
		}
	}

	// Detectar el tipo de imagen
	_, format, err := image.DecodeConfig(bytes.NewReader(imgBytes))
	if err != nil {
		return "", fmt.Errorf("formato de imagen no reconocido: %w", err)
	}

	ext := "png"
	if format == "jpeg" {
		ext = "jpg"
	}

	firmaPath := filepath.Join(dir, "firma."+ext)
	if err := os.WriteFile(firmaPath, imgBytes, 0644); err != nil {
		return "", fmt.Errorf("error guardando imagen de firma: %w", err)
	}
	return firmaPath, nil
}

// generateContratoPDF construye el PDF del contrato con todos los campos y la firma.
func generateContratoPDF(pdfPath string, d models.ContratoReq, firmaPath string) error {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(20, 20, 20)
	pdf.AddPage()

	pageW, _ := pdf.GetPageSize()
	contentW := pageW - 40 // márgenes 20+20

	// ── Título principal ──────────────────────────────────────────
	pdf.SetFont("Arial", "B", 16)
	pdf.SetTextColor(31, 60, 120)
	pdf.CellFormat(contentW, 10, "CONTRATO DE PRESTACIÓN DE SERVICIOS EDUCATIVOS", "", 1, "C", false, 0, "")
	pdf.Ln(4)

	// Fecha de emisión
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(60, 60, 60)
	fecha := fmt.Sprintf("En %s, a %s de %s de %s", d.EDireccion, d.VtaDia, d.VtaMes, d.VtaAgno)
	pdf.CellFormat(contentW, 7, fecha, "", 1, "L", false, 0, "")
	pdf.Ln(4)

	// ── Separador ─────────────────────────────────────────────────
	separador(pdf, contentW)

	// ── Datos de la Empresa ───────────────────────────────────────
	seccion(pdf, "DATOS DE LA EMPRESA")
	fila2col(pdf, contentW, "RUT:", d.Rute, "Razón Social:", d.RSocial)
	fila2col(pdf, contentW, "Nombre Fantasía:", d.NFantasia, "Dirección:", d.EDireccion)
	fila2col(pdf, contentW, "RUT Representante Legal:", d.RLegal, "Nombre Representante:", d.NLegal)
	pdf.Ln(3)

	// ── Datos del Colegio ─────────────────────────────────────────
	seccion(pdf, "DATOS DEL ESTABLECIMIENTO")
	fila2col(pdf, contentW, "Colegio:", d.Colegio, "Comuna:", d.Comuna)
	fila2col(pdf, contentW, "Curso/ID:", d.IdCurso, "Programa:", d.Programa)
	pdf.Ln(3)

	// ── Datos del Apoderado y Alumno ──────────────────────────────
	seccion(pdf, "DATOS DEL APODERADO Y ALUMNO")
	fila2col(pdf, contentW, "Nombre Apoderado:", d.NombreApod, "RUT Apoderado:", d.RutApod)
	fila2col(pdf, contentW, "Correo:", d.CorreoApod, "Teléfono:", d.FonoApod)
	fila1col(pdf, contentW, "Nombre Alumno:", d.NombreAlumno)
	pdf.Ln(3)

	// ── Detalles del Programa ─────────────────────────────────────
	seccion(pdf, "DETALLES DEL PROGRAMA")
	fila2col(pdf, contentW, "Reserva:", d.Reserva, "Tipo de Venta:", d.TypeSale)
	fila2col(pdf, contentW, "Valor Programa:", "$"+d.VPrograma, "Tipo de Cambio:", d.Tc)
	fila2col(pdf, contentW, "Liberados:", d.Liberados, "Fecha Pago:", d.FPago)
	pdf.Ln(3)

	// ── Fechas de Salida ──────────────────────────────────────────
	seccion(pdf, "FECHA DE SALIDA")
	salida := fmt.Sprintf("%s de %s de %s (día %s)", d.FSalidaDia, d.FSalidaMes, d.FSalidaAgno, d.FSalidaDia)
	fila1col(pdf, contentW, "Fecha de Salida:", salida)
	pdf.Ln(3)

	// ── Observaciones ─────────────────────────────────────────────
	if d.Observacion != "" {
		seccion(pdf, "OBSERVACIONES")
		pdf.SetFont("Arial", "", 10)
		pdf.SetTextColor(60, 60, 60)
		pdf.MultiCell(contentW, 6, d.Observacion, "", "L", false)
		pdf.Ln(3)
	}

	// ── Firma ─────────────────────────────────────────────────────
	separador(pdf, contentW)
	pdf.Ln(6)
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(31, 60, 120)
	pdf.CellFormat(contentW, 7, "FIRMA DEL APODERADO", "", 1, "C", false, 0, "")
	pdf.Ln(4)

	// Insertar imagen de firma centrada
	firmaW := 60.0
	firmaH := 25.0
	firmaX := (pageW - firmaW) / 2
	pdf.ImageOptions(firmaPath, firmaX, pdf.GetY(), firmaW, firmaH, false, fpdf.ImageOptions{}, 0, "")
	pdf.Ln(firmaH + 4)

	// Línea y nombre bajo la firma
	lineaX := (pageW - 80) / 2
	pdf.Line(lineaX, pdf.GetY(), lineaX+80, pdf.GetY())
	pdf.Ln(2)
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(80, 80, 80)
	pdf.CellFormat(contentW, 5, d.NombreApod+" — RUT: "+d.RutApod, "", 1, "C", false, 0, "")

	// ── Guardar PDF ───────────────────────────────────────────────
	return pdf.OutputFileAndClose(pdfPath)
}

// ── Helpers de layout ────────────────────────────────────────────

func separador(pdf *fpdf.Fpdf, w float64) {
	pdf.SetDrawColor(31, 60, 120)
	pdf.SetLineWidth(0.5)
	x, y := pdf.GetXY()
	pdf.Line(x, y, x+w, y)
	pdf.Ln(3)
}

func seccion(pdf *fpdf.Fpdf, titulo string) {
	pdf.SetFont("Arial", "B", 11)
	pdf.SetFillColor(220, 230, 245)
	pdf.SetTextColor(31, 60, 120)
	pageW, _ := pdf.GetPageSize()
	w := pageW - 40
	pdf.CellFormat(w, 7, "  "+titulo, "", 1, "L", true, 0, "")
	pdf.Ln(2)
}

func fila2col(pdf *fpdf.Fpdf, w float64, label1, val1, label2, val2 string) {
	col := w / 2
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(50, 50, 50)
	pdf.CellFormat(col*0.38, 6, label1, "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 9)
	pdf.CellFormat(col*0.62, 6, val1, "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(col*0.38, 6, label2, "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 9)
	pdf.CellFormat(col*0.62, 6, val2, "", 1, "L", false, 0, "")
}

func fila1col(pdf *fpdf.Fpdf, w float64, label, val string) {
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(50, 50, 50)
	pdf.CellFormat(w*0.25, 6, label, "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 9)
	pdf.CellFormat(w*0.75, 6, val, "", 1, "L", false, 0, "")
}

func indexOf(s, sep string) int {
	for i := 0; i < len(s)-len(sep)+1; i++ {
		if s[i:i+len(sep)] == sep {
			return i
		}
	}
	return -1
}
