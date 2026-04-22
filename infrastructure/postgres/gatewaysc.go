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
type gatewayscRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewGatewayscRepository(ctx context.Context, db *gorm.DB) ports.GatewayscRepository {
	return &gatewayscRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *gatewayscRepository) Create(ctx context.Context, gatewaysc interface{}) (string, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	// Asegúrate de que el tipo del usuario es correcto
	g := gatewaysc.(models.CreateGatewayscReq)

	var existingReg models.CreateGatewayscReq

	err := DB.Where("gateway_type = ?", g.GatewayType).First(&existingReg).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("la pasarela con el nombre '" + g.GatewayType + "' ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar pasarelas: " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	//asigna el schma a usar
	crDB := infrastructure.GetDBWithSchema(ctx, s.DB)
	if err := crDB.WithContext(ctx).Create(&g).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return g.ID, nil
}

func (s *gatewayscRepository) GetInf(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	return nil, nil
}

func (s *gatewayscRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {

	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var gatewaysc []models.GatewayscResp

	// Crea una consulta base
	query := DB.WithContext(ctx).Model(&models.GatewayscResp{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Find(&gatewaysc).Error; err != nil {
		return nil, err
	}

	if len(gatewaysc) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	response := models.GatewayscListResponse{
		Items:      gatewaysc,
		TotalCount: 0,
	}

	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil
}

func (s *gatewayscRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var gatewaysc models.GatewayscResp

	// Busca el registro en la base de datos utilizando GORM
	if err := DB.WithContext(ctx).Where("id = ?", ID).First(&gatewaysc).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &gatewaysc, nil
}

func (s *gatewayscRepository) GetByIDUpdate(ctx context.Context, ID string) (interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var registro models.GatewayscResp

	// Busca el registro en la base de datos utilizando GORM
	if err := DB.WithContext(ctx).Where("id = ?", ID).First(&registro).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &registro, nil
}

func (s *gatewayscRepository) Update(ctx context.Context, ID string, gatewaysc interface{}) error {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	if err := DB.WithContext(ctx).Model(&models.UpdateGatewayscReq{}).Where("id = ?", ID).Updates(gatewaysc).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *gatewayscRepository) Delete(ctx context.Context, ID string, filter map[string]interface{}) error {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var gatewaysc models.CreateGatewayscReq

	result := DB.WithContext(ctx).Where("id = ?", ID).Delete(&gatewaysc)

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
