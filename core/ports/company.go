package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type CompanyRepository interface {
	repository.Repository
}

// RolesService interface
type CompanyService interface {
	Create(ctx context.Context, company models.CreateCompanyReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.CompanyListResponse, error)
	GetByID(ctx context.Context, ID string) (models.CompanyResp, error)
	Update(ctx context.Context, ID string, company models.UpdateCompanyReq) error
	Delete(ctx context.Context, ID string) error
}
