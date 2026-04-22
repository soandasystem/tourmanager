package ports

import (
	"context"

	"tourmanager/core/models"

	"github.com/antoniomarfa/hexatools/repository"
)

// UserRepositoy interface
type UsersRepository interface {
	repository.Repository
	GetByIDUpdate(ctx context.Context, ID string) (interface{}, error)
}

// SaleService interface
type UsersService interface {
	Create(ctx context.Context, users models.CreateUsersReq) (string, error)
	GetAll(ctx context.Context, filter map[string]interface{}) (*models.UsersListResponse, error)
	GetByID(ctx context.Context, ID string) (models.UsersInf, error)
	GetUpdateByID(ctx context.Context, ID string) (models.UsersResp, error)
	Update(ctx context.Context, ID string, users models.UpdateUsersReq) error
	Delete(ctx context.Context, ID string) error
}
