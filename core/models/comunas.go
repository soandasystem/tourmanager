package models

import "time"

type Comunas struct {
	ID          string    `json:"_id,omitempty"`
	RegionsId   int8      `json:"regions_id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Active      int       `json:"active"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
	Author      string    `json:"author"`
}

// Resp  response struct
type ComunasResp struct {
	ID          string    `json:"id"`
	RegionsId   int8      `json:"regions_id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Active      int       `json:"active"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
	Author      string    `json:"author"`
}

type ComunasListResponse struct {
	Items      []ComunasResp `json:"items"`
	TotalCount int64         `json:"totalCount"`
}

// Create---Req  request struct
type CreateComunasReq struct {
	ID          string    `gorm:"primaryKey;autoIncrement"`
	RegionsId   int8      `json:"regions_id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Active      int       `json:"active"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
	Author      string    `json:"author"`
}

type UpdateComunasReq struct {
	ID          string     `json:"-"`
	RegionsId   *int8      `json:"regions_id"`
	Code        *string    `json:"code"`
	Description *string    `json:"description"`
	Active      *int       `json:"active"`
	CreatedDate *time.Time `gorm:"autoCreateTime"`
	UpdatedDate *time.Time `gorm:"autoUpdateTime"`
	Author      *string    `json:"author"`
}

type ComunasReport struct {
	ID          int64  `json:"id"`
	RegionsId   int8   `json:"regions_id"`
	Description string `json:"description"`
}

func (ComunasReport) TableName() string {
	return "communes" // Nombre de la tabla en la base de datos
}
