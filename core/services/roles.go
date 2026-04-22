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
type rolesService struct {
	config     config.Config
	repository ports.RolesRepository
}

// NewURolesService creates a new user service
func NewRolesService(cfg config.Config, repo ports.RolesRepository) ports.RolesService {
	return &rolesService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *rolesService) Create(ctx context.Context, roles models.CreateRolesReq) (string, error) {

	insertedID, err := p.repository.Create(ctx, models.CreateRolesReq(roles))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *rolesService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.RolesListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.RolesListResponse{
			Items:      []models.RolesResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.RolesListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID user
func (p *rolesService) GetByID(ctx context.Context, ID string) (resp models.RolesResp, err error) {
	roles, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("roles con ID %s no encontrado", ID)
	}

	if roles == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.RolesResp{}, fmt.Errorf("roles con ID %s no encontrado", ID)
	}

	resp = *roles.(*models.RolesResp)

	return
}

// Update user
func (p *rolesService) Update(ctx context.Context, ID string, roles models.UpdateRolesReq) (err error) {

	dbRoles, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	now := time.Now().UTC()
	if roles.Active != nil {
		dbRoles.Active = *roles.Active
	}

	if roles.Author != nil {
		dbRoles.Author = *roles.Author
	}

	if roles.Description != nil {
		dbRoles.Description = *roles.Description
	}

	if roles.Permissions != nil {
		dbRoles.Permissions = *roles.Permissions
	}

	dbRoles.UpdatedDate = now
	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Roles(dbRoles))

	return err
}

// Delete user
func (p *rolesService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
