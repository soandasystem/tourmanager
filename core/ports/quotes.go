package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type QuotesRepository interface {
	repository.Repository
}

// SaleService interface
type QuotesService interface {
	Create(ctx context.Context, sales models.CreateQuoteReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.QuoteListResponse, error)
	GetInforme(ctx context.Context, filter map[string]interface{}) (*models.QuoteInfListResponse, error)
	GetByID(ctx context.Context, ID string) (models.QuoteResp, error)
	Update(ctx context.Context, ID string, sale models.UpdateQuoteReq) error
	Delete(ctx context.Context, ID string) error
}
