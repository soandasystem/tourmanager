package services

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"

	"tourmanager/config"
	"tourmanager/core/models"
	"tourmanager/core/ports"
)

type contratoService struct {
	config  config.Config
	storage ports.UploadStorage
}

// NewContratoService crea una nueva instancia del servicio de contratos
func NewContratoService(cfg config.Config, storage ports.UploadStorage) ports.ContratoService {
	return &contratoService{config: cfg, storage: storage}
}

// GenerarContrato descarga el template DOCX, reemplaza los placeholders {{campo}}
// con los valores del request y guarda el resultado en un directorio temporal.
func (s *contratoService) GenerarContrato(ctx context.Context, req models.ContratoReq) (models.ContratoTempResp, error) {
	if req.TemplateFilename == "" {
		return models.ContratoTempResp{}, fmt.Errorf("el nombre del archivo template (template_filename) es requerido")
	}

	// Construir la URL completa del template usando la URL pública de descarga
	baseURL := strings.TrimRight(s.config.B2PublicURL, "/")
	templateURL := fmt.Sprintf("%s/%s", baseURL, req.TemplateFilename)

	fmt.Println("URL template:", templateURL)

	// 1. Descargar el template DOCX desde B2
	docxBytes, err := downloadFile(ctx, templateURL)
	if err != nil {
		return models.ContratoTempResp{}, fmt.Errorf("error descargando template: %w", err)
	}

	// 2. Construir el mapa de reemplazos
	replacements := buildReplacements(req)

	// 3. Procesar el DOCX
	processedBytes, err := processDocx(docxBytes, replacements)
	if err != nil {
		return models.ContratoTempResp{}, fmt.Errorf("error procesando DOCX: %w", err)
	}

	// 4. Crear directorio temporal
	sessionID := uuid.New().String()

	tempDir := filepath.Join(
		os.TempDir(),
		"contratos",
		sessionID,
	)

	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return models.ContratoTempResp{}, fmt.Errorf("error creando directorio temporal: %w", err)
	}

	// 5. Guardar DOCX con UUID en el nombre para evitar colisiones
	docxName := fmt.Sprintf("contrato-%s.docx", uuid.New().String())
	docxPath := filepath.Join(tempDir, docxName)

	if err := os.WriteFile(docxPath, processedBytes, 0644); err != nil {
		return models.ContratoTempResp{}, fmt.Errorf("error guardando DOCX temporal: %w", err)
	}

	// 6. Subir el DOCX y el data.json a B2 en carpeta temp/
	var docxURL string
	if s.storage != nil {
		// 6a. Subir el DOCX
		f, err := os.Open(docxPath)
		if err != nil {
			return models.ContratoTempResp{}, fmt.Errorf("error abriendo DOCX para subir: %w", err)
		}
		defer f.Close()

		objectKey := fmt.Sprintf("temp/%s/%s", sessionID, docxName)
		docxURL, err = s.storage.Upload(ctx, f, objectKey, "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
		if err != nil {
			return models.ContratoTempResp{}, fmt.Errorf("error subiendo DOCX a B2: %w", err)
		}

		// 6b. Subir data.json (para que Fase 2 no dependa del disco local)
		dataBytes, err := encodeJSON(req)
		if err != nil {
			return models.ContratoTempResp{}, fmt.Errorf("error serializando datos del contrato: %w", err)
		}
		dataKey := fmt.Sprintf("temp/%s/data.json", sessionID)
		_, err = s.storage.Upload(ctx, bytes.NewReader(dataBytes), dataKey, "application/json")
		if err != nil {
			return models.ContratoTempResp{}, fmt.Errorf("error subiendo data.json a B2: %w", err)
		}
	}

	return models.ContratoTempResp{
		SessionID: sessionID,
		DocxURL:   docxURL,
		Message:   "Contrato temporal generado correctamente",
	}, nil
}

/*
func (s *contratoService) GenerarContrato(ctx context.Context, req models.ContratoReq) (models.ContratoTempResp, error) {
	if req.TemplateFilename == "" {
		return models.ContratoTempResp{}, fmt.Errorf("el nombre del archivo template (template_filename) es requerido")
	}

	// Construir la URL completa (garantizar que no haya doble slash si B2Endpoint termina en /)
	baseURL := strings.TrimRight(s.config.B2Endpoint, "/")
	templateURL := fmt.Sprintf("%s/%s", baseURL, req.TemplateFilename)
	fmt.Println("ulr template ", templateURL)
	// 1. Descargar el template DOCX desde B2
	docxBytes, err := downloadFile(ctx, templateURL)
	if err != nil {
		return models.ContratoTempResp{}, fmt.Errorf("error descargando template: %w", err)
	}

	// 2. Construir el mapa de reemplazos {{campo}} → valor
	replacements := buildReplacements(req)

	// 3. Procesar el DOCX (ZIP) reemplazando los placeholders en el XML interno
	processedBytes, err := processDocx(docxBytes, replacements)
	if err != nil {
		return models.ContratoTempResp{}, fmt.Errorf("error procesando DOCX: %w", err)
	}

	// 4. Crear directorio temporal para esta sesión
	sessionID := uuid.New().String()
	tempDir := filepath.Join(os.TempDir(), "contratos", sessionID)
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return models.ContratoTempResp{}, fmt.Errorf("error creando directorio temporal: %w", err)
	}

	// 5. Guardar el DOCX procesado
	docxPath := filepath.Join(tempDir, "contrato.docx")
	if err := os.WriteFile(docxPath, processedBytes, 0644); err != nil {
		return models.ContratoTempResp{}, fmt.Errorf("error guardando DOCX temporal: %w", err)
	}

	// 6. Guardar los datos del contrato en JSON para usar en Fase 2
	dataPath := filepath.Join(tempDir, "data.json")
	dataBytes, err := json.Marshal(req)
	if err != nil {
		return models.ContratoTempResp{}, fmt.Errorf("error serializando datos del contrato: %w", err)
	}
	if err := os.WriteFile(dataPath, dataBytes, 0644); err != nil {
		return models.ContratoTempResp{}, fmt.Errorf("error guardando datos del contrato: %w", err)
	}

	return models.ContratoTempResp{
		SessionID: sessionID,
		TempFile:  docxPath,
		Message:   "Contrato temporal generado correctamente",
	}, nil
}
*/

