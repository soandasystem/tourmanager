package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type PaymentRepository interface {
	repository.Repository
}

// RolesService interface
type PaymentService interface {
	Create(ctx context.Context, pagos models.CreatePaymentReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.PaymentListResponse, error)
	GetInforme(ctx context.Context, filter map[string]interface{}) (*models.PaymentInfListResponse, error)
	GetByID(ctx context.Context, ID string) (models.PaymentResp, error)
	Update(ctx context.Context, ID string, pagos models.UpdatePaymentReq) error
	Delete(ctx context.Context, ID string) error
}
