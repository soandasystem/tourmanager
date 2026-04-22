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

// SetUserRoutes creates Gatewaysc routes
func SetGatewayscRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.GatewayscService) {
	//	r.Use(middlewares.Recover())

	r.POST("/api/v3.5/gatewaysc", createGatewaysc(p))
	r.GET("/api/v3.5/gatewaysc", getAllGatewaysc(p))
	r.GET("/api/v3.5/gatewaysc/:id", getGatewayscByID(p))
	r.PATCH("/api/v3.5/gatewaysc/:id", updateGatewaysc(p))
	r.DELETE("/api/v3.5/gatewaysc/:id", deleteGatewaysc(p))
}

// @Summary Create Gatewaysc
// @Description Creates a new Gatewaysc
// @Tags colegion
// @Param user body models.CreateGatewayscReq true "New colegio to be created"
// @Success 201 {object} models.GatewayscResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/Gatewaysc [post]
func createGatewaysc(p ports.GatewayscService) gin.HandlerFunc {
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
		var gatewaysc models.CreateGatewayscReq
		err = json.Unmarshal(body, &gatewaysc)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		// Llamar al servicio para crear los Gatewaysc
		id, err := p.Create(ctx, gatewaysc)
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

// @Summary Get all Gatewaysc
// @Description Gets all the Gatewaysc
// @Tags Gatewaysc
// @Success 200 {array} models.GatewayscResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/Gatewaysc [get]
func getAllGatewaysc(p ports.GatewayscService) gin.HandlerFunc {
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

// @Summary Get Gatewaysc by ID
// @Description Gets a Gatewaysc by ID
// @Tags Gatewaysc
// @Param id path string true "ID"
// @Success 200 {object} models.GatewayscResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/Gatewaysc/{id} [get]
func getGatewayscByID(p ports.GatewayscService) gin.HandlerFunc {
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

// @Summary Update Gatewaysc
// @Description Updates a Gatewaysc
// @Tags Gatewaysc
// @Param id path string true "ID"
// @Param User body models.UpdateGatewayscReq true "Gatewaysc"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/users/{id} [patch]
func updateGatewaysc(p ports.GatewayscService) gin.HandlerFunc {
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
		var Gatewaysc models.UpdateGatewayscReq
		err = json.Unmarshal(body, &Gatewaysc)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), Gatewaysc)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		result := util.NewMessageResponse("Registro con "+id+" actualizado ", id)

		response := util.NewSuccessResponse(result, http.StatusAccepted)
		c.JSON(response.StatusCode, response)
	}
}

// @Summary Delete Gatewaysc
// @Description Delete a Gatewaysc
// @Tags Gatewaysc
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/Gatewaysc/{id} [delete]
func deleteGatewaysc(p ports.GatewayscService) gin.HandlerFunc {
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
