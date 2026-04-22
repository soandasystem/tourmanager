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

// SetUserRoutes creates user routes
func SetComunasRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.ComunasService) {
	//	r.Use(middlewares.Recover())

	r.POST("/api/v3.5/comunas", createComunas(p))
	r.GET("/api/v3.5/comunas", getAllComunas(p))
	r.GET("/api/v3.5/comunas/:id", getComunasByID(p))
	r.PATCH("/api/v3.5/comunas/:id", updateComunas(p))
	r.DELETE("/api/v3.5/comunas/:id", deleteComunas(p))
}

// @Summary Create comunas
// @Description Creates a new comunas
// @Tags comunas
// @Param user body models.CreatecomunasReq true "New comunas to be created"
// @Success 201 {object} models.comunasResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/comunas [post]
func createComunas(p ports.ComunasService) gin.HandlerFunc {
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
		var Comunas models.CreateComunasReq
		err = json.Unmarshal(body, &Comunas)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		// Llamar al servicio para crear los comunas
		id, err := p.Create(ctx, Comunas)
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

// @Summary Get all comunas
// @Description Gets all the comunas
// @Tags comunas
// @Success 200 {array} models.comunasResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/comunas [get]
func getAllComunas(p ports.ComunasService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// El ctx ya tiene timeout + schema
		ctx := c.Request.Context()

		// Captura el parámetro de la URL
		regionsID := c.Query("regions_id")

		// Crea un mapa de filtro basado en el parámetro
		filter := make(map[string]interface{})
		if regionsID != "" {
			filter["regions_id"] = regionsID
		}

		// Llama al servicio con el filtro
		comunas, err := p.GetAll(ctx, filter)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		// Agregar el total count al header (fuera del wrapper JSON)
		c.Header("X-Total-Count", strconv.FormatInt(comunas.TotalCount, 10))

		// Enviar solo los items en el wrapper de éxito
		response := util.NewSuccessResponse(comunas.Items, http.StatusOK)
		c.JSON(response.StatusCode, response)
	}
}

// @Summary Get comunas by ID
// @Description Gets a comunas by ID
// @Tags comunas
// @Param id path string true "ID"
// @Success 200 {object} models.comunasResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/comunas{id} [get]
func getComunasByID(p ports.ComunasService) gin.HandlerFunc {
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

// @Summary Update comunas
// @Description Updates a comunas
// @Tags comunas
// @Param id path string true "ID"
// @Param User body models.UpdatecomunasReq true "comunas"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/comunas/{id} [patch]
func updateComunas(p ports.ComunasService) gin.HandlerFunc {
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
		var Comunas models.UpdateComunasReq
		err = json.Unmarshal(body, &Comunas)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), Comunas)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		result := util.NewMessageResponse("Registro con id "+id+" actualizado ", id)

		response := util.NewSuccessResponse(result, http.StatusAccepted)
		c.JSON(response.StatusCode, response)
	}
}

// @Summary Delete comunas
// @Description Delete a comunas
// @Tags comunas
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/comunas/{id} [delete]
func deleteComunas(p ports.ComunasService) gin.HandlerFunc {
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
