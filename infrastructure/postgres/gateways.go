package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"tourmanager/core/models"
	"tourmanager/core/ports"

	"github.com/antoniomarfa/hexatools/infrastructure"
	"github.com/antoniomarfa/hexatools/wrappers"

	"gorm.io/gorm"
)

// userRepository adapter of an roles repository for postgres
type gatewaysRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewGatewaysRepository(ctx context.Context, db *gorm.DB) ports.GatewaysRepository {
	return &gatewaysRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *gatewaysRepository) Create(ctx context.Context, gateways interface{}) (string, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	// Asegúrate de que el tipo del usuario es correcto
	g := gateways.(models.CreateGatewaysReq)

	var existingReg models.CreateGatewaysReq

	err := DB.Where("gateway_id = ?", g.GatewayId).First(&existingReg).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("la pasarela con el codigo '" + strconv.FormatInt(g.GatewayId, 10) + "' ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar pasarelas: " + err.Error())
	}

	//asigna el schma a usar
	crDB := infrastructure.GetDBWithSchema(ctx, s.DB)
	// Usamos el contexto y creamos el registro en la base de datos
	if err := crDB.WithContext(ctx).Create(&g).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return g.ID, nil
}

func (s *gatewaysRepository) GetInf(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	return nil, nil
}

func (s *gatewaysRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var gateways []models.GatewaysResp

	// Crea una consulta base
	query := DB.WithContext(ctx).Model(&models.Gateways{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Find(&gateways).Error; err != nil {
		return nil, err
	}

	if len(gateways) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Mapear a la estructura de respuesta
	response := models.GatewaysListResponse{
		Items:      gateways,
		TotalCount: 0,
	}

	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil
}

func (s *gatewaysRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var registro models.GatewaysResp

	// Busca el registro en la base de datos utilizando GORM
	if err := DB.WithContext(ctx).Where("id = ?", ID).First(&registro).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &registro, nil
}

func (s *gatewaysRepository) GetByIDUpdate(ctx context.Context, ID string) (interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var registro models.Gateways

	// Busca el registro en la base de datos utilizando GORM
	if err := DB.WithContext(ctx).Where("id = ?", ID).First(&registro).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &registro, nil
}

func (s *gatewaysRepository) Update(ctx context.Context, ID string, gatewaysc interface{}) error {

	return nil
}

func (s *gatewaysRepository) Delete(ctx context.Context, ID string, filter map[string]interface{}) error {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var gateways models.Gateways

	result := DB.WithContext(ctx).Where("id = ?", ID).Delete(&gateways)

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
