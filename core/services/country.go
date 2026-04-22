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
type countryService struct {
	config     config.Config
	repository ports.CountryRepository
}

// NewURolesService creates a new user service
func NewCountryService(cfg config.Config, repo ports.CountryRepository) ports.CountryService {
	return &countryService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *countryService) Create(ctx context.Context, Country models.CreateCountryReq) (string, error) {

	insertedID, err := p.repository.Create(ctx, models.CreateCountryReq(Country))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *countryService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.CountryListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, map[string]interface{}{}, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.CountryListResponse{
			Items:      []models.CountryResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.CountryListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID user
func (p *countryService) GetByID(ctx context.Context, ID string) (resp models.CountryResp, err error) {
	Country, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("country con ID %s no encontrado", ID)
	}

	if Country == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.CountryResp{}, fmt.Errorf("country con ID %s no encontrado", ID)
	}

	resp = *Country.(*models.CountryResp)

	return
}

// Update user
func (p *countryService) Update(ctx context.Context, ID string, country models.UpdateCountryReq) (err error) {

	dbCountry, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Country(dbCountry))

	return err
}

// Delete user
func (p *countryService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
