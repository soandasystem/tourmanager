package services

import (
	"context"
	"errors"
	"fmt"

	"tourmanager/config"
	"tourmanager/core/entities"
	"tourmanager/core/models"
	"tourmanager/core/ports"

	"github.com/antoniomarfa/hexatools/wrappers"
)

// rolesService adapter of an user service
type programadService struct {
	config     config.Config
	repository ports.ProgramadRepository
}

// NewURolesService creates a new user service
func NewProgramadService(cfg config.Config, repo ports.ProgramadRepository) ports.ProgramadService {
	return &programadService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *programadService) Create(ctx context.Context, programad models.CreateProgramadReq) (string, error) {

	insertedID, err := p.repository.Create(ctx, models.CreateProgramadReq(programad))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// CreateMany users
func (p *programadService) CreateMany(ctx context.Context, programd []models.CreateProgramadReq) (resp models.MultiCreationResp, err error) {
	var create []interface{}

	for _, programad := range programd {
		create = append(create, entities.Programad(programad))
	}

	insertedIDs, err := p.repository.CreateMany(ctx, create)
	if err != nil {
		return
	}

	resp = models.MultiCreationResp{
		InsertedIDs: insertedIDs,
	}
	return
}

// GetAll users
func (p *programadService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.ProgramadListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.ProgramadListResponse{
			Items:      []models.ProgramadResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.ProgramadListResponse)
	if !ok {
		fmt.Println(ok)
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID user
func (p *programadService) GetByID(ctx context.Context, ID string) (resp models.ProgramadResp, err error) {
	program, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("programa con ID %s no encontrado", ID)
	}

	if program == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.ProgramadResp{}, fmt.Errorf("programa con ID %s no encontrado", ID)
	}

	resp = *program.(*models.ProgramadResp)

	return
}

// Update user
func (p *programadService) Update(ctx context.Context, ID string, program models.UpdateProgramadReq) (err error) {

	dbProgramad, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	if program.Desde != nil {
		dbProgramad.Desde = *program.Desde
	}
	if program.Hasta != nil {
		dbProgramad.Hasta = *program.Hasta
	}
	if program.Liberado != nil {
		dbProgramad.Liberado = *program.Liberado
	}
	if program.Valor != nil {
		dbProgramad.Valor = *program.Valor
	}

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Programad(dbProgramad))

	return err
}

// Delete user
func (p *programadService) Delete(ctx context.Context, ID string, filter map[string]interface{}) (err error) {
	err = p.repository.Delete(ctx, ID, filter)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
