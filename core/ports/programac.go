package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type ProgramacRepository interface {
	repository.Repository
}

// SaleService interface
type ProgramacService interface {
	Create(ctx context.Context, program models.CreateProgramacReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.ProgramacListResponse, error)
	GetByID(ctx context.Context, ID string) (models.ProgramacResp, error)
	Update(ctx context.Context, ID string, program models.UpdateProgramacReq) error
	Delete(ctx context.Context, ID string) error
}
