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
type programsService struct {
	config     config.Config
	repository ports.ProgramsRepository
}

// NewURolesService creates a new user service
func NewProgramsService(cfg config.Config, repo ports.ProgramsRepository) ports.ProgramsService {
	return &programsService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *programsService) Create(ctx context.Context, program models.CreateProgramsReq) (string, error) {

	insertedID, err := p.repository.Create(ctx, models.CreateProgramsReq(program))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *programsService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.ProgramsListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.ProgramsListResponse{
			Items:      []models.ProgramsResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.ProgramsListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID user
func (p *programsService) GetByID(ctx context.Context, ID string) (resp models.ProgramsResp, err error) {
	program, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("programa con ID %s no encontrado", ID)
	}

	if program == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.ProgramsResp{}, fmt.Errorf("programa con ID %s no encontrado", ID)
	}

	resp = *program.(*models.ProgramsResp)

	return
}

// Update user
func (p *programsService) Update(ctx context.Context, ID string, program models.UpdateProgramsReq) (err error) {

	dbProgram, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	if program.Name != nil {
		dbProgram.Name = *program.Name
	}
	if program.Active != nil {
		dbProgram.Active = *program.Active
	}
	if program.Reserva != nil {
		dbProgram.Reserva = *program.Reserva
	}
	if program.Author != nil {
		dbProgram.Author = *program.Author
	}
	if program.OriginCode != nil {
		dbProgram.OriginCode = *program.OriginCode
	}
	if program.DestinationCode != nil {
		dbProgram.DestinationCode = *program.DestinationCode
	}
	if program.Origin != nil {
		dbProgram.Origin = *program.Origin
	}
	if program.Destination != nil {
		dbProgram.Destination = *program.Destination
	}
	if program.Matrix != nil {
		dbProgram.Matrix = *program.Matrix
	}

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Programs(dbProgram))

	return err
}

// Delete user
func (p *programsService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
