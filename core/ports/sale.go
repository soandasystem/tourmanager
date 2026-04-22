package ports

import (
	"context"
	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type SaleRepository interface {
	repository.Repository
}

// SaleService interface
type SaleService interface {
	Create(ctx context.Context, sales models.CreateSaleReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.SaleListResponse, error)
	GetInforme(ctx context.Context, filter map[string]interface{}) (*models.SaleInfListResponse, error)
	GetByID(ctx context.Context, ID string) (models.SaleResp, error)
	Update(ctx context.Context, ID string, sale models.UpdateSaleReq) error
	Delete(ctx context.Context, ID string) error
}
