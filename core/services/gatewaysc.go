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
type gatewayscService struct {
	config     config.Config
	repository ports.GatewayscRepository
}

// NewURolesService creates a new user service
func NewGatewayscService(cfg config.Config, repo ports.GatewayscRepository) ports.GatewayscService {
	return &gatewayscService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *gatewayscService) Create(ctx context.Context, gatewaysc models.CreateGatewayscReq) (string, error) {

	insertedID, err := p.repository.Create(ctx, models.CreateGatewayscReq(gatewaysc))
	if err != nil {
		return "", nil
	}

	return insertedID, err
}

// GetAll users
func (p *gatewayscService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.GatewayscListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	//	result, err := p.repository.Get(ctx, map[string]interface{}{}, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados a la estructura de respuesta

	if len(result) == 0 {
		return &models.GatewayscListResponse{
			Items:      []models.GatewayscResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.GatewayscListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID user
func (p *gatewayscService) GetByID(ctx context.Context, ID string) (resp models.GatewayscResp, err error) {
	gatewaysc, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("colegio con ID %s no encontrado", ID)
	}

	if gatewaysc == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.GatewayscResp{}, fmt.Errorf("colegio con ID %s no encontrado", ID)
	}

	// Type assertion seguro
	gwResp, ok := gatewaysc.(*models.GatewayscResp)
	if !ok {
		return models.GatewayscResp{}, fmt.Errorf("error de tipo: se esperaba *models.GatewayscResp pero se obtuvo otro tipo")
	}

	resp = *gwResp
	return resp, nil
}

func (p *gatewayscService) GetUpdateByID(ctx context.Context, ID string) (resp models.GatewayscResp, err error) {
	gatewaysc, err := p.repository.GetByIDUpdate(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("pasarela con ID %s no encontrado", ID)
	}

	if gatewaysc == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.GatewayscResp{}, fmt.Errorf("pasarela con ID %s no encontrado", ID)
	}

	// Type assertion seguro
	gwResp, ok := gatewaysc.(*models.GatewayscResp)
	if !ok {
		return models.GatewayscResp{}, fmt.Errorf("error de tipo: se esperaba *models.GatewayscResp pero se obtuvo otro tipo")
	}

	resp = *gwResp
	return resp, nil
}

// Update user
func (p *gatewayscService) Update(ctx context.Context, ID string, gatewaysc models.UpdateGatewayscReq) (err error) {

	dbGatewaysc, err := p.GetUpdateByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	now := time.Now().UTC()
	if gatewaysc.GatewayType != nil {
		dbGatewaysc.GatewayType = *gatewaysc.GatewayType
	}
	if gatewaysc.GatewayDescription != nil {
		dbGatewaysc.GatewayDescription = *gatewaysc.GatewayDescription
	}
	if gatewaysc.GatewayImage != nil {
		dbGatewaysc.GatewayImage = *gatewaysc.GatewayImage
	}
	if gatewaysc.GatewayUrl != nil {
		dbGatewaysc.GatewayUrl = *gatewaysc.GatewayUrl
	}

	dbGatewaysc.UpdatedDate = now
	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Gatewaysc(dbGatewaysc))

	return err
}

// Delete user
func (p *gatewayscService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
