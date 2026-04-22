package models

import "time"

type Programac struct {
	ID          string    `json:"_id,omitempty"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Reserva     float32   `json:"reserva"`
	Active      int       `json:"active"`
	Author      string    `json:"author"`
	CompanyId   int64     `json:"company_id"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdateDate  time.Time `gorm:"autoUpdateTime"`
}

type ProgramacResp struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Reserva     float32   `json:"reserva"`
	Active      int       `json:"active"`
	Author      string    `json:"author"`
	CompanyId   int64     `json:"company_id"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdateDate  time.Time `gorm:"autoUpdateTime"`
}

type ProgramacListResponse struct {
	Items      []ProgramacResp `json:"items"`
	TotalCount int64           `json:"totalCount"`
}

func (ProgramacResp) TableName() string {
	return "programac" // Nombre de la tabla en la base de datos
}

type CreateProgramacReq struct {
	ID          string     `json:"id" gorm:"primaryKey;autoIncrement"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Reserva     float32    `json:"reserva"`
	Active      int        `json:"active"`
	Author      string     `json:"author"`
	CompanyId   int64      `json:"company_id"`
	CreatedDate *time.Time `gorm:"autoCreateTime"`
	UpdateDate  *time.Time `gorm:"autoUpdateTime"`
}

func (CreateProgramacReq) TableName() string {
	return "programac" // Nombre de la tabla en la base de datos
}

type UpdateProgramacReq struct {
	ID          string     `json:"-"`
	Code        *string    `json:"code"`
	Name        *string    `json:"name"`
	Reserva     *float32   `json:"reserva"`
	Active      *int       `json:"active"`
	Author      *string    `json:"author"`
	CompanyId   *int64     `json:"company_id"`
	CreatedDate *time.Time `gorm:"autoCreateTime"`
	UpdateDate  *time.Time `gorm:"autoUpdateTime"`
}

func (UpdateProgramacReq) TableName() string {
	return "programac" // Nombre de la tabla en la base de datos
}

type ProgramacReport struct {
	ID        int64  `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Active    int    `json:"active"`
	Reserva   int    `json:"reserva"`
	CompanyId int64  `json:"company_id"`
}

func (ProgramacReport) TableName() string {
	return "programac" // Nombre de la tabla en la base de datos
}
