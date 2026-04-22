package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"tourmanager/config"
	"tourmanager/core/models"
	"tourmanager/core/ports"

	"tourmanager/util"

	"github.com/gin-gonic/gin"
)

// SetPermissionRoutes creates permission routes
func SetPermissionRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.PermissionService) {
	//	r.Use(middlewares.Recover())

	r.POST("/api/v3.5/permission", createPermission(p))
	r.GET("/api/v3.5/permission", getAllPermission(p))
	r.GET("/api/v3.5/permission/:id", getPermissionByID(p))
	r.PATCH("/api/v3.5/permission/:id", updatePermission(p))
	r.DELETE("/api/v3.5/permission/:id", deletePermission(p))
}

// @Summary Create permission
// @Description Creates a new permission
// @Tags permission
// @Param user body models.CreatepermissionReq true "New permission to be created"
// @Success 201 {object} models.permissionResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/permission [post]
func createPermission(p ports.PermissionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// El ctx ya tiene timeout + schema
		ctx := c.Request.Context()

		// Leer el cuerpo de la solicitud
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		// Parsear el JSON al modelo adecuado
		var permission models.CreateRolesPermissionsReq
		err = json.Unmarshal(body, &permission)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		// Llamar al servicio para crear los permission
		id, err := p.Create(ctx, permission)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		// Devolver la respuesta con el código de estado adecuado
		// Devolver la respuesta con el código de estado adecuado
		result := util.NewMessageResponse("Registro Creado con id: "+id, id)

		response := util.NewSuccessResponse(result, http.StatusAccepted)
		c.JSON(response.StatusCode, response)
	}
}

// @Summary Get all permission
// @Description Gets all the permission
// @Tags permission
// @Success 200 {array} models.permissionResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/permission [get]
func getAllPermission(p ports.PermissionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// El ctx ya tiene timeout + schema
		ctx := c.Request.Context()

		// Captura el parámetro de la UR
		filter := make(map[string]interface{})
		// Obtener todos los parámetros de consulta
		queryParams := c.Request.URL.Query()
		// Convertir a un map[string]interface{}
		for key, values := range queryParams {
			if len(values) > 1 {
				// Si hay múltiples valores, los almacenamos como slice
				filter[key] = values
			} else {
				// Si hay un solo valor, lo almacenamos directamente
				filter[key] = values[0]
			}
		}
		result, err := p.GetAll(ctx, filter)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		c.Header("X-Total-Count", strconv.FormatInt(result.TotalCount, 10))

		// Enviar solo los items en el wrapper de éxito
		response := util.NewSuccessResponse(result.Items, http.StatusOK)
		c.JSON(response.StatusCode, response)
	}
}

// @Summary Get permission by ID
// @Description Gets a permission by ID
// @Tags permission
// @Param id path string true "ID"
// @Success 200 {object} models.permissionResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/permission/{id} [get]
func getPermissionByID(p ports.PermissionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// El ctx ya tiene timeout + schema
		ctx := c.Request.Context()

		//	var params = mux.Vars(c.Request)
		result, err := p.GetByID(ctx, c.Param("id"))
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		response := util.NewSuccessResponse(result, http.StatusOK)
		c.JSON(response.StatusCode, response)
	}
}

// @Summary Update permission
// @Description Updates a permission
// @Tags permission
// @Param id path string true "ID"
// @Param User body models.UpdatepermissionReq true "permission"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/permission/{id} [patch]
func updatePermission(p ports.PermissionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// El ctx ya tiene timeout + schema
		ctx := c.Request.Context()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		params := c.Params
		var permission models.UpdateRolesPermissionsReq
		err = json.Unmarshal(body, &permission)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), permission)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		result := util.NewMessageResponse("Registro con id "+id+" actualizado", id)

		response := util.NewSuccessResponse(result, http.StatusAccepted)
		c.JSON(response.StatusCode, response)
	}
}

// @Summary Delete permission
// @Description Delete a permission
// @Tags permission
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/permission/{id} [delete]
func deletePermission(p ports.PermissionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// El ctx ya tiene timeout + schema
		ctx := c.Request.Context()

		filter := make(map[string]interface{})

		//	var params = mux.Vars(c.Request)
		id := c.Param("id")
		err := p.Delete(ctx, c.Param("id"), filter)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		result := util.NewMessageResponse("Registro con id "+id+" eliminado ", id)

		response := util.NewSuccessResponse(result, http.StatusAccepted)
		c.JSON(response.StatusCode, response)
	}
}
