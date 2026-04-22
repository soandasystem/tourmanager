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

// SetUserRoutes creates Gateways routes
func SetGatewaysRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.GatewaysService) {
	//	r.Use(middlewares.Recover())

	r.POST("/api/v3.5/gateways", createGateways(p))
	r.GET("/api/v3.5/gateways", getAllGateways(p))
	r.GET("/api/v3.5/gateways/:id", getGatewaysByID(p))
	r.PATCH("/api/v3.5/gateways/:id", updateGateways())
	r.DELETE("/api/v3.5/gateways/:id", deleteGateways(p))
}

// @Summary Create Gateways
// @Description Creates a new Gateways
// @Tags colegion
// @Param user body models.CreateGatewaysReq true "New colegio to be created"
// @Success 201 {object} models.GatewaysResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/Gateways [post]
func createGateways(p ports.GatewaysService) gin.HandlerFunc {
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
		var Gateways models.CreateGatewaysReq
		err = json.Unmarshal(body, &Gateways)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		// Llamar al servicio para crear los Gateways
		id, err := p.Create(ctx, Gateways)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		// Devolver la respuesta con el código de estado adecuado
		result := util.NewMessageResponse("Registro Creado con id: "+id, id)

		response := util.NewSuccessResponse(result, http.StatusAccepted)
		c.JSON(response.StatusCode, response)
	}
}

// @Summary Get all Gateways
// @Description Gets all the Gateways
// @Tags Gateways
// @Success 200 {array} models.GatewaysResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/Gateways [get]
func getAllGateways(p ports.GatewaysService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// El ctx ya tiene timeout + schema
		ctx := c.Request.Context()

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
		// Agregar el total count al header (fuera del wrapper JSON)
		c.Header("X-Total-Count", strconv.FormatInt(result.TotalCount, 10))

		// Enviar solo los items en el wrapper de éxito
		response := util.NewSuccessResponse(result.Items, http.StatusOK)
		c.JSON(response.StatusCode, response)
	}
}

// @Summary Get Gateways by ID
// @Description Gets a Gateways by ID
// @Tags Gateways
// @Param id path string true "ID"
// @Success 200 {object} models.GatewaysResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/Gateways/{id} [get]
func getGatewaysByID(p ports.GatewaysService) gin.HandlerFunc {
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

// @Summary Update Gateways
// @Description Updates a Gateways
// @Tags Gateways
// @Param id path string true "ID"
// @Param User body models.UpdateGatewaysReq true "Gateways"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/users/{id} [patch]
func updateGateways() gin.HandlerFunc {
	return func(c *gin.Context) {

		result := util.NewMessageResponse("Registro ", "")

		response := util.NewSuccessResponse(result, http.StatusAccepted)
		c.JSON(response.StatusCode, response)
	}
}

// @Summary Delete Gateways
// @Description Delete a Gateways
// @Tags Gateways
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/Gateways/{id} [delete]
func deleteGateways(p ports.GatewaysService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// El ctx ya tiene timeout + schema
		ctx := c.Request.Context()

		//	var params = mux.Vars(c.Request)
		id := c.Param("id")
		err := p.Delete(ctx, c.Param("id"))
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
