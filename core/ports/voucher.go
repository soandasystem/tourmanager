package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type VoucherRepository interface {
	repository.Repository
}

// VoucherService interface
type VoucherService interface {
	Create(ctx context.Context, voucher models.CreateVoucherReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.VoucherListResponse, error)
	GetByID(ctx context.Context, ID string) (models.VoucherResp, error)
	Update(ctx context.Context, ID string, voucher models.UpdateVoucherReq) error
	Delete(ctx context.Context, ID string) error
}
