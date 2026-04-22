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
	"github.com/google/uuid"
)

// rolesService adapter of an user service
type companyService struct {
	config     config.Config
	repository ports.CompanyRepository
}

// NewURolesService creates a new user service
func NewCompanyService(cfg config.Config, repo ports.CompanyRepository) ports.CompanyService {
	return &companyService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *companyService) Create(ctx context.Context, company models.CreateCompanyReq) (string, error) {

	now := time.Now().UTC()
	company.CreatedDate = now
	company.UpdatedDate = now

	var str string = company.Razonsocial
	prefix := str[0:3]          // Esto obtiene los caracteres desde la posición 0 hasta la 2 (inclusive)
	uuidStr := uuid.NewString() // UUID como string

	prefixedID := prefix + "_" + uuidStr
	company.Identificador = prefixedID

	insertedID, err := p.repository.Create(ctx, models.CreateCompanyReq(company))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *companyService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.CompanyListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.CompanyListResponse{
			Items:      []models.CompanyResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.CompanyListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID user
func (p *companyService) GetByID(ctx context.Context, ID string) (resp models.CompanyResp, err error) {
	company, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("colegio con ID %s no encontrado", ID)
	}

	if company == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.CompanyResp{}, fmt.Errorf("colegio con ID %s no encontrado", ID)
	}

	resp = *company.(*models.CompanyResp)

	return
}

// Update user
func (p *companyService) Update(ctx context.Context, ID string, company models.UpdateCompanyReq) (err error) {

	dbCompany, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	if company.Rut != nil {
		dbCompany.Rut = *company.Rut
	}

	if company.Razonsocial != nil {
		dbCompany.Razonsocial = *company.Razonsocial
	}

	if company.Nomfantasia != nil {
		dbCompany.Nomfantasia = *company.Nomfantasia
	}

	if company.Rutreplegal != nil {
		dbCompany.Rutreplegal = *company.Rutreplegal
	}

	if company.Nomreplegal != nil {
		dbCompany.Nomreplegal = *company.Nomreplegal
	}

	if company.Contrato != nil {
		dbCompany.Contrato = *company.Contrato
	}

	if company.ContratoVg != nil {
		dbCompany.ContratoVg = *company.ContratoVg
	}

	if company.Website != nil {
		dbCompany.Website = *company.Website
	}

	if company.ComunaId != nil {
		dbCompany.ComunaId = *company.ComunaId
	}

	if company.Direccion != nil {
		dbCompany.Direccion = *company.Direccion
	}

	if company.RegionId != nil {
		dbCompany.RegionId = *company.RegionId
	}

	if company.Nombrecontacto1 != nil {
		dbCompany.Nombrecontacto1 = *company.Nombrecontacto1
	}

	if company.Fonocontacto1 != nil {
		dbCompany.Fonocontacto1 = *company.Fonocontacto1
	}

	if company.Emailcontacto1 != nil {
		dbCompany.Emailcontacto1 = *company.Emailcontacto1
	}

	if company.Nombrecontacto2 != nil {
		dbCompany.Nombrecontacto2 = *company.Nombrecontacto2
	}

	if company.Fonocontacto2 != nil {
		dbCompany.Fonocontacto2 = *company.Fonocontacto2
	}

	if company.Emailcontacto2 != nil {
		dbCompany.Emailcontacto2 = *company.Emailcontacto2
	}

	if company.Author != nil {
		dbCompany.Author = *company.Author
	}

	if company.Email != nil {
		dbCompany.Email = *company.Email
	}

	// Asegúrate de que Active no sea nil antes de asignarlo
	if company.Active != nil {
		dbCompany.Active = *company.Active
	}

	if company.PlancodeId != nil {
		dbCompany.PlancodeId = *company.PlancodeId
	}

	if company.Additionaluser != nil {
		dbCompany.Additionaluser = *company.Additionaluser
	}

	if company.Maxusers != nil {
		dbCompany.Maxusers = *company.Maxusers
	}

	if company.Maxquote != nil {
		dbCompany.Maxquote = *company.Maxquote
	}

	if company.Maxsales != nil {
		dbCompany.Maxsales = *company.Maxsales
	}

	if company.Terminoscondiciones != nil {
		dbCompany.Terminoscondiciones = *company.Terminoscondiciones
	}

	if company.Politicasdeuso != nil {
		dbCompany.Politicasdeuso = *company.Politicasdeuso
	}

	// Actualizar la fecha de modificación
	dbCompany.UpdatedDate = time.Now().UTC()

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Company(dbCompany))

	return err
}

// Delete user
func (p *companyService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
