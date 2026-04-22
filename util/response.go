package util

type ApiResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"` // Opcional: mensaje descriptivo
	Data       interface{} `json:"data,omitempty"`    // Datos principales (ej: items)
	Error      string      `json:"error,omitempty"`   // Error legible (si Success = false)
	StatusCode int         `json:"-"`                 // Usado internamente (no se serializa a JSON)
}

type ApiMessage struct {
	ReturnID string `json:"return_id"`
	Message  string `json:"message,omitempty"`
}

func NewSuccessResponse(data interface{}, statusCode int) *ApiResponse {
	return &ApiResponse{
		Success:    true,
		Data:       data,
		StatusCode: statusCode,
	}
}

func NewMessageResponse(message string, returnid string) *ApiMessage {
	return &ApiMessage{
		ReturnID: returnid,
		Message:  message,
	}
}

func NewErrorResponse(err error, statusCode int) *ApiResponse {
	return &ApiResponse{
		Success:    false,
		Error:      err.Error(),
		StatusCode: statusCode,
	}
}
