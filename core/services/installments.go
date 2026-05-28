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
type installmentsService struct {
	config     config.Config
	repository ports.InstallmentsRepository
}

// NewURolesService creates a new user service
func NewInstallmentService(cfg config.Config, repo ports.InstallmentsRepository) ports.InstallmentsService {
	return &installmentsService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *installmentsService) Create(ctx context.Context, pagos models.CreateInstallmentReq) (string, error) {

	now := time.Now().UTC()
	pagos.CreatedDate = now
	pagos.UpdatedDate = now
	insertedID, err := p.repository.Create(ctx, models.CreateInstallmentReq(pagos))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *installmentsService) GetInforme(ctx context.Context, filter map[string]interface{}) (*models.InstallmentInfListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.GetInf(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.InstallmentInfListResponse{
			Items:      []models.InstallmentInf{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.InstallmentInfListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetAll users
func (p *installmentsService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.InstallmentListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.InstallmentListResponse{
			Items:      []models.InstallmentResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.InstallmentListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil

}

// GetByID user
func (p *installmentsService) GetByID(ctx context.Context, ID string) (resp models.InstallmentResp, err error) {
	installment, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	if installment == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.InstallmentResp{}, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	resp = *installment.(*models.InstallmentResp)

	return
}

// get by id for update
func (p *installmentsService) GetByIDu(ctx context.Context, ID string) (resp models.UpdateInstallmentReq, err error) {
	installment, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	if installment == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.UpdateInstallmentReq{}, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	resp = *installment.(*models.UpdateInstallmentReq)

	return
}

// Update user
func (p *installmentsService) Update(ctx context.Context, ID string, installment models.UpdateInstallmentReq) (err error) {

	dbInstallment, err := p.GetByIDu(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	//dbPagos.UpdatedDate = time.Now().UTC()

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.UpdateInstallmentReq(dbInstallment))

	return err
}

// Delete user
func (p *installmentsService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
