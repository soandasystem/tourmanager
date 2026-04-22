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
func SetCountryRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.CountryService) {
	//	r.Use(middlewares.Recover())

	r.POST("/api/v3.5/country", createCountry(p))
	r.GET("/api/v3.5/country", getAllCountry(p))
	r.GET("/api/v3.5/country/:id", getCountryByID(p))
	r.PATCH("/api/v3.5/country/:id", updateCountry(p))
	r.DELETE("/api/v3.5/country/:id", deleteCountry(p))
}

// @Summary Create Country
// @Description Creates a new Country
// @Tags Country
// @Param user body models.CreateCountryReq true "New Country to be created"
// @Success 201 {object} models.CountryResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/Country [post]
func createCountry(p ports.CountryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Establecer un timeout en el contexto  c.Request.Context()
		ctx := c.Request.Context()

		// Leer el cuerpo de la solicitud
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		// Parsear el JSON al modelo adecuado
		var Country models.CreateCountryReq
		err = json.Unmarshal(body, &Country)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		// Llamar al servicio para crear los Country
		id, err := p.Create(ctx, Country)
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

// @Summary Get all Country
// @Description Gets all the Country
// @Tags Country
// @Success 200 {array} models.CountryResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/Country [get]
func getAllCountry(p ports.CountryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		// Captura el parámetro de la URL
		regionsID := c.Query("regions_id")

		// Crea un mapa de filtro basado en el parámetro
		filter := make(map[string]interface{})
		if regionsID != "" {
			filter["regions_id"] = regionsID
		}

		// Llama al servicio con el filtro
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

// @Summary Get Country by ID
// @Description Gets a Country by ID
// @Tags Country
// @Param id path string true "ID"
// @Success 200 {object} models.CountryResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/Country{id} [get]
func getCountryByID(p ports.CountryService) gin.HandlerFunc {
	return func(c *gin.Context) {
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

// @Summary Update Country
// @Description Updates a Country
// @Tags Country
// @Param id path string true "ID"
// @Param User body models.UpdateCountryReq true "Country"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/Country/{id} [patch]
func updateCountry(p ports.CountryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		params := c.Params
		var Country models.UpdateCountryReq
		err = json.Unmarshal(body, &Country)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), Country)
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

// @Summary Delete Country
// @Description Delete a Country
// @Tags Country
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/Country/{id} [delete]
func deleteCountry(p ports.CountryService) gin.HandlerFunc {
	return func(c *gin.Context) {
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
