package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type GatewayscRepository interface {
	repository.Repository
	GetByIDUpdate(ctx context.Context, ID string) (interface{}, error)
}

// RolesService interface
type GatewayscService interface {
	Create(ctx context.Context, gatewaysc models.CreateGatewayscReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.GatewayscListResponse, error)
	GetByID(ctx context.Context, ID string) (models.GatewayscResp, error)
	GetUpdateByID(ctx context.Context, ID string) (models.GatewayscResp, error)
	Update(ctx context.Context, ID string, gatewaysc models.UpdateGatewayscReq) error
	Delete(ctx context.Context, ID string) error
}
