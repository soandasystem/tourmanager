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
type airportsService struct {
	config     config.Config
	repository ports.AirportsRepository
}

// NewURolesService creates a new user service
func NewAirportsService(cfg config.Config, repo ports.AirportsRepository) ports.AirportsService {
	return &airportsService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *airportsService) Create(ctx context.Context, airports models.CreateAirportsReq) (string, error) {

	insertedID, err := p.repository.Create(ctx, models.CreateAirportsReq(airports))
	if err != nil {
		return "", err
	}

	return insertedID, err
}

// GetAll users
func (p *airportsService) GetAll(ctx context.Context, filter map[string]interface{}) (*models.AirportsListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.AirportsListResponse{
			Items:      []models.AirportsResp{},
			TotalCount: 0,
		}, nil
	}

	response, ok := result[0].(models.AirportsListResponse)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	return &response, nil
}

// GetAll users
func (p *airportsService) GetInforme(ctx context.Context, filter map[string]interface{}) (*models.AirportsListResponse, error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.GetInf(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados
	if len(result) == 0 {
		return &models.AirportsListResponse{
			Items:      []models.AirportsResp{},
			TotalCount: 0,
		}, nil
	}

	airportsInf, ok := result[0].([]models.AirportsInf)
	if !ok {
		return nil, fmt.Errorf("tipo de respuesta inesperado del repositorio")
	}

	// Mapeas manualmente a []AirportsResp
	respList := make([]models.AirportsResp, len(airportsInf))
	for i, a := range airportsInf {
		respList[i] = models.AirportsResp{
			City:      a.City,
			ID:        a.ID,
			Icao:      a.Icao,
			Iata:      a.Iata,
			Name:      a.Name,
			State:     a.State,
			Country:   a.Country,
			Elevation: a.Elevation,
			Lat:       a.Lat,
			Lon:       a.Lon,
			Tz:        a.Tz,
		}
	}

	return &models.AirportsListResponse{
		Items:      respList,
		TotalCount: int64(len(respList)),
	}, nil
}

// GetByID user
func (p *airportsService) GetByID(ctx context.Context, ID string) (resp models.AirportsResp, err error) {
	Airports, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("colegio con ID %s no encontrado", ID)
	}

	if Airports == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.AirportsResp{}, fmt.Errorf("colegio con ID %s no encontrado", ID)
	}

	resp = *Airports.(*models.AirportsResp)

	return
}

// Update user
func (p *airportsService) Update(ctx context.Context, ID string, airports models.UpdateAirportsReq) (err error) {

	dbAirports, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	if airports.Icao != nil {
		dbAirports.Iata = *airports.Icao
	}

	if airports.Iata != nil {
		dbAirports.Iata = *airports.Iata
	}

	if airports.Name != nil {
		dbAirports.Name = *airports.Name
	}

	if airports.City != nil {
		dbAirports.City = *airports.City
	}

	if airports.State != nil {
		dbAirports.State = *airports.State
	}

	if airports.Country != nil {
		dbAirports.Country = *airports.Country
	}

	if airports.Elevation != nil {
		dbAirports.Elevation = *airports.Elevation
	}

	if airports.Lat != nil {
		dbAirports.Lat = *airports.Lat
	}

	if airports.Lon != nil {
		dbAirports.Lon = *airports.Lon
	}

	if airports.Tz != nil {
		dbAirports.Tz = *airports.Tz
	}

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Airports(dbAirports))

	return err
}

// Delete user
func (p *airportsService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID, nil)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
