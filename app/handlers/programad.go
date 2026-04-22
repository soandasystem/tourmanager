package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"tourmanager/config"
	"tourmanager/core/models"
	"tourmanager/core/ports"

	//	"travel/tools/api/middlewares"
	"tourmanager/util"

	"github.com/gin-gonic/gin"
)

// SetProgramRoutes creates program routes
func SetProgramadRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.ProgramadService) {
	//	r.Use(middlewares.Recover())

	r.POST("/api/v3.5/programad", createProgramad(p))
	r.POST("/v2.0/programad/many", createManyProgramad(p))
	r.GET("/api/v3.5/programad", getAllProgramad(p))
	r.GET("/api/v3.5/programad/:id", getProgramadByID(p))
	r.PATCH("/api/v3.5/programad/:id", updateProgramad(p))
	r.DELETE("/api/v3.5/programad/:id", deleteProgramad(p))
}

// @Summary Create program
// @Description Creates a new program
// @Tags program
// @Param user body models.CreateprogramReq true "New program to be created"
// @Success 201 {object} models.programResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/program [post]
func createProgramad(p ports.ProgramadService) gin.HandlerFunc {
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
		var program models.CreateProgramadReq
		err = json.Unmarshal(body, &program)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		// Llamar al servicio para crear los program
		id, err := p.Create(ctx, program)
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

// @Summary Create program
// @Description Creates a new program
// @Tags program
// @Param user body models.CreateprogramReq true "New program to be created"
// @Success 201 {object} models.programResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /v2.0/program [post]
func createManyProgramad(p ports.ProgramadService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// El ctx ya tiene timeout + schema
		ctx := c.Request.Context()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		var users []models.CreateProgramadReq
		err = json.Unmarshal(body, &users)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		resultids, err := p.CreateMany(ctx, users)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		// Concatenar IDs en un string (puedes adaptar si necesitas otro formato)
		returnIDs := strings.Join(resultids.InsertedIDs, ", ")

		// Devolver la respuesta con el código de estado adecuado
		result := util.NewMessageResponse("Registro creados con ID(s): ", returnIDs)

		response := util.NewSuccessResponse(result, http.StatusAccepted)
		c.JSON(response.StatusCode, response)

	}
}

// @Summary Get all program
// @Description Gets all the program
// @Tags program
// @Success 200 {array} models.programResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/program [get]
func getAllProgramad(p ports.ProgramadService) gin.HandlerFunc {
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

// @Summary Get program by ID
// @Description Gets a program by ID
// @Tags program
// @Param id path string true "ID"
// @Success 200 {object} models.programResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/program/{id} [get]
func getProgramadByID(p ports.ProgramadService) gin.HandlerFunc {
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

// @Summary Update program
// @Description Updates a program
// @Tags program
// @Param id path string true "ID"
// @Param User body models.UpdateprogramReq true "program"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/program/{id} [patch]
func updateProgramad(p ports.ProgramadService) gin.HandlerFunc {
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
		var program models.UpdateProgramadReq
		err = json.Unmarshal(body, &program)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), program)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		result := util.NewMessageResponse("Registro con id "+id+" modificado ", id)

		response := util.NewSuccessResponse(result, http.StatusAccepted)
		c.JSON(response.StatusCode, response)
	}
}

// @Summary Delete program
// @Description Delete a program
// @Tags program
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/program/{id} [delete]
func deleteProgramad(p ports.ProgramadService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// El ctx ya tiene timeout + schema
		ctx := c.Request.Context()
		// Captura el parámetro de la URL
		programaID := c.Query("programa_id")

		filter := make(map[string]interface{})
		if programaID != "" {
			filter["programa_id"] = programaID
		}
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
