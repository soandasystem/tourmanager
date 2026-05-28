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

// paymentInstallmentService adapter of an user service
type paymentInstallmentService struct {
	config     config.Config
	repository ports.PaymentInstallmentRepository
}

// NewPaymentInstallmentService creates a new user service
func NewPaymentInstallmentService(cfg config.Config, repo ports.PaymentInstallmentRepository) ports.PaymentInstallmentService {
	return &paymentInstallmentService{
		config:     cfg,
		repository: repo,
	}
}

// Create paymentInstallment
func (p *paymentInstallmentService) Create(ctx context.Context, paymentInstallment models.CreatePaymentInstallmentReq) (string, error) {

	now := time.Now().UTC()
	paymentInstallment.CreatedDate = now
	paymentInstallment.UpdatedDate = now
	insertedID, err := p.repository.Create(ctx, models.CreatePaymentInstallmentReq(paymentInstallment))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *paymentInstallmentService) GetInforme(ctx context.Context, filter map[string]interface{}) (*models.PaymentInstallmentInfListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.GetInf(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.PaymentInstallmentInfListResponse{
			Items:      []models.PaymentInstallmentInf{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.PaymentInstallmentInfListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetAll users
func (p *paymentInstallmentService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.PaymentInstallmentListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.PaymentInstallmentListResponse{
			Items:      []models.PaymentInstallmentResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.PaymentInstallmentListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil

}

// GetByID user
func (p *paymentInstallmentService) GetByID(ctx context.Context, ID string) (resp models.PaymentInstallmentResp, err error) {
	paymentInstallment, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	if paymentInstallment == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.PaymentInstallmentResp{}, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	resp = *paymentInstallment.(*models.PaymentInstallmentResp)

	return
}

// get by id for update
func (p *paymentInstallmentService) GetByIDu(ctx context.Context, ID string) (resp models.UpdatePaymentInstallmentReq, err error) {
	paymentInstallment, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	if paymentInstallment == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.UpdatePaymentInstallmentReq{}, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	resp = *paymentInstallment.(*models.UpdatePaymentInstallmentReq)

	return
}

// Update payment installment
func (p *paymentInstallmentService) Update(ctx context.Context, ID string, paymentInstallment models.UpdatePaymentInstallmentReq) (err error) {

	dbPaymentInstallment, err := p.GetByIDu(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	//dbPagos.UpdatedDate = time.Now().UTC()

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.UpdatePaymentInstallmentReq(dbPaymentInstallment))

	return err
}

// Delete payment installment
func (p *paymentInstallmentService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
