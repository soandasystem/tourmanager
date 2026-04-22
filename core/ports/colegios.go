package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type ColegiosRepository interface {
	repository.Repository
}

// RolesService interface
type ColegiosService interface {
	Create(ctx context.Context, colegios models.CreateColegiosReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.ColegiosListResponse, error)
	GetByID(ctx context.Context, ID string) (models.ColegiosResp, error)
	Update(ctx context.Context, ID string, colegios models.UpdateColegiosReq) error
	Delete(ctx context.Context, ID string) error
}
