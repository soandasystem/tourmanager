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
type saleService struct {
	config     config.Config
	repository ports.CompanyRepository
}

// NewURolesService creates a new user service
func NewSaleService(cfg config.Config, repo ports.SaleRepository) ports.SaleService {
	return &saleService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *saleService) Create(ctx context.Context, sale models.CreateSaleReq) (string, error) {

	insertedID, err := p.repository.Create(ctx, models.CreateSaleReq(sale))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *saleService) GetInforme(ctx context.Context, filter map[string]interface{}) (*models.SaleInfListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.GetInf(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.SaleInfListResponse{
			Items:      []models.SaleInforme{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.SaleInfListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetAll users
func (p *saleService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.SaleListResponse, error) {
	// Obtiene las vemtas desde el repositorio

	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.SaleListResponse{
			Items:      []models.SaleResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.SaleListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetByID user
func (p *saleService) GetByID(ctx context.Context, ID string) (resp models.SaleResp, err error) {
	sale, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, nil
	}

	if sale == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.SaleResp{}, fmt.Errorf("venta con ID %s no encontrado", ID)
	}

	//	resp = models.ColegiosResp(*colegios.(*models.ColegiosResp))
	resp = *sale.(*models.SaleResp)
	return resp, nil
	// return
}

// Update user
func (p *saleService) Update(ctx context.Context, ID string, sale models.UpdateSaleReq) (err error) {
	dbSale, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}

	if sale.Fecha != nil {
		dbSale.Fecha = *sale.Fecha
	}
	if sale.SellerId != nil {
		dbSale.SellerId = *sale.SellerId
	}
	if sale.Identificador != nil {
		dbSale.Identificador = *sale.Identificador
	}
	if sale.EstablecimientoId != nil {
		dbSale.EstablecimientoId = *sale.EstablecimientoId
	}
	if sale.ProgramId != nil {
		dbSale.ProgramId = *sale.ProgramId
	}
	if sale.Curso != nil {
		dbSale.Curso = *sale.Curso
	}
	if sale.Idcurso != nil {
		dbSale.Idcurso = *sale.Idcurso
	}
	if sale.Nroalumno != nil {
		dbSale.Nroalumno = *sale.Nroalumno
	}
	if sale.Liberados != nil {
		dbSale.Liberados = *sale.Liberados
	}
	if sale.Subtotal != nil {
		dbSale.Subtotal = *sale.Subtotal
	}
	if sale.Descm != nil {
		dbSale.Descm = *sale.Descm
	}
	if sale.Vprograma != nil {
		dbSale.Vprograma = *sale.Vprograma
	}
	if sale.Description != nil {
		dbSale.Description = *sale.Description
	}
	if sale.Obs != nil {
		dbSale.Obs = *sale.Obs
	}
	if sale.Fechasalida != nil {
		dbSale.Fechasalida = *sale.Fechasalida
	}
	if sale.Activo != nil {
		dbSale.Activo = *sale.Activo
	}
	if sale.State != nil {
		dbSale.State = *sale.State
	}
	if sale.CorreoEncargado != nil {
		dbSale.CorreoEncargado = *sale.CorreoEncargado
	}
	if sale.Password != nil {
		dbSale.Password = *sale.Password
	}
	if sale.FechaUltpag != nil {
		dbSale.FechaUltpag = *sale.FechaUltpag
	}
	if sale.FechaCierre != nil {
		dbSale.FechaCierre = *sale.FechaCierre
	}
	if sale.Sendemail != nil {
		dbSale.Sendemail = *sale.Sendemail
	}
	if sale.Author != nil {
		dbSale.Author = *sale.Author
	}
	if sale.Encargado != nil {
		dbSale.Encargado = *sale.Encargado
	}
	if sale.Comision != nil {
		dbSale.Comision = *sale.Comision
	}
	if sale.Tipocambio != nil {
		dbSale.Tipocambio = *sale.Tipocambio
	}
	if sale.ComisionPagada != nil {
		dbSale.ComisionPagada = *sale.ComisionPagada
	}
	if sale.CompanyId != nil {
		dbSale.CompanyId = *sale.CompanyId
	}
	err = p.repository.Update(ctx, ID, models.SaleResp(dbSale))

	return err

}

// Delete user
func (p *saleService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
