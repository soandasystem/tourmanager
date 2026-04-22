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
type cursoService struct {
	config     config.Config
	repository ports.CursoRepository
}

// NewURolesService creates a new user service
func NewCursoService(cfg config.Config, repo ports.CursoRepository) ports.CursoService {
	return &cursoService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *cursoService) Create(ctx context.Context, curso models.CreateCursoReq) (string, error) {

	now := time.Now().UTC()
	curso.CreatedDate = now
	curso.UpdatedDate = now
	insertedID, err := p.repository.Create(ctx, models.CreateCursoReq(curso))
	if err != nil {
		return "", err
	}

	return insertedID, err
}
func (p *cursoService) GetInforme(ctx context.Context, filter map[string]interface{}) (*models.CursoInfListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.GetInf(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.CursoInfListResponse{
			Items:      []models.CursoInf{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.CursoInfListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetAll users
func (p *cursoService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.CursoListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.CursoListResponse{
			Items:      []models.CursoResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.CursoListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID user
func (p *cursoService) GetByID(ctx context.Context, ID string) (resp models.CursoResp, err error) {
	curso, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	if curso == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.CursoResp{}, fmt.Errorf("curso con ID %s no encontrado", ID)
	}

	resp = *curso.(*models.CursoResp)

	return
}

// Update user
func (p *cursoService) Update(ctx context.Context, ID string, curso models.UpdateCursoReq) (err error) {

	dbCurso, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	if curso.SaleId != nil {
		dbCurso.SaleId = *curso.SaleId
	}
	if curso.Rutalumno != nil {
		dbCurso.Rutalumno = *curso.Rutalumno
	}
	if curso.Nombrealumno != nil {
		dbCurso.Nombrealumno = *curso.Nombrealumno
	}
	if curso.Fechanac != nil {
		dbCurso.Fechanac = *curso.Fechanac
	}
	if curso.Rutapod != nil {
		dbCurso.Rutapod = *curso.Rutapod
	}
	if curso.Nombreapod != nil {
		dbCurso.Nombreapod = *curso.Nombreapod
	}
	if curso.Dircalle != nil {
		dbCurso.Dircalle = *curso.Dircalle
	}
	if curso.Dirnumero != nil {
		dbCurso.Dirnumero = *curso.Dirnumero
	}
	if curso.Nrodepto != nil {
		dbCurso.Nrodepto = *curso.Nrodepto
	}
	if curso.RegionId != nil {
		dbCurso.RegionId = *curso.RegionId
	}
	if curso.ComunaId != nil {
		dbCurso.ComunaId = *curso.ComunaId
	}
	if curso.Fono != nil {
		dbCurso.Fono = *curso.Fono
	}
	if curso.Celular != nil {
		dbCurso.Celular = *curso.Celular
	}
	if curso.Correo != nil {
		dbCurso.Correo = *curso.Correo
	}
	if curso.Vpagar != nil {
		dbCurso.Vpagar = *curso.Vpagar
	}
	if curso.Descto != nil {
		dbCurso.Descto = *curso.Descto
	}
	if curso.Apagar != nil {
		dbCurso.Apagar = *curso.Apagar
	}
	if curso.Liberado != nil {
		dbCurso.Liberado = *curso.Liberado
	}
	if curso.Enviado != nil {
		dbCurso.Enviado = *curso.Enviado
	}
	if curso.Estado != nil {
		dbCurso.Estado = *curso.Estado
	}
	if curso.Password != nil {
		dbCurso.Password = *curso.Password
	}
	if curso.AceptaContrato != nil {
		dbCurso.AceptaContrato = *curso.AceptaContrato
	}
	if curso.Signature != nil {
		dbCurso.Signature = *curso.Signature
	}
	if curso.Author != nil {
		dbCurso.Author = *curso.Author
	}
	if curso.Pasaporte != nil {
		dbCurso.Pasaporte = *curso.Pasaporte
	}
	/*
		if curso.CompanyId != nil {
			dbCurso.CompanyId = *curso
		}
	*/
	// Actualizar la fecha de modificación
	dbCurso.UpdatedDate = time.Now().UTC()

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Curso(dbCurso))

	return err
}

// Delete user
func (p *cursoService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
