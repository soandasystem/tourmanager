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
type permissionRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewPermissionRepository(ctx context.Context, db *gorm.DB) ports.PermissionRepository {
	return &permissionRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *permissionRepository) Create(ctx context.Context, permission interface{}) (string, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)
	// Asegúrate de que el tipo del usuario es correcto
	u := permission.(models.CreateRolesPermissionsReq)

	// Usamos el contexto y creamos el registro en la base de datos
	if err := DB.WithContext(ctx).Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *permissionRepository) GetInf(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	return nil, nil
}

func (s *permissionRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var registro []models.RolesPermissionsResp
	//   var consulta:=""
	//   var userID:=""
	// Crea una consulta base
	query := DB.WithContext(ctx).Model(&models.RolesPermissionsResp{})

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
	response := models.RolesPermissionsListResponse{
		Items:      registro,
		TotalCount: 0,
	}

	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil
}

func (s *permissionRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var registro models.RolesPermissionsResp

	// Busca el registro en la base de datos utilizando GORM
	if err := DB.WithContext(ctx).Table("roles_permission").Where("id = ?", ID).First(&registro).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &registro, nil
}

func (s *permissionRepository) Update(ctx context.Context, ID string, permission interface{}) error {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	if err := DB.WithContext(ctx).Model(&models.UpdateRolesPermissionsReq{}).Where("id = ?", ID).Updates(permission).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *permissionRepository) Delete(ctx context.Context, ID string, filter map[string]interface{}) error {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var permission models.CreateRolesPermissionsReq

	// Aquí puedes manejar el caso donde el filtro está vacío
	result := DB.WithContext(ctx).Where("roles_id = ?", ID).Delete(&permission)
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
