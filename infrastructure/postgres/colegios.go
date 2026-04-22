package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"tourmanager/core/models"
	"tourmanager/core/ports"

	"github.com/antoniomarfa/hexatools/infrastructure"
	"github.com/antoniomarfa/hexatools/wrappers"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Definimos un tipo personalizado para los errores
type ForeignKeyViolationError struct {
	Message string
}

func (e *ForeignKeyViolationError) Error() string {
	return e.Message
}

// userRepository adapter of an roles repository for postgres
type colegiosRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewColegiosRepository(ctx context.Context, db *gorm.DB) ports.ColegiosRepository {
	return &colegiosRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *colegiosRepository) Create(ctx context.Context, colegios interface{}) (string, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	// Asegúrate de que el tipo del usuario es correcto
	u := colegios.(models.CreateColegiosReq)

	var existingColegios models.CreateColegiosReq

	err := DB.Where("codigo = ?", u.Codigo).First(&existingColegios).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("El Colegio con el Codigo '" + u.Codigo + "' ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar colegio: " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	//asigna el schma a usar
	crDB := infrastructure.GetDBWithSchema(ctx, s.DB)
	if err := crDB.WithContext(ctx).Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *colegiosRepository) GetInf(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	return nil, nil
}

func (s *colegiosRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var registro []models.ColegiosResp

	// Crea una consulta base
	query := DB.WithContext(ctx).Model(&models.ColegiosResp{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Order("id ASC").Find(&registro).Error; err != nil {
		return nil, err
	}

	if len(registro) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Mapear a la estructura de respuesta
	response := models.ColegiosListResponse{
		Items:      registro,
		TotalCount: 0,
	}

	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil

}

func (s *colegiosRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var registro models.ColegiosResp

	// Busca el registro en la base de datos utilizando GORM
	if err := DB.WithContext(ctx).Where("id = ?", ID).First(&registro).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, err
	}

	return &registro, nil
}

func (s *colegiosRepository) Update(ctx context.Context, ID string, colegios interface{}) error {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	err := DB.WithContext(ctx).
		Table("establecimientos"). // Usamos directamente el nombre de la tabla
		Where("id = ?", ID).
		Updates(colegios).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // Uso de errors.Is para comparar errores
			return wrappers.NewNonExistentErr(err)
		}
		return err // Devuelve directamente otros errores
	}

	return nil // Retorna nil si no hubo errores

}

func (s *colegiosRepository) Delete(ctx context.Context, ID string, filter map[string]interface{}) error {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)
	var colegios models.CreateColegiosReq

	result := DB.WithContext(ctx).Where("id = ?", ID).Delete(&colegios)

	// Comprueba si ocurrió un error durante la eliminación
	if result.Error != nil {
		if pqErr, ok := result.Error.(*pq.Error); ok {
			// Comprobamos si el código del error es 23503 (violación de clave foránea)
			if pqErr.Code == "23503" {
				return wrappers.NewNonExistentErr(&ForeignKeyViolationError{
					Message: "No se puede eliminar el colegio debido a una violación de clave foránea en otra tabla",
				})
			}
		}
		return wrappers.NewNonExistentErr(result.Error)
	}

	// Comprueba si se eliminaron filas
	if result.RowsAffected == 0 {
		return wrappers.NewNonExistentErr(sql.ErrNoRows) // Manejo de error si no se encontró el registro
	}

	return nil
}

func (s *colegiosRepository) DeleteAll(ctx context.Context, ID string) error {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	//	var colegio models.CreateColegiosReq

	//	result := DB.WithContext(ctx).Where("id = ?", ID).Delete(&colegio)
	result := DB.WithContext(ctx).Delete(&models.Colegios{}, "id = ?", ID)

	// Comprueba si ocurrió un error durante la eliminación
	if result.Error != nil {
		if pqErr, ok := result.Error.(*pq.Error); ok {
			// Comprobamos si el código del error es 23503 (violación de clave foránea)
			if pqErr.Code == "23503" {
				return fmt.Errorf("Error: No se puede eliminar el colegio debido a una violación de clave foránea en otra tabla")
			}
		}
		return wrappers.NewNonExistentErr(result.Error)
	}

	// Comprueba si se eliminaron filas
	if result.RowsAffected == 0 {
		return wrappers.NewNonExistentErr(sql.ErrNoRows) // Manejo de error si no se encontró el registro
	}

	return nil
}
