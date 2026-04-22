package models

import "time"

type Region struct {
	ID           string    `json:"_id,omitempty"`
	Code         string    `json:"code"`
	CodeInternal string    `json:"code_internal"`
	Description  string    `json:"description"`
	Position     int       `json:"position"`
	Active       int       `json:"active"`
	Author       string    `json:"author"`
	CompanyId    int64     `json:"company_id"`
	CreatedDate  time.Time `gorm:"autoCreateTime"`
	UpdatedDate  time.Time `gorm:"autoUpdateTime"`
}

// Resp  response struct
type RegionResp struct {
	ID           string    `json:"id"`
	Code         string    `json:"code"`
	CodeInternal string    `json:"code_internal"`
	Description  string    `json:"description"`
	Position     int       `json:"position"`
	Active       int       `json:"active"`
	Author       string    `json:"author"`
	CompanyId    int64     `json:"company_id"`
	CreatedDate  time.Time `gorm:"autoCreateTime"`
	UpdatedDate  time.Time `gorm:"autoUpdateTime"`
}

type RegionListResponse struct {
	Items      []RegionResp `json:"items"`
	TotalCount int64        `json:"totalCount"`
}

// Create---Req  request struct
type CreateRegionReq struct {
	ID           string    `gorm:"primaryKey;autoIncrement"`
	Code         string    `json:"code"`
	CodeInternal string    `json:"code_internal"`
	Description  string    `json:"description"`
	Position     int       `json:"position"`
	Active       int       `json:"active"`
	Author       string    `json:"author"`
	CompanyId    int64     `json:"company_id"`
	CreatedDate  time.Time `gorm:"autoCreateTime"`
	UpdatedDate  time.Time `gorm:"autoUpdateTime"`
}

type UpdateRegionReq struct {
	ID           string     `json:"-"`
	Code         *string    `json:"code"`
	CodeInternal *string    `json:"code_internal"`
	Description  *string    `json:"description"`
	Position     *int       `json:"position"`
	Active       *int       `json:"active"`
	Author       *string    `json:"author"`
	CompanyId    *int64     `json:"company_id"`
	CreatedDate  *time.Time `gorm:"autoCreateTime"`
	UpdatedDate  *time.Time `gorm:"autoUpdateTime"`
}

type RegionReport struct {
	ID           int64  `json:"id"`
	Code         string `json:"code"`
	CodeInternal string `json:"code_internal"`
	Description  string `json:"description"`
}

func (RegionReport) TableName() string {
	return "regions" // Nombre de la tabla en la base de datos
}
