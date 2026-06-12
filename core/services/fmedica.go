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
	if ficha.SaleId != nil {
		dbFmedica.SaleId = *ficha.SaleId
	}
	if ficha.CursoId != nil {
		dbFmedica.CursoId = *ficha.CursoId
	}
	if ficha.CompanyId != nil {
		dbFmedica.CompanyId = *ficha.CompanyId
	}
	if ficha.GrupoSanguineo != nil {
		dbFmedica.GrupoSanguineo = *ficha.GrupoSanguineo
	}
	if ficha.Edad != nil {
		dbFmedica.Edad = *ficha.Edad
	}
	if ficha.Peso != nil {
		dbFmedica.Peso = *ficha.Peso
	}
	if ficha.Estatura != nil {
		dbFmedica.Estatura = *ficha.Estatura
	}
	if ficha.Hipertension != nil {
		dbFmedica.Hipertension = *ficha.Hipertension
	}
	if ficha.Diabetes != nil {
		dbFmedica.Diabetes = *ficha.Diabetes
	}
	if ficha.Asma != nil {
		dbFmedica.Asma = *ficha.Asma
	}
	if ficha.Epilepsia != nil {
		dbFmedica.Epilepsia = *ficha.Epilepsia
	}
	if ficha.Arritmias != nil {
		dbFmedica.Arritmias = *ficha.Arritmias
	}
	if ficha.EnfermedadesCardiacas != nil {
		dbFmedica.EnfermedadesCardiacas = *ficha.EnfermedadesCardiacas
	}
	if ficha.EnfermedadesRespiratorias != nil {
		dbFmedica.EnfermedadesRespiratorias = *ficha.EnfermedadesRespiratorias
	}
	if ficha.EnfermedadesRenales != nil {
		dbFmedica.EnfermedadesRenales = *ficha.EnfermedadesRenales
	}
	if ficha.OtrasEnfermedades != nil {
		dbFmedica.OtrasEnfermedades = *ficha.OtrasEnfermedades
	}
	if ficha.BajoTratamiento != nil {
		dbFmedica.BajoTratamiento = *ficha.BajoTratamiento
	}
	if ficha.Diagnostico != nil {
		dbFmedica.Diagnostico = *ficha.Diagnostico
	}
	if ficha.MedicamentosTratamiento != nil {
		dbFmedica.MedicamentosTratamiento = *ficha.MedicamentosTratamiento
	}
	if ficha.DosisTratamiento != nil {
		dbFmedica.DosisTratamiento = *ficha.DosisTratamiento
	}
	if ficha.AlergiaMedicamentos != nil {
		dbFmedica.AlergiaMedicamentos = *ficha.AlergiaMedicamentos
	}
	if ficha.AlergiaAlimentos != nil {
		dbFmedica.AlergiaAlimentos = *ficha.AlergiaAlimentos
	}
	if ficha.AlergiaInsectos != nil {
		dbFmedica.AlergiaInsectos = *ficha.AlergiaInsectos
	}
	if ficha.AlergiaOtras != nil {
		dbFmedica.AlergiaOtras = *ficha.AlergiaOtras
	}
	if ficha.DificultadMovilidad != nil {
		dbFmedica.DificultadMovilidad = *ficha.DificultadMovilidad
	}
	if ficha.AsistenciaMovilidad != nil {
		dbFmedica.AsistenciaMovilidad = *ficha.AsistenciaMovilidad
	}
	if ficha.AsistenciaEspecial != nil {
		dbFmedica.AsistenciaEspecial = *ficha.AsistenciaEspecial
	}
	if ficha.Vegetariano != nil {
		dbFmedica.Vegetariano = *ficha.Vegetariano
	}
	if ficha.Vegano != nil {
		dbFmedica.Vegano = *ficha.Vegano
	}
	if ficha.Celiaco != nil {
		dbFmedica.Celiaco = *ficha.Celiaco
	}
	if ficha.IntoleranciaLactosa != nil {
		dbFmedica.IntoleranciaLactosa = *ficha.IntoleranciaLactosa
	}
	if ficha.DiabeticoAlim != nil {
		dbFmedica.DiabeticoAlim = *ficha.DiabeticoAlim
	}
	if ficha.OtraRestriccionAlim != nil {
		dbFmedica.OtraRestriccionAlim = *ficha.OtraRestriccionAlim
	}
	if ficha.ContactoNombre != nil {
		dbFmedica.ContactoNombre = *ficha.ContactoNombre
	}
	if ficha.ContactoRelacion != nil {
		dbFmedica.ContactoRelacion = *ficha.ContactoRelacion
	}
	if ficha.ContactoTelefono1 != nil {
		dbFmedica.ContactoTelefono1 = *ficha.ContactoTelefono1
	}
	if ficha.ContactoTelefono2 != nil {
		dbFmedica.ContactoTelefono2 = *ficha.ContactoTelefono2
	}
	if ficha.Observaciones != nil {
		dbFmedica.Observaciones = *ficha.Observaciones
	}
	if ficha.Author != nil {
		dbFmedica.Author = *ficha.Author
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
