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
type gatewaysService struct {
	config     config.Config
	repository ports.GatewaysRepository
}

// NewURolesService creates a new user service
func NewGatewaysService(cfg config.Config, repo ports.GatewaysRepository) ports.GatewaysService {
	return &gatewaysService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *gatewaysService) Create(ctx context.Context, gateways models.CreateGatewaysReq) (string, error) {

	insertedID, err := p.repository.Create(ctx, models.CreateGatewaysReq(gateways))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *gatewaysService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.GatewaysListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}
	// Convierte los resultados
	if len(result) == 0 {
		return &models.GatewaysListResponse{
			Items:      []models.GatewaysResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.GatewaysListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID user
func (p *gatewaysService) GetByID(ctx context.Context, ID string) (resp models.GatewaysResp, err error) {
	gateways, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("colegio con ID %s no encontrado", ID)
	}

	if gateways == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.GatewaysResp{}, fmt.Errorf("colegio con ID %s no encontrado", ID)
	}

	resp = *gateways.(*models.GatewaysResp)

	return
}

func (p *gatewaysService) GetUpdateByID(ctx context.Context, ID string) (resp models.Gateways, err error) {
	gateways, err := p.repository.GetByIDUpdate(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("pasarela con ID %s no encontrado", ID)
	}

	if gateways == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.Gateways{}, fmt.Errorf("pasarela con ID %s no encontrado", ID)
	}

	resp = *gateways.(*models.Gateways)

	return
}

// Update user
func (p *gatewaysService) Update(ctx context.Context, ID string, Gateways models.Gateways) (err error) {

	return err
}

// Delete user
func (p *gatewaysService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
