package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type ComunasRepository interface {
	repository.Repository
}

// RolesService interface
type ComunasService interface {
	Create(ctx context.Context, comunas models.CreateComunasReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.ComunasListResponse, error)
	GetByID(ctx context.Context, ID string) (models.ComunasResp, error)
	Update(ctx context.Context, ID string, comunas models.UpdateComunasReq) error
	Delete(ctx context.Context, ID string) error
}
