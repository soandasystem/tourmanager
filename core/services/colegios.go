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
type colegiosService struct {
	config     config.Config
	repository ports.CompanyRepository
}

// NewURolesService creates a new user service
func NewColegiosService(cfg config.Config, repo ports.ColegiosRepository) ports.ColegiosService {
	return &colegiosService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *colegiosService) Create(ctx context.Context, colegios models.CreateColegiosReq) (string, error) {

	insertedID, err := p.repository.Create(ctx, models.CreateColegiosReq(colegios))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *colegiosService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.ColegiosListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.ColegiosListResponse{
			Items:      []models.ColegiosResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.ColegiosListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID user
func (p *colegiosService) GetByID(ctx context.Context, ID string) (resp models.ColegiosResp, err error) {
	colegios, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("colegio con ID %s no encontrado", ID)
	}

	if colegios == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.ColegiosResp{}, fmt.Errorf("colegio con ID %s no encontrado", ID)
	}

	//	resp = models.ColegiosResp(*colegios.(*models.ColegiosResp))
	resp = *colegios.(*models.ColegiosResp)
	return resp, nil
	// return
}

// Update user
func (p *colegiosService) Update(ctx context.Context, ID string, colegios models.UpdateColegiosReq) (err error) {
	dbColegios, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}

	if colegios.Codigo != nil {
		dbColegios.Codigo = *colegios.Codigo
	}

	if colegios.Nombre != nil {
		dbColegios.Nombre = *colegios.Nombre
	}

	if colegios.Direccion != nil {
		dbColegios.Direccion = *colegios.Direccion
	}

	if colegios.Comuna != nil {
		dbColegios.Comuna = *colegios.Comuna
	}

	if colegios.Latitud != nil {
		dbColegios.Latitud = *colegios.Latitud
	}

	if colegios.Longitud != nil {
		dbColegios.Longitud = *colegios.Longitud
	}

	if colegios.RegionId != nil {
		dbColegios.RegionId = *colegios.RegionId
	}

	if colegios.ComunaId != nil {
		dbColegios.ComunaId = *colegios.ComunaId
	}

	if colegios.CompanyId != nil {
		dbColegios.CompanyId = *colegios.CompanyId
	}

	err = p.repository.Update(ctx, ID, models.Colegios(dbColegios))
	return err

}

// Delete user
func (p *colegiosService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		fmt.Println("error ", err)
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
		//		err = wrappers.NewNonExistentErr(err)
	}

	return err
}
