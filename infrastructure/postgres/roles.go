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
type rolesRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewRolesRepository(ctx context.Context, db *gorm.DB) ports.RolesRepository {
	return &rolesRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *rolesRepository) Create(ctx context.Context, roles interface{}) (string, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)
	// Asegúrate de que el tipo del usuario es correcto
	u := roles.(models.CreateRolesReq)

	var existingReg models.CreateRolesReq

	err := DB.Where("description = ?", u.Description).First(&existingReg).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("El Role con la descripcion '" + u.Description + "' ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar roles: " + err.Error())
	}

	//asigna el schma a usar
	crDB := infrastructure.GetDBWithSchema(ctx, s.DB)
	// Usamos el contexto y creamos el registro en la base de datos
	if err := crDB.WithContext(ctx).Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *rolesRepository) GetInf(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	return nil, nil
}

func (s *rolesRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)
	var registro []models.RolesResp
	// Crea una consulta base
	query := DB.WithContext(ctx).Model(&models.RolesResp{})

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
	response := models.RolesListResponse{
		Items:      registro,
		TotalCount: 0,
	}

	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil
}

func (s *rolesRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var registro models.RolesResp

	// Busca el registro en la base de datos utilizando GORM
	if err := DB.WithContext(ctx).Where("id = ?", ID).First(&registro).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &registro, nil
}

func (s *rolesRepository) Update(ctx context.Context, ID string, roles interface{}) error {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	if err := DB.WithContext(ctx).Model(&models.UpdateRolesReq{}).Where("id = ?", ID).Updates(roles).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *rolesRepository) Delete(ctx context.Context, ID string, filter map[string]interface{}) error {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var roles models.CreateRolesReq

	result := DB.WithContext(ctx).Where("id = ?", ID).Delete(&roles)

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
