package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"tourmanager/config"
	"tourmanager/core/models"
	"tourmanager/core/ports"

	"github.com/antoniomarfa/hexatools/wrappers"
)

// rolesService adapter of an user service
type regionService struct {
	config     config.Config
	repository ports.RegionRepository
}

// NewURolesService creates a new user service
func NewRegionService(cfg config.Config, repo ports.RegionRepository) ports.RegionService {
	return &regionService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *regionService) Create(ctx context.Context, region models.CreateRegionReq) (string, error) {

	insertedID, err := p.repository.Create(ctx, models.CreateRegionReq(region))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *regionService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.RegionListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, map[string]interface{}{}, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.RegionListResponse{
			Items:      []models.RegionResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.RegionListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID user
func (p *regionService) GetByID(ctx context.Context, ID string) (resp models.RegionResp, err error) {
	region, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("region con ID %s no encontrado", ID)
	}

	if region == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.RegionResp{}, fmt.Errorf("region con ID %s no encontrado", ID)
	}

	resp = *region.(*models.RegionResp)

	return
}

// Update user
func (p *regionService) Update(ctx context.Context, ID string, region models.UpdateRegionReq) (err error) {

	dbRegion, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	now := time.Now().UTC()
	if region.Description != nil {
		dbRegion.Description = *region.Description
	}
	dbRegion.UpdatedDate = now
	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Region(dbRegion))

	return err
}

// Delete user
func (p *regionService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
