package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// ProgramsRepository interface
type ProgramsRepository interface {
	repository.Repository
}

// ProgramsService interface
type ProgramsService interface {
	Create(ctx context.Context, program models.CreateProgramsReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.ProgramsListResponse, error)
	GetByID(ctx context.Context, ID string) (models.ProgramsResp, error)
	Update(ctx context.Context, ID string, program models.UpdateProgramsReq) error
	Delete(ctx context.Context, ID string) error
}
