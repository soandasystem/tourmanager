package util

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SchemaMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		schema := c.GetHeader("X-Tenant-Schema")
		if schema == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "schema header required"})
			return
		}
		c.Set("schema", schema)
		c.Next()
	}
}

func ContextWithSchemaMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Leer el schema del header o de c.Get (si ya viene de otro middleware)
		schemaVal, exists := c.Get("schema")
		if !exists {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "schema not found in context"})
			return
		}

		schema, ok := schemaVal.(string)
		if !ok || schema == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid schema format"})
			return
		}

		// Crear un contexto nuevo con timeout y con schema
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		// Muy importante cancelar después de terminar el request
		c.Set("cancelFunc", cancel)

		ctxWithSchema := context.WithValue(ctx, "schema", schema)

		// Reemplazar el contexto original de la request con el nuevo
		c.Request = c.Request.WithContext(ctxWithSchema)

		// Continuar
		c.Next()

		// Cancelar al final de la petición
		cancel()
	}
}

func CreateNewSchema(newSchema string, db *gorm.DB) error {

	// Crear el nuevo esquema
	err := db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", newSchema)).Error
	if err != nil {
		//	log.Fatalf("Error al crear esquema: %v", err)
		return fmt.Errorf("Error al crear esquema: %v", err)
	}

	// Obtener lista de tablas (esto puede variar según GORM)
	var tables []string
	err = db.Table("information_schema.tables").
		Where("table_schema = ? AND table_type = ?", "template", "BASE TABLE").
		Pluck("table_name", &tables).Error
	if err != nil {
		//	log.Fatalf("Error al obtener tablas: %v", err)
		return fmt.Errorf("Error al obtener tablas: %v", err)
	}

	// Copiar cada tabla
	for _, table := range tables {
		err = db.Exec(fmt.Sprintf(`
			CREATE TABLE %s.%s AS 
			SELECT * FROM template.%s
		`, newSchema, table, table)).Error
		if err != nil {
			log.Printf("Error al copiar tabla %s: %v", table, err)
			continue
		}
		fmt.Printf("Tabla %s copiada exitosamente", table)
	}
	return nil
}
