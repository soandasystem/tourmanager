package services

import (
	"context"
	"errors"
	"fmt"

	"tourmanager/config"
	"tourmanager/core/models"
	"tourmanager/core/ports"

	"github.com/antoniomarfa/hexatools/wrappers"
)

// rolesService adapter of an user service
type permissionService struct {
	config     config.Config
	repository ports.PermissionRepository
}

// NewURolesService creates a new user service
func NewPermissionService(cfg config.Config, repo ports.PermissionRepository) ports.PermissionService {
	return &permissionService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *permissionService) Create(ctx context.Context, ermission models.CreateRolesPermissionsReq) (string, error) {

	insertedID, err := p.repository.Create(ctx, models.CreateRolesPermissionsReq(ermission))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *permissionService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.RolesPermissionsListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.RolesPermissionsListResponse{
			Items:      []models.RolesPermissionsResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.RolesPermissionsListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID roles permission
func (p *permissionService) GetByID(ctx context.Context, ID string) (resp models.RolesPermissionsResp, err error) {
	permission, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("permission con ID %s no encontrado", ID)
	}

	if permission == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.RolesPermissionsResp{}, fmt.Errorf("permission con ID %s no encontrado", ID)
	}

	resp = *permission.(*models.RolesPermissionsResp)

	return
}

// Update user
func (p *permissionService) Update(ctx context.Context, ID string, permission models.UpdateRolesPermissionsReq) (err error) {

	dbPermission, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	if permission.RolesId != nil {
		dbPermission.RolesId = *permission.RolesId
	}
	if permission.Permission != nil {
		dbPermission.Permission = *permission.Permission
	}
	if permission.Actions != nil {
		dbPermission.Actions = *permission.Actions
	}

	err = p.repository.Update(ctx, ID, models.RolesPermissions(dbPermission))

	return err
}

// Delete user
func (p *permissionService) Delete(ctx context.Context, ID string, filter map[string]interface{}) (err error) {
	err = p.repository.Delete(ctx, ID, filter)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
