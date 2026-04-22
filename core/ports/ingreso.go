package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// IngresoRepositoy interface
type IngresoRepository interface {
	repository.Repository
}

// IngresoService interface
type IngresoService interface {
	Create(ctx context.Context, ingreso models.CreateIngresoReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.IngresoListResponse, error)
	GetInforme(ctx context.Context, filter map[string]interface{}) (*models.IngresoInfListResponse, error)
	GetByID(ctx context.Context, ID string) (models.IngresoResp, error)
	Update(ctx context.Context, ID string, ingreso models.UpdateIngresoReq) error
	Delete(ctx context.Context, ID string) error
}
