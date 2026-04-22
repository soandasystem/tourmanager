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
type voucherService struct {
	config     config.Config
	repository ports.VoucherRepository
}

// NewURolesService creates a new user service
func NewVoucherService(cfg config.Config, repo ports.VoucherRepository) ports.VoucherService {
	return &voucherService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *voucherService) Create(ctx context.Context, voucher models.CreateVoucherReq) (string, error) {

	insertedID, err := p.repository.Create(ctx, models.CreateVoucherReq(voucher))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *voucherService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.VoucherListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.VoucherListResponse{
			Items:      []models.VoucherResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.VoucherListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID user
func (p *voucherService) GetByID(ctx context.Context, ID string) (resp models.VoucherResp, err error) {
	voucher, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("voucher con ID %s no encontrado", ID)
	}

	if voucher == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.VoucherResp{}, fmt.Errorf("voucher con ID %s no encontrado", ID)
	}

	resp = *voucher.(*models.VoucherResp)

	return
}

// Update user
func (p *voucherService) Update(ctx context.Context, ID string, voucher models.UpdateVoucherReq) (err error) {

	dbVoucher, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	dbVoucher.Used = voucher.Used

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.VoucherResp(dbVoucher))

	return err
}

// Delete user
func (p *voucherService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
