package models

import "time"

type Gatewaysc struct {
	ID                 string    `json:"_id,omitempty"`
	GatewayType        string    `json:"gateway_type"`
	GatewayDescription string    `json:"gateway_description"`
	GatewayImage       string    `json:"gateway_image"`
	GatewayUrl         string    `json:"gateway_url"`
	Active             int       `json:"active"`
	CreatedDate        time.Time `gorm:"autoCreateTime"`
	UpdatedDate        time.Time `gorm:"autoUpdateTime"`
}

// Resp  response struct
type GatewayscResp struct {
	ID                 string    `json:"id"`
	GatewayType        string    `json:"gateway_type"`
	GatewayDescription string    `json:"gateway_description"`
	GatewayImage       string    `json:"gateway_image"`
	GatewayUrl         string    `json:"gateway_url"`
	Active             int       `json:"active"`
	CreatedDate        time.Time `gorm:"autoCreateTime"`
	UpdatedDate        time.Time `gorm:"autoUpdateTime"`
}

func (GatewayscResp) TableName() string {
	return "gatewaysc" // Nombre de la tabla en la base de datos
}

type GatewayscListResponse struct {
	Items      []GatewayscResp `json:"items"`
	TotalCount int64           `json:"totalCount"`
}

// Create---Req  request struct
type CreateGatewayscReq struct {
	ID                 string    `gorm:"primaryKey;autoIncrement"`
	GatewayType        string    `json:"gateway_type"`
	GatewayDescription string    `json:"gateway_description"`
	GatewayImage       string    `json:"gateway_image"`
	GatewayUrl         string    `json:"gateway_url"`
	Active             int       `json:"active"`
	CreatedDate        time.Time `gorm:"autoCreateTime"`
	UpdatedDate        time.Time `gorm:"autoUpdateTime"`
}

func (CreateGatewayscReq) TableName() string {
	return "gatewaysc" // Nombre de la tabla en la base de datos
}

type UpdateGatewayscReq struct {
	ID                 string     `json:"-"`
	GatewayType        *string    `json:"gateway_type"`
	GatewayDescription *string    `json:"gateway_description"`
	GatewayImage       *string    `json:"gateway_image"`
	GatewayUrl         *string    `json:"gateway_url"`
	Active             *int       `json:"active"`
	CreatedDate        *time.Time `gorm:"autoCreateTime"`
	UpdatedDate        *time.Time `gorm:"autoUpdateTime"`
}

func (UpdateGatewayscReq) TableName() string {
	return "gatewaysc" // Nombre de la tabla en la base de datos
}

type GatewayscReport struct {
	ID                 int64  `json:"id"`
	GatewayType        string `json:"gateway_type"`
	GatewayDescription string `json:"gateway_description"`
	GatewayImage       string `json:"gateway_image"`
	GatewayUrl         string `json:"gateway_url"`
}

func (GatewayscReport) TableName() string {
	return "gatewaysc" // Nombre de la tabla en la base de datos
}
