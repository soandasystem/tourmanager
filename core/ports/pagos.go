package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type PagosRepository interface {
	repository.Repository
}

// RolesService interface
type PagosService interface {
	Create(ctx context.Context, pagos models.CreatePagosReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.PagosListResponse, error)
	GetInforme(ctx context.Context, filter map[string]interface{}) (*models.PagosInfListResponse, error)
	GetByID(ctx context.Context, ID string) (models.PagosResp, error)
	Update(ctx context.Context, ID string, pagos models.UpdatePagosReq) error
	Delete(ctx context.Context, ID string) error
}
