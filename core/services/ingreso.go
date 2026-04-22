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

// ingresoService adapter of an user service
type ingresoService struct {
	config     config.Config
	repository ports.IngresoRepository
}

// NewURolesService creates a new user service
func NewIngresoService(cfg config.Config, repo ports.IngresoRepository) ports.IngresoService {
	return &ingresoService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *ingresoService) Create(ctx context.Context, ingreso models.CreateIngresoReq) (string, error) {

	now := time.Now().UTC()
	ingreso.CreatedDate = now
	ingreso.UpdatedDate = now
	insertedID, err := p.repository.Create(ctx, models.CreateIngresoReq(ingreso))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *ingresoService) GetInforme(ctx context.Context, filter map[string]interface{}) (*models.IngresoInfListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.GetInf(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.IngresoInfListResponse{
			Items:      []models.IngresoInf{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.IngresoInfListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetAll users
func (p *ingresoService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.IngresoListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.IngresoListResponse{
			Items:      []models.IngresoResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.IngresoListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID user
func (p *ingresoService) GetByID(ctx context.Context, ID string) (resp models.IngresoResp, err error) {
	ingreso, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("ingreso con ID %s no encontrado", ID)
	}

	if ingreso == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.IngresoResp{}, fmt.Errorf("ingreso con ID %s no encontrado", ID)
	}

	resp = *ingreso.(*models.IngresoResp)

	return
}

// Update user
func (p *ingresoService) Update(ctx context.Context, ID string, ingreso models.UpdateIngresoReq) (err error) {

	dbIngreso, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	dbIngreso.UpdatedDate = time.Now().UTC()

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Ingreso(dbIngreso))

	return err
}

// Delete user
func (p *ingresoService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
