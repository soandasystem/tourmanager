package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type GatewaysRepository interface {
	repository.Repository
	GetByIDUpdate(ctx context.Context, ID string) (interface{}, error)
}

// RolesService interface
type GatewaysService interface {
	Create(ctx context.Context, gateways models.CreateGatewaysReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.GatewaysListResponse, error)
	GetByID(ctx context.Context, ID string) (models.GatewaysResp, error)
	GetUpdateByID(ctx context.Context, ID string) (models.Gateways, error)
	Update(ctx context.Context, ID string, gateways models.Gateways) error
	Delete(ctx context.Context, ID string) error
}
