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
type countryRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewCountryRepository(ctx context.Context, db *gorm.DB) ports.CountryRepository {
	return &countryRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *countryRepository) Create(ctx context.Context, country interface{}) (string, error) {
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)
	// Asegúrate de que el tipo del usuario es correcto
	u := country.(models.CreateCountryReq)

	var existingReg models.CreateCountryReq

	err := DB.Model(&existingReg).Where("code = ?", u.Code).First(&existingReg).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("El pais con el codigo '" + u.Code + "' ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar pais: " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	crDB := infrastructure.GetDBWithSchema(ctx, s.DB)
	if err := crDB.WithContext(ctx).Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *countryRepository) GetInf(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	return nil, nil
}

func (s *countryRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)
	var registro []models.CountryResp

	// Crea una consulta base
	query := DB.WithContext(ctx).Model(&models.CountryResp{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Order("id asc").Find(&registro).Error; err != nil {
		return nil, err
	}

	if len(registro) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Mapear a la estructura de respuesta
	response := models.CountryListResponse{
		Items:      registro,
		TotalCount: 0,
	}

	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil
}

func (s *countryRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)
	var registro models.CountryResp

	// Busca el registro en la base de datos utilizando GORM
	if err := DB.WithContext(ctx).Where("id = ?", ID).First(&registro).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &registro, nil
}

func (s *countryRepository) Update(ctx context.Context, ID string, country interface{}) error {
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)
	if err := DB.WithContext(ctx).Model(&models.UpdateCountryReq{}).Where("id = ?", ID).Updates(country).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *countryRepository) Delete(ctx context.Context, ID string, filter map[string]interface{}) error {
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)
	var country models.CreateCountryReq

	result := DB.WithContext(ctx).Where("id = ?", ID).Delete(&country)

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
