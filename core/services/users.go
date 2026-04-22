package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"tourmanager/config"
	"tourmanager/core/models"
	"tourmanager/core/ports"

	"github.com/antoniomarfa/hexatools/wrappers"
)

// rolesService adapter of an user service
type usersService struct {
	config     config.Config
	repository ports.UsersRepository
}

// NewURolesService creates a new user service
func NewUsersService(cfg config.Config, repo ports.UsersRepository) ports.UsersService {
	return &usersService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *usersService) Create(ctx context.Context, users models.CreateUsersReq) (string, error) {

	insertedID, err := p.repository.Create(ctx, models.CreateUsersReq(users))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *usersService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.UsersListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.UsersListResponse{
			Items:      []models.UsersInf{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.UsersListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID user
func (p *usersService) GetByID(ctx context.Context, ID string) (resp models.UsersInf, err error) {
	users, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("usuario con ID %s no encontrado", ID)
	}

	if users == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.UsersInf{}, fmt.Errorf("usuario con ID %s no encontrado", ID)
	}

	resp = *users.(*models.UsersInf)

	return
}

func (p *usersService) GetUpdateByID(ctx context.Context, ID string) (resp models.UsersResp, err error) {
	users, err := p.repository.GetByIDUpdate(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("usuario con ID %s no encontrado", ID)
	}

	if users == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.UsersResp{}, fmt.Errorf("usuario con ID %s no encontrado", ID)
	}

	resp = *users.(*models.UsersResp)

	return
}

// Update user
func (p *usersService) Update(ctx context.Context, ID string, users models.UpdateUsersReq) (err error) {

	dbUsers, err := p.GetUpdateByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	now := time.Now().UTC()
	if users.Username != nil {
		dbUsers.Username = *users.Username
	}
	if users.Name != nil {
		dbUsers.Name = *users.Name
	}
	if users.Email != nil {
		dbUsers.Email = *users.Email
	}
	if users.Password != nil {
		dbUsers.Password = *users.Password
	}
	if users.Phone != nil {
		dbUsers.Phone = *users.Phone
	}
	if users.RolesId != nil {
		dbUsers.RolesId = *users.RolesId
	}
	if users.Active != nil {
		dbUsers.Active = *users.Active
	}
	if users.Author != nil {
		dbUsers.Author = *users.Author
	}
	if users.ResetToken != nil {
		dbUsers.ResetToken = *users.ResetToken
	}
	if users.ResetTokenExpira != nil {
		dbUsers.ResetTokenExpira = *users.ResetTokenExpira
	}

	dbUsers.UpdatedDate = now
	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Users(dbUsers))

	return err
}

// Delete user
func (p *usersService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
