package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type CountryRepository interface {
	repository.Repository
}

// RolesService interface
type CountryService interface {
	Create(ctx context.Context, country models.CreateCountryReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.CountryListResponse, error)
	GetByID(ctx context.Context, ID string) (models.CountryResp, error)
	Update(ctx context.Context, ID string, country models.UpdateCountryReq) error
	Delete(ctx context.Context, ID string) error
}
