package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"tourmanager/core/models"
	"tourmanager/core/ports"

	"github.com/antoniomarfa/hexatools/infrastructure"
	"github.com/antoniomarfa/hexatools/wrappers"

	"gorm.io/gorm"
)

// userRepository adapter of an roles repository for postgres
type cursoRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewCursoRepository(ctx context.Context, db *gorm.DB) ports.CursoRepository {
	return &cursoRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *cursoRepository) Create(ctx context.Context, curso interface{}) (string, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	tx := DB.Begin()
	if tx.Error != nil {
		return "", errors.New("No se pudo iniciar la transaccion: " + DB.Error.Error())
	}

	// Garantiza rollback si ocurre un panic o error
	// rollback garantizado en caso de error o panic
	defer func() {
		if r := recover(); r != nil {
			DB.Rollback()
			//	panic(r) // puedes remover esto si no quieres propagar el panic
		} else if DB.Error != nil || DB.Statement != nil && DB.Statement.ConnPool != nil {
			_ = DB.Rollback() // rollback silencioso si no se hizo commit
		}
	}()

	// Asegúrate de que el tipo del usuario es correcto
	u, ok := curso.(models.CreateCursoReq)
	if !ok {
		tx.Rollback()
		return "", errors.New("el tipo de curso no es correcto")
	}

	var existingCurso models.CreateCursoReq

	if err := tx.Where("rutalumno = ?", u.Rutalumno).First(&existingCurso).Error; err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		tx.Rollback()
		return "error", errors.New("El pasajero con el rut '" + u.Rutalumno + "' ya existe")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar pasajero " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	if err := tx.WithContext(ctx).Create(&u).Error; err != nil {
		tx.Rollback()
		return "", errors.New("error al crear pasajero " + err.Error())
	}

	// Confirma la transacción
	if err := tx.Commit().Error; err != nil {
		return "", errors.New("no se pudo confirmar la transacción: " + err.Error())
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *cursoRepository) GetInf(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var registro []models.CursoInf
	var totalCount int64

	// 1. Procesar el filtro para separar offset/limit de los WHERE
	if len(filter) != 0 {
		for key, value := range filter {
			if key == "offset" || key == "limit" {
				// Conversión segura a int
				var intVal int
				switch v := value.(type) {
				case int:
					intVal = v
				case float64:
					intVal = int(v)
				case string:
					parsed, err := strconv.Atoi(v)
					if err != nil {
						return nil, nil
					}
					intVal = parsed
				default:
					return nil, nil
				}

				switch key {
				case "offset":
					newSkip := intVal
					skip = &newSkip
				case "limit":
					newTake := intVal
					take = &newTake
				}

				delete(filter, key)
			}
		}
	}

	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	// 2. Consulta para contar el total (sin paginación)
	countQuery := DB.WithContext(ctx).Model(&models.CursoInf{})
	if len(filter) > 0 {
		for key, value := range filter {
			countQuery = countQuery.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}
	if err := countQuery.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	//asigna el schma a usar
	rsDB := infrastructure.GetDBWithSchema(ctx, s.DB)

	// Crea una consulta base
	query := rsDB.WithContext(ctx).Model(&models.CursoInf{})

	// Aplica filtros si existen
	if len(filter) != 0 {
		for key, value := range filter {
			query = query.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}
	// Verificar si skip y take son nil
	fmt.Println("skip:", skip, "take:", take)

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	if err := query.Preload("Sale").Preload("Ingreso").Order("sale_id").Find(&registro).Error; err != nil {
		return nil, err
	}

	//------------------------
	if len(registro) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Mapear a la estructura de respuesta
	response := models.CursoInfListResponse{
		Items:      registro,
		TotalCount: totalCount,
	}
	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil

}

func (s *cursoRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var registro []models.CursoResp
	var totalCount int64

	// 1. Procesar el filtro para separar offset/limit de los WHERE
	if len(filter) != 0 {
		for key, value := range filter {
			if key == "offset" || key == "limit" {
				// Conversión segura a int
				var intVal int
				switch v := value.(type) {
				case int:
					intVal = v
				case float64:
					intVal = int(v)
				case string:
					parsed, err := strconv.Atoi(v)
					if err != nil {
						return nil, nil
					}
					intVal = parsed
				default:
					return nil, nil
				}

				// Asignación correcta
				switch key {
				case "offset":
					newSkip := intVal
					skip = &newSkip
				case "limit":
					newTake := intVal
					take = &newTake
				}

				delete(filter, key)
			}
		}
	}

	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	// 2. Consulta para contar el total (sin paginación)
	countQuery := DB.WithContext(ctx).Model(&models.CursoResp{})
	if len(filter) > 0 {
		for key, value := range filter {
			countQuery = countQuery.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}
	if err := countQuery.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	//asigna el schma a usar
	rsDB := infrastructure.GetDBWithSchema(ctx, s.DB)

	// Crea una consulta base
	query := rsDB.WithContext(ctx).Model(&models.CursoResp{})

	// Aplica filtros si existen
	if len(filter) != 0 {
		for key, value := range filter {
			query = query.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}
	// Verificar si skip y take son nil
	fmt.Println("skip:", skip, "take:", take)

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	if err := query.Find(&registro).Error; err != nil {
		return nil, err
	}

	//------------------------
	if len(registro) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Mapear a la estructura de respuesta
	response := models.CursoListResponse{
		Items:      registro,
		TotalCount: totalCount,
	}
	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil
}

func (s *cursoRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var curso models.CursoResp

	// Busca el registro en la base de datos utilizando GORM
	if err := DB.WithContext(ctx).Where("id = ?", ID).First(&curso).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &curso, nil
}

func (s *cursoRepository) Update(ctx context.Context, ID string, curso interface{}) error {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB).Begin()

	// rollback garantizado en caso de error o panic
	defer func() {
		if r := recover(); r != nil {
			DB.Rollback()
			//	panic(r) // puedes remover esto si no quieres propagar el panic
		} else if DB.Error != nil || DB.Statement != nil && DB.Statement.ConnPool != nil {
			_ = DB.Rollback() // rollback silencioso si no se hizo commit
		}
	}()

	if err := DB.WithContext(ctx).Model(&models.UpdateCursoReq{}).Where("id = ?", ID).Updates(curso).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			DB.Rollback()
			return wrappers.NewNonExistentErr(err)
		} else {
			DB.Rollback()
			return err
		}
	}

	// Confirma la transacción
	if err := DB.Commit().Error; err != nil {
		return errors.New("no se pudo confirmar la transacción: " + err.Error())
	}

	return nil
}

func (s *cursoRepository) Delete(ctx context.Context, ID string, filter map[string]interface{}) error {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB).Begin()

	// rollback garantizado en caso de error o panic
	defer func() {
		if r := recover(); r != nil {
			DB.Rollback()
			//	panic(r) // puedes remover esto si no quieres propagar el panic
		} else if DB.Error != nil || DB.Statement != nil && DB.Statement.ConnPool != nil {
			_ = DB.Rollback() // rollback silencioso si no se hizo commit
		}
	}()

	result := DB.WithContext(ctx).Delete(&models.Curso{}, "id = ?", ID)

	// Comprueba si ocurrió un error durante la eliminación
	if result.Error != nil {
		DB.Rollback()
		return wrappers.NewNonExistentErr(result.Error)
	}

	// Comprueba si se eliminaron filas
	if result.RowsAffected == 0 {
		DB.Rollback()
		return errors.New("pasajero no encontrado ") // Manejo de error si no se encontró el registro
	}

	// Confirma la transacción
	if err := DB.Commit().Error; err != nil {
		return errors.New("no se pudo confirmar la transacción: " + err.Error())
	}

	return nil
}
