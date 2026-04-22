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
type companyRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// GetUpdateByID implements ports.CompanyRepository.
func (s *companyRepository) GetUpdateByID(ctx context.Context, ID string) (interface{}, error) {
	panic("unimplemented")
}

// NewUserRepository creates a roles repository for postgres
func NewCompanyRepository(ctx context.Context, db *gorm.DB) ports.CompanyRepository {
	return &companyRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *companyRepository) Create(ctx context.Context, company interface{}) (string, error) {
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
	u, ok := company.(models.CreateCompanyReq)
	if !ok {
		tx.Rollback()
		return "", errors.New("el tipo de voucher no es correcto")
	}

	var existing models.CreateCompanyReq

	if err := tx.Where("rut = ?", u.Rut).First(&existing).Error; err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		tx.Rollback()
		return "error", errors.New("la compañia con el rut '" + u.Rut + "' ya existe")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar compañias " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	if err := tx.WithContext(ctx).Create(&u).Error; err != nil {
		tx.Rollback()
		return "", errors.New("error al crear compañia " + err.Error())
	}

	//crea el nuevo esquema y las tablas de trabajo
	/*
		if u.SchemaName != "" {
			if err := util.CreateNewSchema(u.SchemaName, tx); err != nil {
				log.Printf("Error al crear el esquema '%s': %v", u.SchemaName, err)
			}
		}
	*/

	// Confirma la transacción
	if err := tx.Commit().Error; err != nil {
		return "", errors.New("no se pudo confirmar la transacción: " + err.Error())
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *companyRepository) GetInf(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	return nil, nil
}

func (s *companyRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var registro []models.CompanyResp
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
	countQuery := DB.WithContext(ctx).Model(&models.CompanyResp{})
	if len(filter) > 0 {
		for key, value := range filter {
			countQuery = countQuery.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}
	if err := countQuery.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	// Crea una consulta base
	nDB := infrastructure.GetDBWithSchema(ctx, s.DB)

	query := nDB.WithContext(ctx).Model(&models.CompanyResp{})

	// Aplica filtros si existen
	if len(filter) != 0 {
		for key, value := range filter {
			query = query.Where(fmt.Sprintf("%s = ?", key), value)
		}
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
	response := models.CompanyListResponse{
		Items:      registro,
		TotalCount: totalCount,
	}

	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil
}

func (s *companyRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var company models.CompanyResp

	// Busca el registro en la base de datos utilizando GORM
	if err := DB.WithContext(ctx).Where("id = ?", ID).First(&company).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &company, nil
}

func (s *companyRepository) Update(ctx context.Context, ID string, company interface{}) error {
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

	if err := DB.WithContext(ctx).Model(&models.UpdateCompanyReq{}).Where("id = ?", ID).Updates(company).Error; err != nil {
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

func (s *companyRepository) Delete(ctx context.Context, ID string, filter map[string]interface{}) error {
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

	result := DB.WithContext(ctx).Delete(&models.Company{}, "id = ?", ID)

	// Comprueba si ocurrió un error durante la eliminación
	if result.Error != nil {
		DB.Rollback()
		return wrappers.NewNonExistentErr(result.Error)
	}

	// Comprueba si se eliminaron filas
	if result.RowsAffected == 0 {
		DB.Rollback()
		return errors.New("compañia no encontrado ") // Manejo de error si no se encontró el registro
	}

	// Confirma la transacción
	if err := DB.Commit().Error; err != nil {
		return errors.New("no se pudo confirmar la transacción: " + err.Error())
	}

	return nil
}
