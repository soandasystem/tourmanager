package models

import (
	"time"

	"gorm.io/datatypes"
)

type Roles struct {
	ID          string         `json:"_id,omitempty"`
	Description string         `json:"description"`
	Active      int            `json:"active"`
	CreatedDate time.Time      `gorm:"autoCreateTime"`
	UpdatedDate time.Time      `gorm:"autoUpdateTime"`
	Author      string         `json:"author"`
	CompanyId   int8           `json:"company_id"`
	Permissions datatypes.JSON `json:"permissions" swaggertype:"object"`
}

// Resp  response struct
type RolesResp struct {
	ID          string         `json:"id"`
	Description string         `json:"description"`
	Active      int            `json:"active"`
	CreatedDate time.Time      `gorm:"autoCreateTime"`
	UpdatedDate time.Time      `gorm:"autoUpdateTime"`
	Author      string         `json:"author"`
	CompanyId   int8           `json:"company_id"`
	Permissions datatypes.JSON `json:"permissions" swaggertype:"object"`
}

func (RolesResp) TableName() string {
	return "roles" // Nombre de la tabla en la base de datos
}

type RolesListResponse struct {
	Items      []RolesResp `json:"items"`
	TotalCount int64       `json:"totalCount"`
}

// Create---Req  request struct
type CreateRolesReq struct {
	ID          string         `gorm:"primaryKey;autoIncrement"`
	Description string         `json:"description"`
	Active      int            `json:"active"`
	CreatedDate time.Time      `gorm:"autoCreateTime"`
	UpdatedDate time.Time      `gorm:"autoUpdateTime"`
	Author      string         `json:"author"`
	CompanyId   int8           `json:"company_id"`
	Permissions datatypes.JSON `json:"permissions" swaggertype:"object"`
}

func (CreateRolesReq) TableName() string {
	return "roles" // Nombre de la tabla en la base de datos
}

type UpdateRolesReq struct {
	ID          string          `json:"-"`
	Description *string         `json:"description"`
	Active      *int            `json:"active"`
	CreatedDate *time.Time      `gorm:"autoCreateTime"`
	UpdatedDate *time.Time      `gorm:"autoUpdateTime"`
	Author      *string         `json:"author"`
	CompanyId   *int8           `json:"company_id"`
	Permissions *datatypes.JSON `json:"permissions" swaggertype:"object"`
}

func (UpdateRolesReq) TableName() string {
	return "roles" // Nombre de la tabla en la base de datos
}

type RolesReport struct {
	ID          int64          `json:"id"`
	Description string         `json:"description"`
	CompanyId   int8           `json:"company_id"`
	Permissions datatypes.JSON `json:"permissions" swaggertype:"object"`
}

func (RolesReport) TableName() string {
	return "roles" // Nombre de la tabla en la base de datos
}
