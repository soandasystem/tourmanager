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
type saleRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewSaleRepository(ctx context.Context, db *gorm.DB) ports.SaleRepository {
	return &saleRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *saleRepository) Create(ctx context.Context, sale interface{}) (string, error) {
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
	u, ok := sale.(models.CreateSaleReq)
	if !ok {
		tx.Rollback()
		return "", errors.New("el tipo de venta no es correcto")
	}

	if u.TypeSale == "GE" {
		var existingSale models.CreateSaleReq

		if err := tx.Where("establecimiento_id = ? and curso = ? and idcurso = ?", u.EstablecimientoId, u.Curso, u.Idcurso).First(&existingSale).Error; err == nil {
			// Si no hay error, significa que se encontró un rol con ese nombre
			tx.Rollback()
			return "error", errors.New("hay una venta para este colegio y curso")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			// Si el error no es de registro no encontrado, es un error inesperado
			return "error", errors.New("Error al buscar ventas " + err.Error())
		}
	}

	// Usamos el contexto y creamos el registro en la base de datos
	if err := tx.WithContext(ctx).Create(&u).Error; err != nil {
		tx.Rollback()
		return "", errors.New("error al crear venta " + err.Error())
	}

	// Confirma la transacción
	if err := tx.Commit().Error; err != nil {
		return "", errors.New("no se pudo confirmar la transacción: " + err.Error())
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *saleRepository) GetInf(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var registro []models.SaleInforme
	var totalCount int64

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
	// Crea una consulta base con los joins necesarios
	query := DB.WithContext(ctx).Model(&models.SaleInforme{}).
		Select("s.*, c.total_curso, e.nombre AS establecimiento_nombre, u.name AS seller_name, p.name AS program_name").
		Joins("LEFT JOIN (SELECT c.sale_id, COUNT(*) AS total_curso FROM cursos c GROUP BY c.sale_id) c ON c.sale_id = s.id").
		Joins("LEFT JOIN establecimientos e ON e.id = s.establecimiento_id").
		Joins("LEFT JOIN users u ON u.id = s.seller_id").
		Joins("LEFT JOIN programac p ON p.id = s.program_id").
		Table("sales s")

	// Aplica filtros si existen
	if len(filter) != 0 {
		for key, value := range filter {
			switch key {
			case "start_date":
				query = query.Where("s.fecha >= ?", value)
			case "end_date":
				query = query.Where("s.fecha <= ?", value)
			default:
				query = query.Where(fmt.Sprintf("s.%s = ?", key), value)
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
	response := models.SaleInfListResponse{
		Items:      registro,
		TotalCount: totalCount,
	}
	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil
}

func (s *saleRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var registro []models.SaleResp
	var totalCount int64

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
	query := DB.WithContext(ctx).Model(&models.SaleResp{})

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
	response := models.SaleListResponse{
		Items:      registro,
		TotalCount: totalCount,
	}
	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil
}

func (s *saleRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)
	var sale models.SaleResp

	// Busca el registro en la base de datos utilizando GORM
	if err := DB.WithContext(ctx).Where("id = ?", ID).First(&sale).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, err
	}

	return &sale, nil
}

func (s *saleRepository) Update(ctx context.Context, ID string, sale interface{}) error {
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

	if err := DB.WithContext(ctx).Model(&models.UpdateSaleReq{}).Where("id = ?", ID).Updates(sale).Error; err != nil {
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

func (s *saleRepository) Delete(ctx context.Context, ID string, filter map[string]interface{}) error {
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

	result := DB.WithContext(ctx).Delete(&models.Sale{}, "id = ?", ID)

	// Comprueba si ocurrió un error durante la eliminación
	if result.Error != nil {
		DB.Rollback()
		return wrappers.NewNonExistentErr(result.Error)
	}

	// Comprueba si se eliminaron filas
	if result.RowsAffected == 0 {
		DB.Rollback()
		return errors.New("venta no encontrada ") // Manejo de error si no se encontró el registro
	}

	// Confirma la transacción
	if err := DB.Commit().Error; err != nil {
		return errors.New("no se pudo confirmar la transacción: " + err.Error())
	}
	return nil
}
