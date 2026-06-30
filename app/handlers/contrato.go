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

// SetContratoRoutes creates contrato routes
func SetContratoRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.ContratoService) {
	r.POST("/api/v3.5/contrato", generarContrato(p))
	r.POST("/api/v3.5/contrato/firma", firmarContrato(p))
}

// @Summary Create contrato DOCX temporal
// @Description Generates a temporary DOCX with the replaced placeholders from the request
// @Tags contrato
// @Param request body models.ContratoReq true "Contrato data"
// @Success 202 {object} util.ApiResponse "Accepted"
// @Failure 400 {object} util.ApiResponse
// @Failure 500 {object} util.ApiResponse
// @Router /api/v3.5/contrato [post]
func generarContrato(p ports.ContratoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		var req models.ContratoReq
		if err := json.Unmarshal(body, &req); err != nil {
			response := util.NewErrorResponse(err, http.StatusBadRequest)
			c.JSON(response.StatusCode, response)
			return
		}

		res, err := p.GenerarContrato(ctx, req)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		response := util.NewSuccessResponse(res, http.StatusAccepted)
		c.JSON(response.StatusCode, response)
	}
}

// @Summary Create contrato PDF definitivo con firma
// @Description Creates the final PDF by filling the DOCX data and injecting the signature
// @Tags contrato
// @Param request body models.ContratoFirmaReq true "Firma request"
// @Success 202 {object} util.ApiResponse "Accepted"
// @Failure 400 {object} util.ApiResponse
// @Failure 500 {object} util.ApiResponse
// @Router /api/v3.5/contrato/firma [post]
func firmarContrato(p ports.ContratoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		var req models.ContratoFirmaReq
		if err := json.Unmarshal(body, &req); err != nil {
			response := util.NewErrorResponse(err, http.StatusBadRequest)
			c.JSON(response.StatusCode, response)
			return
		}

		res, err := p.FirmarContrato(ctx, req)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		response := util.NewSuccessResponse(res, http.StatusAccepted)
		c.JSON(response.StatusCode, response)
	}
}
