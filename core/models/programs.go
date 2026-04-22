package models

import (
	"time"

	"gorm.io/datatypes"
)

type Programs struct {
	ID              string         `json:"_id,omitempty"`
	Code            string         `json:"code"`
	Name            string         `json:"name"`
	Reserva         float32        `json:"reserva"`
	Active          int            `json:"active"`
	Author          string         `json:"author"`
	CompanyId       int64          `json:"company_id"`
	OriginCode      string         `json:"origin_code"`
	DestinationCode string         `json:"destination_code"`
	Origin          string         `json:"origin"`
	Destination     string         `json:"destination"`
	Matrix          datatypes.JSON `json:"matrix" swaggertype:"object"`
	CreatedDate     time.Time      `gorm:"autoCreateTime"`
	UpdateDate      time.Time      `gorm:"autoUpdateTime"`
}

type ProgramsResp struct {
	ID              string         `json:"id"`
	Code            string         `json:"code"`
	Name            string         `json:"name"`
	Reserva         float32        `json:"reserva"`
	Active          int            `json:"active"`
	Author          string         `json:"author"`
	CompanyId       int64          `json:"company_id"`
	OriginCode      string         `json:"origin_code"`
	DestinationCode string         `json:"destination_code"`
	Origin          string         `json:"origin"`
	Destination     string         `json:"destination"`
	Matrix          datatypes.JSON `json:"matrix" swaggertype:"object"`
	CreatedDate     time.Time      `gorm:"autoCreateTime"`
	UpdateDate      time.Time      `gorm:"autoUpdateTime"`
}

type ProgramsListResponse struct {
	Items      []ProgramsResp `json:"items"`
	TotalCount int64          `json:"totalCount"`
}

func (ProgramsResp) TableName() string {
	return "programas" // Nombre de la tabla en la base de datos
}

type CreateProgramsReq struct {
	ID              string         `json:"id" gorm:"primaryKey;autoIncrement"`
	Code            string         `json:"code"`
	Name            string         `json:"name"`
	Reserva         float32        `json:"reserva"`
	Active          int            `json:"active"`
	Author          string         `json:"author"`
	CompanyId       int64          `json:"company_id"`
	OriginCode      string         `json:"origin_code"`
	DestinationCode string         `json:"destination_code"`
	Origin          string         `json:"origin"`
	Destination     string         `json:"destination"`
	Matrix          datatypes.JSON `json:"matrix" swaggertype:"object"`
	CreatedDate     *time.Time     `gorm:"autoCreateTime"`
	UpdateDate      *time.Time     `gorm:"autoUpdateTime"`
}

func (CreateProgramsReq) TableName() string {
	return "programas" // Nombre de la tabla en la base de datos
}

type UpdateProgramsReq struct {
	ID              string          `json:"-"`
	Code            *string         `json:"code"`
	Name            *string         `json:"name"`
	Reserva         *float32        `json:"reserva"`
	Active          *int            `json:"active"`
	Author          *string         `json:"author"`
	CompanyId       *int64          `json:"company_id"`
	OriginCode      *string         `json:"origin_code"`
	DestinationCode *string         `json:"destination_code"`
	Origin          *string         `json:"origin"`
	Destination     *string         `json:"destination"`
	Matrix          *datatypes.JSON `json:"matrix" swaggertype:"object"`
	CreatedDate     *time.Time      `gorm:"autoCreateTime"`
	UpdateDate      *time.Time      `gorm:"autoUpdateTime"`
}

func (UpdateProgramsReq) TableName() string {
	return "programas" // Nombre de la tabla en la base de datos
}

type ProgramsReport struct {
	ID              int64          `json:"id"`
	Code            string         `json:"code"`
	Name            string         `json:"name"`
	Active          int            `json:"active"`
	Reserva         int            `json:"reserva"`
	CompanyId       int64          `json:"company_id"`
	OriginCode      string         `json:"origin_code"`
	DestinationCode string         `json:"destination_code"`
	Origin          string         `json:"origin"`
	Destination     string         `json:"destination"`
	Matrix          datatypes.JSON `json:"matrix" swaggertype:"object"`
}

func (ProgramsReport) TableName() string {
	return "programs" // Nombre de la tabla en la base de datos
}
