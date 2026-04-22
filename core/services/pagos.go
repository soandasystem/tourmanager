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
type pagosService struct {
	config     config.Config
	repository ports.PagosRepository
}

// NewURolesService creates a new user service
func NewPagosService(cfg config.Config, repo ports.PagosRepository) ports.PagosService {
	return &pagosService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *pagosService) Create(ctx context.Context, pagos models.CreatePagosReq) (string, error) {

	now := time.Now().UTC()
	pagos.CreatedDate = now
	pagos.UpdatedDate = now
	insertedID, err := p.repository.Create(ctx, models.CreatePagosReq(pagos))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *pagosService) GetInforme(ctx context.Context, filter map[string]interface{}) (*models.PagosInfListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.GetInf(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.PagosInfListResponse{
			Items:      []models.PagosInf{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.PagosInfListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetAll users
func (p *pagosService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.PagosListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.PagosListResponse{
			Items:      []models.PagosResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.PagosListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil

}

// GetByID user
func (p *pagosService) GetByID(ctx context.Context, ID string) (resp models.PagosResp, err error) {
	pagos, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	if pagos == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.PagosResp{}, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	resp = *pagos.(*models.PagosResp)

	return
}

// get by id for update
func (p *pagosService) GetByIDu(ctx context.Context, ID string) (resp models.UpdatePagosReq, err error) {
	pagos, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	if pagos == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.UpdatePagosReq{}, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	resp = *pagos.(*models.UpdatePagosReq)

	return
}

// Update user
func (p *pagosService) Update(ctx context.Context, ID string, pagos models.UpdatePagosReq) (err error) {

	dbPagos, err := p.GetByIDu(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	//dbPagos.UpdatedDate = time.Now().UTC()

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.UpdatePagosReq(dbPagos))

	return err
}

// Delete user
func (p *pagosService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
