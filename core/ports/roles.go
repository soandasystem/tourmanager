package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type RolesRepository interface {
	repository.Repository
}

// SaleService interface
type RolesService interface {
	Create(ctx context.Context, roles models.CreateRolesReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.RolesListResponse, error)
	GetByID(ctx context.Context, ID string) (models.RolesResp, error)
	Update(ctx context.Context, ID string, roles models.UpdateRolesReq) error
	Delete(ctx context.Context, ID string) error
}
