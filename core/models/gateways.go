package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type GatewayAdditionalConfig struct {
	FlowAPIKey         string `json:"flow_apikey"`
	FlowSecretKey      string `json:"flow_secretkey"`
	TrbkCommercialCode string `json:"trbk_commercialcode"`
	TrbkKeySecret      string `json:"trbk_keysecret"`
	MpPublickey        string `json:"mp_publickey"`
	MpAccesstoken      string `json:"mp_accesstoken"`
	MpUsersid          string `json:"mp_usersid"`
}

func (c GatewayAdditionalConfig) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *GatewayAdditionalConfig) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, c)
}

type Gateways struct {
	ID               string                  `json:"_id,omitempty"`
	CompanyId        int64                   `json:"company_id"`
	GatewayId        int64                   `json:"gateway_id"`
	AdditionalConfig GatewayAdditionalConfig `gorm:"type:jsonb" json:"additional_config"` // GORM serializa automáticamente
	Active           int                     `json:"active"`
	CreatedDate      time.Time               `gorm:"autoCreateTime"`
	UpdatedDate      time.Time               `gorm:"autoUpdateTime"`
}

type GatewaysResp struct {
	ID               string                  `json:"id"`
	CompanyId        int64                   `json:"company_id"`
	GatewayId        int64                   `json:"gateway_id"`
	AdditionalConfig GatewayAdditionalConfig `gorm:"type:jsonb" json:"additional_config"` // GORM serializa automáticamente
	Active           int                     `json:"active"`
	CreatedDate      time.Time               `gorm:"autoCreateTime"`
	UpdatedDate      time.Time               `gorm:"autoUpdateTime"`
}

func (GatewaysResp) TableName() string {
	return "gateways" // Nombre de la tabla en la base de datos
}

type GatewaysListResponse struct {
	Items      []GatewaysResp `json:"items"`
	TotalCount int64          `json:"totalCount"`
}

type CreateGatewaysReq struct {
	ID               string                  `gorm:"primaryKey;autoIncrement"`
	CompanyId        int64                   `json:"company_id"`
	GatewayId        int64                   `json:"Gateway_id"`
	AdditionalConfig GatewayAdditionalConfig `json:"additional_config" gorm:"type:jsonb"` // GORM serializa automáticamente
	Active           int                     `json:"active"`
	CreatedDate      time.Time               `gorm:"autoCreateTime"`
	UpdatedDate      time.Time               `gorm:"autoUpdateTime"`
}

func (CreateGatewaysReq) TableName() string {
	return "gateways" // Nombre de la tabla en la base de datos
}

type GatewaysInforme struct {
	ID               string                  `json:"-"`
	CompanyId        int64                   `json:"company_id"`
	GatewayId        int64                   `json:"Gateway_id"`
	AdditionalConfig GatewayAdditionalConfig `json:"additional_config" gorm:"type:jsonb"` // GORM serializa automáticamente
	Active           int                     `json:"active"`
	CreatedDate      time.Time               `gorm:"autoCreateTime"`
	UpdatedDate      time.Time               `gorm:"autoUpdateTime"`
}

func (GatewaysInforme) TableName() string {
	return "gateways" // Nombre de la tabla en la base de datos
}
