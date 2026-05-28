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

// SetPaymentInstallmentRoutes creates payment installment routes
func SetPaymentInstallmentsRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.PaymentInstallmentService) {
	//	r.Use(middlewares.Recover())

	r.POST("/api/v3.5/payment_installment", createPaymentInstallment(ctx, cfg, p))
	r.GET("/api/v3.5/payment_installment", getAllPaymentInstallment(ctx, cfg, p))
	r.GET("/api/v3.5/payment_installment/informe", getInfPaymentInstallment(ctx, cfg, p))
	r.GET("/api/v3.5/payment_installment/:id", getPaymentInstallmentByID(ctx, cfg, p))
	r.PATCH("/api/v3.5/payment_installment/:id", updatePaymentInstallment(ctx, cfg, p))
	r.DELETE("/api/v3.5/payment_installment/:id", deletePaymentInstallment(ctx, cfg, p))
}

// @Summary Create payment installment
// @Description Creates a new payment installment
// @Tags payment_installment
// @Param user body models.CreatePaymentInstallmentReq true "New payment installment to be created"
// @Success 201 {object} models.PaymentInstallmentResp "OK"
// @Failure 400 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/payment_installment [post]
func createPaymentInstallment(ctx context.Context, cfg config.Config, p ports.PaymentInstallmentService) gin.HandlerFunc {
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
		var ingreso models.CreatePaymentInstallmentReq
		err = json.Unmarshal(body, &ingreso)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		// Llamar al servicio para crear los ingreso
		id, err := p.Create(ctx, ingreso)
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

// @Summary Get inf all payment_installment
// @Description Gets all the payment_installment
// @Tags payment_installment
// @Success 200 {array} models.PaymentInstallmentResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/payment_installment/informe [get]
func getInfPaymentInstallment(ctx context.Context, cfg config.Config, p ports.PaymentInstallmentService) gin.HandlerFunc {
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
		result, err := p.GetInforme(ctx, filter)
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

// @Summary Get all payment_installment
// @Description Gets all the payment_installment
// @Tags payment_installment
// @Success 200 {array} models.PaymentInstallmentResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/payment_installment [get]
func getAllPaymentInstallment(ctx context.Context, cfg config.Config, p ports.PaymentInstallmentService) gin.HandlerFunc {
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

// @Summary Get payment_installment by ID
// @Description Gets a payment_installment by ID
// @Tags payment_installment
// @Param id path string true "ID"
// @Success 200 {object} models.PaymentInstallmentResp "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/payment_installment/{id} [get]
func getPaymentInstallmentByID(ctx context.Context, cfg config.Config, p ports.PaymentInstallmentService) gin.HandlerFunc {
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
		//	utils.ResponseJSON(c.Writer, c.Request, nil, http.StatusAccepted, users)
		response := util.NewSuccessResponse(result, http.StatusOK)
		c.JSON(response.StatusCode, response)
	}
}

// @Summary Update payment_installment
// @Description Updates a payment_installment
// @Tags payment_installment
// @Param id path string true "ID"
// @Param User body models.UpdatePaymentInstallmentReq true "payment_installment"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/payment_installment/{id} [patch]
func updatePaymentInstallment(ctx context.Context, cfg config.Config, p ports.PaymentInstallmentService) gin.HandlerFunc {
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
		var ingreso models.UpdatePaymentInstallmentReq
		err = json.Unmarshal(body, &ingreso)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}
		id := params.ByName("id")
		err = p.Update(ctx, params.ByName("id"), ingreso)
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

// @Summary Delete payment_installment
// @Description Delete a payment_installment
// @Tags payment_installment
// @Param id path string true "ID"
// @Success 200 "OK"
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 408 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/payment_installment/{id} [delete]
func deletePaymentInstallment(ctx context.Context, cfg config.Config, p ports.PaymentInstallmentService) gin.HandlerFunc {
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
