package ports

import (
	"context"

	"tourmanager/core/models"
)

// ContratoService define las operaciones disponibles para generación de contratos
type ContratoService interface {
	// GenerarContrato descarga el template DOCX de B2, reemplaza placeholders y guarda un archivo temporal.
	// Devuelve un ContratoTempResp con el session_id y la ruta del DOCX temporal.
	GenerarContrato(ctx context.Context, req models.ContratoReq) (models.ContratoTempResp, error)

	// FirmarContrato recibe el session_id y la firma en base64, genera el PDF definitivo
	// con todos los datos del contrato y la firma incrustada.
	FirmarContrato(ctx context.Context, req models.ContratoFirmaReq) (models.ContratoPDFResp, error)
}
