package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type InstallmentsRepository interface {
	repository.Repository
}

// RolesService interface
type InstallmentsService interface {
	Create(ctx context.Context, pagos models.CreateInstallmentReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.InstallmentListResponse, error)
	GetInforme(ctx context.Context, filter map[string]interface{}) (*models.InstallmentInfListResponse, error)
	GetByID(ctx context.Context, ID string) (models.InstallmentResp, error)
	Update(ctx context.Context, ID string, pagos models.UpdateInstallmentReq) error
	Delete(ctx context.Context, ID string) error
}
