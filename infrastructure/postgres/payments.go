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

// paymentsRepository adapter of an payments repository for postgres
type paymentsRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewPaymentsRepository creates a payments repository for postgres
func NewPaymentsRepository(ctx context.Context, db *gorm.DB) ports.PaymentRepository {
	return &paymentsRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *paymentsRepository) Create(ctx context.Context, payments interface{}) (string, error) {
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
	u, ok := payments.(models.CreatePaymentReq)
	if !ok {
		tx.Rollback()
		return "", errors.New("el tipo de pagos no es correcto")
	}
	var existingPagos models.CreatePaymentReq

	if err := tx.Where("id = ?", u.ID).First(&existingPagos).Error; err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		tx.Rollback()
		return "error", errors.New("El Ingreso con el identificador '" + u.ID + "' ya existe")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar pagos " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	if err := tx.WithContext(ctx).Create(&u).Error; err != nil {
		tx.Rollback()
		return "", errors.New("error al crear el pago " + err.Error())
	}

	// Confirma la transacción
	if err := tx.Commit().Error; err != nil {
		return "", errors.New("no se pudo confirmar la transacción: " + err.Error())
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *paymentsRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var registro []models.PaymentResp
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
	query := DB.WithContext(ctx).Model(&models.PaymentResp{})

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

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Table("pagos").Order("id ASC").Find(&registro).Error; err != nil {
		return nil, err
	}

	if len(registro) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Mapear a la estructura de respuesta
	response := models.PaymentListResponse{
		Items:      registro,
		TotalCount: totalCount,
	}
	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil
}

func (s *paymentsRepository) GetInf(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var registro []models.PaymentInf
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
	query := DB.WithContext(ctx).Model(&models.PaymentInf{})

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
	response := models.PaymentInfListResponse{
		Items:      registro,
		TotalCount: totalCount,
	}
	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil
}

func (s *paymentsRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var pagos models.PaymentResp

	// Busca el registro en la base de datos utilizando GORM
	if err := DB.WithContext(ctx).Where("id = ?", ID).First(&pagos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &pagos, nil
}

func (s *paymentsRepository) Update(ctx context.Context, ID string, payments interface{}) error {
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

	if err := DB.WithContext(ctx).Model(&models.UpdatePaymentReq{}).Where("id = ?", ID).Updates(payments).Error; err != nil {
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

func (s *paymentsRepository) Delete(ctx context.Context, ID string, filter map[string]interface{}) error {
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

	result := DB.WithContext(ctx).Delete(&models.Pagos{}, "id = ?", ID)

	// Comprueba si ocurrió un error durante la eliminación
	if result.Error != nil {
		DB.Rollback()
		return wrappers.NewNonExistentErr(result.Error)
	}

	// Comprueba si se eliminaron filas
	if result.RowsAffected == 0 {
		DB.Rollback()
		return errors.New("pago no encontrado ") // Manejo de error si no se encontró el registro	}
	}

	// Confirma la transacción
	if err := DB.Commit().Error; err != nil {
		return errors.New("no se pudo confirmar la transacción: " + err.Error())
	}
	return nil
}
