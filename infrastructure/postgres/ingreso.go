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
type ingresoRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewIngresoRepository(ctx context.Context, db *gorm.DB) ports.IngresoRepository {
	return &ingresoRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *ingresoRepository) Create(ctx context.Context, ingreso interface{}) (string, error) {
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
	u := ingreso.(models.CreateIngresoReq)

	var existingIngreso models.CreateIngresoReq

	if err := tx.Where("identificador = ?", u.Identificador).First(&existingIngreso).Error; err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		tx.Rollback()
		return "error", errors.New("El Ingreso con el identificador '" + u.Identificador + "' ya existe")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar pagos " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	if err := tx.WithContext(ctx).Create(&u).Error; err != nil {
		tx.Rollback()
		return "", errors.New("error al crear el ingreso " + err.Error())

	}

	// Confirma la transacción
	if err := tx.Commit().Error; err != nil {
		return "", errors.New("no se pudo confirmar la transacción: " + err.Error())
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *ingresoRepository) GetInf(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var registro []models.IngresoInf
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

	// Crea una consulta base
	query := DB.WithContext(ctx).Model(&models.IngresoInf{})
	// Aplica filtros si existen

	if len(filter) != 0 {
		for key, value := range filter {
			switch key {
			case "start_date":
				query = query.Where("fecha >= ?", value)
			case "end_date":
				query = query.Where("fecha <= ?", value)
			default:
				query = query.Where(fmt.Sprintf("%s = ?", key), value)
			}
		}
	}

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	// Verificar si skip y take son nil
	fmt.Println("skip:", skip, "take:", take)

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	if err := query.Preload("Pago").Preload("Sale").Preload("Curso").Order("id ASC").Find(&registro).Error; err != nil {
		return nil, err
	}

	//------------------------
	if len(registro) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Mapear a la estructura de respuesta
	response := models.IngresoInfListResponse{
		Items:      registro,
		TotalCount: totalCount,
	}
	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil
}

func (s *ingresoRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var registro []models.IngresoResp
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

	// Crea una consulta base
	query := DB.WithContext(ctx).Model(&models.IngresoResp{})
	// Aplica filtros si existen
	if len(filter) != 0 {
		for key, value := range filter {
			switch key {
			case "start_date":
				query = query.Where("fecha >= ?", value)
			case "end_date":
				query = query.Where("fecha <= ?", value)
			default:
				query = query.Where(fmt.Sprintf("%s = ?", key), value)
			}

		}
	}

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	// Verificar si skip y take son nil
	fmt.Println("skip:", skip, "take:", take)

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
	response := models.IngresoListResponse{
		Items:      registro,
		TotalCount: totalCount,
	}
	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil
}

func (s *ingresoRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var ingreso models.IngresoResp

	// Busca el registro en la base de datos utilizando GORM
	if err := DB.WithContext(ctx).Where("id = ?", ID).First(&ingreso).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &ingreso, nil
}

func (s *ingresoRepository) Update(ctx context.Context, ID string, ingreso interface{}) error {
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

	if err := DB.WithContext(ctx).Model(&models.UpdateIngresoReq{}).Where("id = ?", ID).Updates(ingreso).Error; err != nil {
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

func (s *ingresoRepository) Delete(ctx context.Context, ID string, filter map[string]interface{}) error {
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

	result := DB.WithContext(ctx).Delete(&models.Ingreso{}, "id = ?", ID)

	// Comprueba si ocurrió un error durante la eliminación
	if result.Error != nil {
		DB.Rollback()
		return wrappers.NewNonExistentErr(result.Error)
	}

	// Comprueba si se eliminaron filas
	if result.RowsAffected == 0 {
		DB.Rollback()
		return errors.New("ingreso no encontrado ") // Manejo de error si no se encontró el registro
	}

	return nil
}
