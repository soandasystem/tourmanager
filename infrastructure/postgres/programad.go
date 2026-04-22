package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"tourmanager/core/entities"
	"tourmanager/core/models"
	"tourmanager/core/ports"

	"github.com/antoniomarfa/hexatools/infrastructure"
	"github.com/antoniomarfa/hexatools/wrappers"

	"gorm.io/gorm"
)

// userRepository adapter of an roles repository for postgres
type ProgramadRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewProgramadRepository(ctx context.Context, db *gorm.DB) ports.ProgramadRepository {
	return &ProgramadRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *ProgramadRepository) Create(ctx context.Context, program interface{}) (string, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)
	// Asegúrate de que el tipo del usuario es correcto
	u := program.(models.CreateProgramadReq)

	// Usamos el contexto y creamos el registro en la base de datos
	if err := DB.WithContext(ctx).Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *ProgramadRepository) GetInf(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	return nil, nil
}

func (s *ProgramadRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var registro []models.ProgramadResp

	// Crea una consulta base
	query := DB.WithContext(ctx).Model(&models.ProgramadResp{})

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
		return []interface{}{}, err
	}

	if len(registro) < 1 {
		return []interface{}{}, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Mapear a la estructura de respuesta
	response := models.ProgramadListResponse{
		Items:      registro,
		TotalCount: 0,
	}

	// Devolver como slice de interface{} para cumplir con la interfaz
	return []interface{}{response}, nil
}

func (s *ProgramadRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var registro models.ProgramadResp

	// Busca el registro en la base de datos utilizando GORM
	if err := DB.WithContext(ctx).Where("id = ?", ID).First(&registro).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &registro, nil
}

func (s *ProgramadRepository) Update(ctx context.Context, ID string, program interface{}) error {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	if err := DB.WithContext(ctx).Model(&models.UpdateProgramadReq{}).Where("id = ?", ID).Updates(program).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *ProgramadRepository) Delete(ctx context.Context, ID string, filter map[string]interface{}) error {
	//asigna el schma a usar
	DB := infrastructure.GetDBWithSchema(ctx, s.DB)

	var program models.CreateProgramadReq

	if len(filter) == 0 {
		// Aquí puedes manejar el caso donde el filtro está vacío
		result := DB.WithContext(ctx).Where("programa_id = ?", ID).Delete(&program)
		// Comprueba si ocurrió un error durante la eliminación
		if result.Error != nil {
			return wrappers.NewNonExistentErr(result.Error)
		}

		// Comprueba si se eliminaron filas
		if result.RowsAffected == 0 {
			return wrappers.NewNonExistentErr(sql.ErrNoRows) // Manejo de error si no se encontró el registro
		}
	} else {
		// Crea una consulta base
		query := DB.WithContext(ctx).Model(&models.ProgramadResp{})

		// Aquí puedes manejar el caso donde el filtro tiene contenido
		for key, value := range filter {
			query = query.Where(fmt.Sprintf("%s = ?", key), value)
		}
		result := query.Where(filter).Delete(nil)
		if result.Error != nil {
			// Manejo de error
			return wrappers.NewNonExistentErr(result.Error)
		} else {
			return nil
		}
	}
	return nil
}

func (s *ProgramadRepository) CreateMany(ctx context.Context, programad []interface{}) ([]string, error) {
	// `tx` is an instance of `*sql.Tx` through which we can execute our queries
	tx := s.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []string
	for _, entity := range programad {
		// Verifica y convierte a entities.Programad
		u, ok := entity.(entities.Programad)
		if !ok {
			tx.Rollback()
			return nil, fmt.Errorf("invalid type: expected entities.Programad")
		}

		// Inserta el registro en la base de datos
		if err := tx.Table("programad").Create(&u).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		result = append(result, u.ID)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return result, nil
}
