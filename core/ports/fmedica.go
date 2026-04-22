package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type FmedicaRepository interface {
	repository.Repository
}

// SaleService interface
type FmedicaService interface {
	Create(ctx context.Context, fmedica models.CreateFmedicaReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.FmedicaListResponse, error)
	GetByID(ctx context.Context, ID string) (models.FmedicaResp, error)
	Update(ctx context.Context, ID string, fmedica models.UpdateFmedicaReq) error
	Delete(ctx context.Context, ID string) error
}
