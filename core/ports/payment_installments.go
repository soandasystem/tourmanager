package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type PaymentInstallmentRepository interface {
	repository.Repository
}

// RolesService interface
type PaymentInstallmentService interface {
	Create(ctx context.Context, pagos models.CreatePaymentInstallmentReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.PaymentInstallmentListResponse, error)
	GetInforme(ctx context.Context, filter map[string]interface{}) (*models.PaymentInstallmentInfListResponse, error)
	GetByID(ctx context.Context, ID string) (models.PaymentInstallmentResp, error)
	Update(ctx context.Context, ID string, pagos models.UpdatePaymentInstallmentReq) error
	Delete(ctx context.Context, ID string) error
}
