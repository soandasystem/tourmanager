package handlers

import (
	"context"
	"net/http"
	"tourmanager/config"
	"tourmanager/core/ports"
	"tourmanager/util"

	"github.com/gin-gonic/gin"
)

// SetUploadRoutes creates upload routes
func SetUploadRoutes(ctx context.Context, cfg config.Config, r *gin.Engine, p ports.UploadService) {
	r.POST("/api/v3.5/upload", uploadFile(p))
}

// @Summary Upload file to B2
// @Description Uploads a file to Backblaze B2 storage
// @Tags upload
// @Accept multipart/form-data
// @Param file formData file true "File to upload"
// @Success 202 {object} object
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /api/v3.5/upload [post]
func uploadFile(p ports.UploadService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		// 1. Limitar el tamaño del archivo (ej. 10 MB) para proteger tu memoria
		err := c.Request.ParseMultipartForm(10 << 20)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusBadRequest)
			c.JSON(response.StatusCode, response)
			return
		}

		// 2. Obtener el archivo desde el "FormData" bajo la llave "file"
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusBadRequest)
			c.JSON(response.StatusCode, response)
			return
		}
		defer file.Close()

		contentType := header.Header.Get("Content-Type")

		// Llamar al servicio para subir el archivo
		publicURL, err := p.UploadFile(ctx, file, header.Filename, contentType)
		if err != nil {
			response := util.NewErrorResponse(err, http.StatusInternalServerError)
			c.JSON(response.StatusCode, response)
			return
		}

		result := map[string]string{
			"message": "Subido correctamente",
			"url":     publicURL,
		}

		response := util.NewSuccessResponse(result, http.StatusAccepted)
		c.JSON(response.StatusCode, response)
	}
}
