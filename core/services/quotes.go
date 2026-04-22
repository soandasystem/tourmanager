package services

import (
	"context"
	"errors"
	"fmt"

	"tourmanager/config"
	"tourmanager/core/models"
	"tourmanager/core/ports"

	"github.com/antoniomarfa/hexatools/wrappers"
)

// rolesService adapter of an user service
type QuotesService struct {
	config     config.Config
	repository ports.QuotesRepository
}

// NewURolesService creates a new user service
func NewQuotesService(cfg config.Config, repo ports.QuotesRepository) ports.QuotesService {
	return &QuotesService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *QuotesService) Create(ctx context.Context, quote models.CreateQuoteReq) (string, error) {

	insertedID, err := p.repository.Create(ctx, models.CreateQuoteReq(quote))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *QuotesService) GetInforme(ctx context.Context, filter map[string]interface{}) (*models.QuoteInfListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.GetInf(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.QuoteInfListResponse{
			Items:      []models.QuoteInforme{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.QuoteInfListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetAll users
func (p *QuotesService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.QuoteListResponse, error) {
	// Obtiene las vemtas desde el repositorio

	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.QuoteListResponse{
			Items:      []models.QuoteResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.QuoteListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID user
func (p *QuotesService) GetByID(ctx context.Context, ID string) (resp models.QuoteResp, err error) {
	quotes, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, nil
	}

	if quotes == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.QuoteResp{}, fmt.Errorf("venta con ID %s no encontrado", ID)
	}

	//	resp = models.ColegiosResp(*colegios.(*models.ColegiosResp))
	resp = *quotes.(*models.QuoteResp)
	return resp, nil
	// return
}

// Update user
func (p *QuotesService) Update(ctx context.Context, ID string, quotes models.UpdateQuoteReq) (err error) {
	dbQuotes, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}

	if quotes.Fecha != nil {
		dbQuotes.Fecha = *quotes.Fecha
	}

	if quotes.Estado != nil {
		dbQuotes.Estado = *quotes.Estado
	}

	if quotes.SaleId != nil {
		dbQuotes.SaleId = *quotes.SaleId
	}

	err = p.repository.Update(ctx, ID, models.QuoteResp(dbQuotes))

	return err

}

// Delete user
func (p *QuotesService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
