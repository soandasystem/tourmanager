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
type paymentService struct {
	config     config.Config
	repository ports.PaymentRepository
}

// NewPagosService creates a new payment service
func NewPaymentService(cfg config.Config, repo ports.PaymentRepository) ports.PaymentService {
	return &paymentService{
		config:     cfg,
		repository: repo,
	}
}

// Create payment
func (p *paymentService) Create(ctx context.Context, payment models.CreatePaymentReq) (string, error) {

	now := time.Now().UTC()
	payment.CreatedDate = now
	payment.UpdatedDate = now
	insertedID, err := p.repository.Create(ctx, models.CreatePaymentReq(payment))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll payments
func (p *paymentService) GetInforme(ctx context.Context, filter map[string]interface{}) (*models.PaymentInfListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.GetInf(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.PaymentInfListResponse{
			Items:      []models.PaymentInf{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.PaymentInfListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetAll payments
func (p *paymentService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.PaymentListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.PaymentListResponse{
			Items:      []models.PaymentResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.PaymentListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil

}

// GetByID payments
func (p *paymentService) GetByID(ctx context.Context, ID string) (resp models.PaymentResp, err error) {
	payment, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	if payment == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.PaymentResp{}, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	resp = *payment.(*models.PaymentResp)

	return
}

// get by id for update
func (p *paymentService) GetByIDu(ctx context.Context, ID string) (resp models.UpdatePaymentReq, err error) {
	payment, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	if payment == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.UpdatePaymentReq{}, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	resp = *payment.(*models.UpdatePaymentReq)

	return
}

// Update payment
func (p *paymentService) Update(ctx context.Context, ID string, payment models.UpdatePaymentReq) (err error) {

	dbPayment, err := p.GetByIDu(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	//dbPagos.UpdatedDate = time.Now().UTC()

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.UpdatePaymentReq(dbPayment))

	return err
}

// Delete payment
func (p *paymentService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
