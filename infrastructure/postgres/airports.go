package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"tourmanager/core/models"
	"tourmanager/core/ports"

	"github.com/antoniomarfa/hexatools/infrastructure"
	"github.com/antoniomarfa/hexatools/wrappers"

	"gorm.io/gorm"
)

// userRepository adapter of an roles repository for postgres
type airportsRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewAirportsRepository(ctx context.Context, db *gorm.DB) ports.AirportsRepository {
	return &airportsRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *airportsRepository) Create(ctx context.Context, airports interface{}) (string, error) {

	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	// Asegúrate de que el tipo del usuario es correcto
	u := airports.(models.CreateAirportsReq)

	var existingAir models.CreateAirportsReq

	err := DB.Model(&existingAir).Where("description = ?", u.Iata).First(&existingAir).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("el aeropuerto con el codigo '" + u.Iata + "' ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar aeropuerto: " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	crDB := infrastructure.GetDBWithSchema(ctx, s.DB)
	if err := crDB.WithContext(ctx).Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *airportsRepository) GetInf(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {

	var registros []models.AirportsInf

	DB := infrastructure.GetDBWithSchema(ctx, s.DB)
	rawQuery := `
	SELECT DISTINCT ON (city) * 
	FROM airports 
	WHERE city <> '' 
	ORDER BY city, id;
`

	if err := DB.Raw(rawQuery).Scan(&registros).Error; err != nil {
		return nil, err
	}

	if len(registros) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	return []interface{}{registros}, nil

}

func (s *airportsRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var registro []models.AirportsResp

	// Crea una consulta base
	query := DB.WithContext(ctx).Model(&models.AirportsResp{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Find(&registro).Error; err != nil {
		return nil, err
	}

	if len(registro) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Mapear a la estructura de respuesta
	response := models.AirportsListResponse{
		Items:      registro,
		TotalCount: 0,
	}

	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil
}

func (s *airportsRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var airports models.AirportsResp

	// Busca el registro en la base de datos utilizando GORM
	if err := DB.WithContext(ctx).Where("id = ?", ID).First(&airports).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &airports, nil
}

func (s *airportsRepository) Update(ctx context.Context, ID string, airports interface{}) error {
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	if err := DB.WithContext(ctx).Model(&models.UpdateAirportsReq{}).Where("id = ?", ID).Updates(airports).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *airportsRepository) Delete(ctx context.Context, ID string, filter map[string]interface{}) error {
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var airports models.CreateAirportsReq

	result := DB.WithContext(ctx).Where("id = ?", ID).Delete(&airports)

	// Comprueba si ocurrió un error durante la eliminación
	if result.Error != nil {
		return wrappers.NewNonExistentErr(result.Error)
	}

	// Comprueba si se eliminaron filas
	if result.RowsAffected == 0 {
		return wrappers.NewNonExistentErr(sql.ErrNoRows) // Manejo de error si no se encontró el registro
	}

	return nil
}
