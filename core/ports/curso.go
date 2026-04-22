package ports

import (
	"context"
	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type CursoRepository interface {
	repository.Repository
}

// SaleService interface
type CursoService interface {
	Create(ctx context.Context, curso models.CreateCursoReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.CursoListResponse, error)
	GetInforme(ctx context.Context, filter map[string]interface{}) (*models.CursoInfListResponse, error)
	GetByID(ctx context.Context, ID string) (models.CursoResp, error)
	Update(ctx context.Context, ID string, curso models.UpdateCursoReq) error
	Delete(ctx context.Context, ID string) error
}
