package ports

import (
	"context"
	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type RegionRepository interface {
	repository.Repository
}

// SaleService interface
type RegionService interface {
	Create(ctx context.Context, region models.CreateRegionReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.RegionListResponse, error)
	GetByID(ctx context.Context, ID string) (models.RegionResp, error)
	Update(ctx context.Context, ID string, region models.UpdateRegionReq) error
	Delete(ctx context.Context, ID string) error
}
