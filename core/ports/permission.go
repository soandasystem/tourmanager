package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type PermissionRepository interface {
	repository.Repository
}

// SaleService interface
type PermissionService interface {
	Create(ctx context.Context, permisson models.CreateRolesPermissionsReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.RolesPermissionsListResponse, error)
	GetByID(ctx context.Context, ID string) (models.RolesPermissionsResp, error)
	Update(ctx context.Context, ID string, permission models.UpdateRolesPermissionsReq) error
	Delete(ctx context.Context, ID string, filter map[string]interface{}) error
}
