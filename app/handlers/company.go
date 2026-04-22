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
func SetCompanyRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.CompanyService) {
	//	r.Use(middlewares.Recover())

	r.POST("/api/v3.5/company", createCompany(p))
	r.GET("/api/v3.5/company", getAllCompany(p))
	r.GET("/api/v3.5/company/:id", getCompanyByID(p))
	r.PATCH("/api/v3.5/company/:id", updateCompany(p))
	r.DELETE("/api/v3.5/company/:id", deleteCompany(p))
}

// @Summary Create company
// @Description Creates a new company
// @Tags company
// @Param user body models.CreatecompanyReq true "New company to be created"
// @Success 201 {object} models.companyResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/company [post]
func createCompany(p ports.CompanyService) gin.HandlerFunc {
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
		var Company models.CreateCompanyReq
		err = json.Unmarshal(body, &Company)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		// Llamar al servicio para crear los company
		id, err := p.Create(ctx, Company)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		result := util.NewMessageResponse("Registro Creado con id: "+id, id)

		response := util.NewSuccessResponse(result, http.StatusAccepted)
		c.JSON(response.StatusCode, response)
	}
}

// @Summary Get all company
// @Description Gets all the company
// @Tags company
// @Success 200 {array} models.companyResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/company [get]
func getAllCompany(p ports.CompanyService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// El ctx ya tiene timeout + schema
		ctx := c.Request.Context()

		// Crea un mapa de filtro basado en el parámetro
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

		company, err := p.GetAll(ctx, filter)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		// Agregar el total count al header (fuera del wrapper JSON)
		c.Header("X-Total-Count", strconv.FormatInt(company.TotalCount, 10))

		// Enviar solo los items en el wrapper de éxito
		response := util.NewSuccessResponse(company.Items, http.StatusOK)
		c.JSON(response.StatusCode, response)
	}
}

// @Summary Get company by ID
// @Description Gets a company by ID
// @Tags company
// @Param id path string true "ID"
// @Success 200 {object} models.companyResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/company{id} [get]
func getCompanyByID(p ports.CompanyService) gin.HandlerFunc {
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

// @Summary Update company
// @Description Updates a company
// @Tags company
// @Param id path string true "ID"
// @Param User body models.UpdatecompanyReq true "company"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/company/{id} [patch]
func updateCompany(p ports.CompanyService) gin.HandlerFunc {
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
		var Company models.UpdateCompanyReq
		err = json.Unmarshal(body, &Company)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), Company)
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

// @Summary Delete company
// @Description Delete a company
// @Tags company
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/company/{id} [delete]
func deleteCompany(p ports.CompanyService) gin.HandlerFunc {
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
