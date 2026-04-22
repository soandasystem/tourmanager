package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type ProgramadRepository interface {
	repository.Repository
	CreateMany(ctx context.Context, entities []interface{}) ([]string, error)
}

// SaleService interface
type ProgramadService interface {
	Create(ctx context.Context, program models.CreateProgramadReq) (string, error)
	CreateMany(ctx context.Context, program []models.CreateProgramadReq) (models.MultiCreationResp, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.ProgramadListResponse, error)
	GetByID(ctx context.Context, ID string) (models.ProgramadResp, error)
	Update(ctx context.Context, ID string, program models.UpdateProgramadReq) error
	Delete(ctx context.Context, ID string, filter map[string]interface{}) error
}
