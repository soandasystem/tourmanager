package services

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"tourmanager/core/models"
	"tourmanager/util"
)

// FirmarContrato descarga el DOCX temporal desde B2, inserta la firma en el marcador "Firma"
// y utiliza Aspose Cloud para generar el PDF definitivo.
func (s *contratoService) FirmarContrato(ctx context.Context, req models.ContratoFirmaReq) (models.ContratoPDFResp, error) {
	if req.DocxURL == "" {
		return models.ContratoPDFResp{}, fmt.Errorf("docx_url es requerido")
	}

	// 1. Crear un directorio temporal
	tempDir, err := os.MkdirTemp("", fmt.Sprintf("contrato-firma-%s-", req.SessionID))
	if err != nil {
		return models.ContratoPDFResp{}, fmt.Errorf("error creando directorio temporal: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// 2. Decodificar la firma de base64 y guardarla localmente
	firmaBytes, err := decodeBase64Image(req.FirmaBase64)
	if err != nil {
		return models.ContratoPDFResp{}, fmt.Errorf("error decodificando la firma base64: %w", err)
	}
	firmaPath := filepath.Join(tempDir, "firma.png")
	if err := os.WriteFile(firmaPath, firmaBytes, 0644); err != nil {
		return models.ContratoPDFResp{}, fmt.Errorf("error guardando imagen de firma: %w", err)
	}

	// 3. Descargar el archivo DOCX temporal de B2 al disco local
	docxBytes, err := downloadFile(ctx, req.DocxURL)
	if err != nil {
		return models.ContratoPDFResp{}, fmt.Errorf("error descargando DOCX temporal desde B2: %w", err)
	}
	localDocxPath := filepath.Join(tempDir, "contrato_temp.docx")
	if err := os.WriteFile(localDocxPath, docxBytes, 0644); err != nil {
		return models.ContratoPDFResp{}, fmt.Errorf("error guardando DOCX localmente: %w", err)
	}

	// 4. Configurar Aspose Client
	if s.config.AsposeClientID == "" || s.config.AsposeClientSecret == "" {
		return models.ContratoPDFResp{}, fmt.Errorf("las credenciales de Aspose (Client ID / Secret) no están configuradas")
	}
	asposeClient := util.NewAsposeClient(s.config.AsposeClientID, s.config.AsposeClientSecret)

	// Obtener token
	token, err := asposeClient.GetToken(ctx)
	if err != nil {
		return models.ContratoPDFResp{}, fmt.Errorf("error obteniendo token de Aspose: %w", err)
	}

	// Nombres para Aspose
	remoteDocxName := fmt.Sprintf("%s.docx", uuid.New().String())
	remotePdfName := fmt.Sprintf("%s.pdf", uuid.New().String())
	localPdfName := "output.pdf"
	if req.FileNameFirma != "" {
		localPdfName = req.FileNameFirma
		if !strings.HasSuffix(strings.ToLower(localPdfName), ".pdf") {
			localPdfName += ".pdf"
		}
	}
	localPdfPath := filepath.Join(tempDir, localPdfName)

	// 5. Subir DOCX original a Aspose
	if err := asposeClient.UploadFile(ctx, token, localDocxPath, remoteDocxName); err != nil {
		return models.ContratoPDFResp{}, fmt.Errorf("error subiendo archivo a Aspose: %w", err)
	}

	// 6. Insertar imagen en el marcador "Firma" usando Aspose
	if err := asposeClient.InsertImageAtBookmark(ctx, token, remoteDocxName, "Firma", firmaPath); err != nil {
		return models.ContratoPDFResp{}, fmt.Errorf("error insertando firma en el marcador 'Firma': %w", err)
	}

	// 7. Convertir DOCX a PDF y descargarlo
	if err := asposeClient.ConvertToPDF(ctx, token, remoteDocxName, remotePdfName, localPdfPath); err != nil {
		return models.ContratoPDFResp{}, fmt.Errorf("error convirtiendo/descargando PDF con Aspose: %w", err)
	}

	// 8. Subir el PDF resultante a B2
	var finalPDFUrl string
	if s.storage != nil {
		f, err := os.Open(localPdfPath)
		if err != nil {
			return models.ContratoPDFResp{}, fmt.Errorf("error abriendo pdf para subir a B2: %w", err)
		}
		defer f.Close()

		objectKey := fmt.Sprintf("contratos_firmados/%s/%s", req.SessionID, localPdfName)
		url, err := s.storage.Upload(ctx, f, objectKey, "application/pdf")
		if err != nil {
			return models.ContratoPDFResp{}, fmt.Errorf("error subiendo PDF a B2: %w", err)
		}
		finalPDFUrl = url
	} else {
		finalPDFUrl = localPdfPath
	}

	return models.ContratoPDFResp{
		PDFFile: finalPDFUrl,
		Message: "Contrato PDF generado y firmado correctamente con Aspose (Marcador)",
	}, nil
}

// decodeBase64Image remueve el prefijo data URI (si existe) y decodifica la cadena base64.
func decodeBase64Image(firmaBase64 string) ([]byte, error) {
	raw := firmaBase64
	if idx := strings.Index(raw, ","); idx >= 0 {
		raw = raw[idx+1:]
	}

	imgBytes, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		// Intentar con RawStdEncoding (sin padding)
		imgBytes, err = base64.RawStdEncoding.DecodeString(raw)
		if err != nil {
			return nil, fmt.Errorf("base64 inválido: %w", err)
		}
	}
	return imgBytes, nil
}
