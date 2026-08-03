package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"tourmanager/config"
	"tourmanager/core/models"
	"tourmanager/core/ports"
	"tourmanager/util"

	"github.com/gin-gonic/gin"
)

// SetFlowRoutes creates flow payment routes
func SetFlowRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, s ports.FlowService) {
	r.POST("/api/v3.5/flow/init", initFlowPayment(ctx, cfg, s))
}

// @Summary Init Flow Payment
// @Description Inicializa el pago a través de Flow y devuelve la URL de redirección
// @Tags flow
// @Param request body models.InitFlowPaymentReq true "Configuración inicial del pago"
// @Success 200 {object} models.InitFlowPaymentResp "OK"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/gateways/flow/init [post]
func initFlowPayment(ctx context.Context, cfg config.Config, s ports.FlowService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		var req models.InitFlowPaymentReq
		if err := json.Unmarshal(body, &req); err != nil {
			response := util.NewErrorResponse(err, http.StatusBadRequest)
			c.JSON(response.StatusCode, response)
			return
		}

		resp, err := s.InitPayment(ctx, req)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}
