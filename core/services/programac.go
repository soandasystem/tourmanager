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
type programacService struct {
	config     config.Config
	repository ports.ProgramacRepository
}

// NewURolesService creates a new user service
func NewProgramacService(cfg config.Config, repo ports.ProgramacRepository) ports.ProgramacService {
	return &programacService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *programacService) Create(ctx context.Context, program models.CreateProgramacReq) (string, error) {

	insertedID, err := p.repository.Create(ctx, models.CreateProgramacReq(program))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *programacService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.ProgramacListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.ProgramacListResponse{
			Items:      []models.ProgramacResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.ProgramacListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID user
func (p *programacService) GetByID(ctx context.Context, ID string) (resp models.ProgramacResp, err error) {
	program, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("programa con ID %s no encontrado", ID)
	}

	if program == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.ProgramacResp{}, fmt.Errorf("programa con ID %s no encontrado", ID)
	}

	resp = *program.(*models.ProgramacResp)

	return
}

// Update user
func (p *programacService) Update(ctx context.Context, ID string, program models.UpdateProgramacReq) (err error) {

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

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Programac(dbProgram))

	return err
}

// Delete user
func (p *programacService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
