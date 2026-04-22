package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type AirportsRepository interface {
	repository.Repository
}

// RolesService interface
type AirportsService interface {
	Create(ctx context.Context, airports models.CreateAirportsReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.AirportsListResponse, error)
	GetInforme(ctx context.Context, filter map[string]interface{}) (*models.AirportsListResponse, error)
	GetByID(ctx context.Context, ID string) (models.AirportsResp, error)
	Update(ctx context.Context, ID string, Airports models.UpdateAirportsReq) error
	Delete(ctx context.Context, ID string) error
}
