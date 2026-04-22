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
type fmedicaService struct {
	config     config.Config
	repository ports.FmedicaRepository
}

// NewURolesService creates a new user service
func NewFmedicaService(cfg config.Config, repo ports.FmedicaRepository) ports.FmedicaService {
	return &fmedicaService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *fmedicaService) Create(ctx context.Context, ficha models.CreateFmedicaReq) (string, error) {

	insertedID, err := p.repository.Create(ctx, models.CreateFmedicaReq(ficha))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *fmedicaService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.FmedicaListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}
	// Convierte los resultados
	if len(result) == 0 {
		return &models.FmedicaListResponse{
			Items:      []models.FmedicaResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.FmedicaListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil

}

// GetByID user
func (p *fmedicaService) GetByID(ctx context.Context, ID string) (resp models.FmedicaResp, err error) {
	ficha, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("ficha medica con ID %s no encontrado", ID)
	}

	if ficha == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.FmedicaResp{}, fmt.Errorf("ficha medica con  con ID %s no encontrado", ID)
	}

	resp = *ficha.(*models.FmedicaResp)

	return
}

// Update user
func (p *fmedicaService) Update(ctx context.Context, ID string, ficha models.UpdateFmedicaReq) (err error) {

	dbFmedica, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	if ficha.Dato1 != nil {
		dbFmedica.Dato1 = *ficha.Dato1
	}
	if ficha.Dato2 != nil {
		dbFmedica.Dato2 = *ficha.Dato2
	}
	if ficha.Dato31 != nil {
		dbFmedica.Dato31 = *ficha.Dato31
	}
	if ficha.Dato32 != nil {
		dbFmedica.Dato32 = *ficha.Dato1
	}
	if ficha.Dato4 != nil {
		dbFmedica.Dato4 = *ficha.Dato4
	}
	if ficha.Dato5 != nil {
		dbFmedica.Dato5 = *ficha.Dato5
	}
	if ficha.Dato6 != nil {
		dbFmedica.Dato6 = *ficha.Dato6
	}
	if ficha.Dato7 != nil {
		dbFmedica.Dato7 = *ficha.Dato7
	}
	if ficha.Dato8 != nil {
		dbFmedica.Dato8 = *ficha.Dato8
	}
	if ficha.Dato9 != nil {
		dbFmedica.Dato9 = *ficha.Dato9
	}
	if ficha.Dato91 != nil {
		dbFmedica.Dato91 = *ficha.Dato91
	}
	if ficha.Dato92 != nil {
		dbFmedica.Dato92 = *ficha.Dato92
	}
	if ficha.Dato10 != nil {
		dbFmedica.Dato10 = *ficha.Dato10
	}
	if ficha.Dato101 != nil {
		dbFmedica.Dato101 = *ficha.Dato101
	}
	if ficha.Dato11 != nil {
		dbFmedica.Dato11 = *ficha.Dato11
	}
	if ficha.Dato111 != nil {
		dbFmedica.Dato111 = *ficha.Dato111
	}
	if ficha.Dato12 != nil {
		dbFmedica.Dato12 = *ficha.Dato12
	}
	if ficha.Dato13 != nil {
		dbFmedica.Dato13 = *ficha.Dato13
	}
	if ficha.Dato141 != nil {
		dbFmedica.Dato141 = *ficha.Dato141
	}
	if ficha.Dato142 != nil {
		dbFmedica.Dato142 = *ficha.Dato142
	}
	if ficha.Dato151 != nil {
		dbFmedica.Dato151 = *ficha.Dato151
	}
	if ficha.Dato152 != nil {
		dbFmedica.Dato152 = *ficha.Dato152
	}
	if ficha.Dato161 != nil {
		dbFmedica.Dato161 = *ficha.Dato161
	}
	if ficha.Dato162 != nil {
		dbFmedica.Dato162 = *ficha.Dato162
	}
	if ficha.Dato17 != nil {
		dbFmedica.Dato17 = *ficha.Dato17
	}
	if ficha.Dato18 != nil {
		dbFmedica.Dato18 = *ficha.Dato18
	}
	if ficha.Dato19 != nil {
		dbFmedica.Dato19 = *ficha.Dato19
	}
	if ficha.Dato20 != nil {
		dbFmedica.Dato20 = *ficha.Dato20
	}
	if ficha.Dato21 != nil {
		dbFmedica.Dato21 = *ficha.Dato21
	}
	if ficha.Dato22 != nil {
		dbFmedica.Dato22 = *ficha.Dato22
	}

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Fmedicas(dbFmedica))

	return err
}

// Delete user
func (p *fmedicaService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