// downloadFile realiza un HTTP GET y retorna el contenido como bytes
func downloadFile(ctx context.Context, url string) ([]byte, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("respuesta inesperada del servidor: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

// buildReplacements construye el mapa de {{campo}} → valor a partir del ContratoReq
func buildReplacements(req models.ContratoReq) map[string]string {
	return map[string]string{
		"{{vtaDia}}":       req.VtaDia,
		"{{vtaMes}}":       req.VtaMes,
		"{{vtaAgno}}":      req.VtaAgno,
		"{{rute}}":         req.Rute,
		"{{rsocial}}":      req.RSocial,
		"{{nfantasia}}":    req.NFantasia,
		"{{rlegal}}":       req.RLegal,
		"{{nlegal}}":       req.NLegal,
		"{{edireccion}}":   req.EDireccion,
		"{{colegio}}":      req.Colegio,
		"{{comuna}}":       req.Comuna,
		"{{idcurso}}":      req.IdCurso,
		"{{programa}}":     req.Programa,
		"{{reserva}}":      req.Reserva,
		"{{nombreapod}}":   req.NombreApod,
		"{{nombrealumno}}": req.NombreAlumno,
		"{{rutapod}}":      req.RutApod,
		"{{correoapod}}":   req.CorreoApod,
		"{{fonoapod}}":     req.FonoApod,
		"{{observacion}}":  req.Observacion,
		"{{vprograma}}":    req.VPrograma,
		"{{tc}}":           req.Tc,
		"{{liberados}}":    req.Liberados,
		"{{fsalida}}":      req.FSalida,
		"{{fsalidames}}":   req.FSalidaMes,
		"{{fsalidaaño}}":   req.FSalidaAgno,
		"{{fsalidadia}}":   req.FSalidaDia,
		"{{fpago}}":        req.FPago,
		"{{type_sale}}":    req.TypeSale,
	}
}

// processDocx lee el DOCX como ZIP, reemplaza placeholders en los archivos XML
// relevantes (document.xml, header*.xml, footer*.xml) y retorna el DOCX modificado.
func processDocx(docxBytes []byte, replacements map[string]string) ([]byte, error) {
	// Abrir el ZIP en memoria
	reader, err := zip.NewReader(bytes.NewReader(docxBytes), int64(len(docxBytes)))
	if err != nil {
		return nil, fmt.Errorf("error abriendo DOCX como ZIP: %w", err)
	}

	// Buffer de salida para el nuevo ZIP
	var outBuf bytes.Buffer
	writer := zip.NewWriter(&outBuf)

	for _, f := range reader.File {
		rc, err := f.Open()
		if err != nil {
			return nil, fmt.Errorf("error leyendo archivo ZIP %s: %w", f.Name, err)
		}

		content, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return nil, fmt.Errorf("error leyendo contenido de %s: %w", f.Name, err)
		}

		// Aplicar reemplazos solo en los archivos XML del contenido del documento
		if isContentFile(f.Name) {
			content = applyReplacements(content, replacements)
		}

		// Escribir el archivo (modificado o no) al nuevo ZIP
		w, err := writer.CreateHeader(&f.FileHeader)
		if err != nil {
			return nil, fmt.Errorf("error creando entrada ZIP %s: %w", f.Name, err)
		}
		if _, err := w.Write(content); err != nil {
			return nil, fmt.Errorf("error escribiendo contenido ZIP %s: %w", f.Name, err)
		}
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("error cerrando ZIP: %w", err)
	}

	return outBuf.Bytes(), nil
}

// isContentFile retorna true si el archivo XML del DOCX contiene texto editable
func isContentFile(name string) bool {
	name = strings.ToLower(name)
	return strings.Contains(name, "word/document") ||
		strings.Contains(name, "word/header") ||
		strings.Contains(name, "word/footer") ||
		strings.Contains(name, "word/endnotes") ||
		strings.Contains(name, "word/footnotes")
}

// applyReplacements reemplaza cada placeholder en el contenido XML
func applyReplacements(content []byte, replacements map[string]string) []byte {
	text := string(content)
	for placeholder, value := range replacements {
		text = strings.ReplaceAll(text, placeholder, value)
	}
	return []byte(text)
}

// encodeJSON serializa el ContratoReq a JSON (helper para subir data.json a B2)
func encodeJSON(v any) ([]byte, error) {
	return json.Marshal(v)
}
