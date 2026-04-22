package models

import (
	"time"
)

type Sale struct {
	ID                string    `json:"_id,omitempty"`
	Fecha             time.Time `json:"fecha"`
	SellerId          int64     `json:"seller_id"`
	Identificador     string    `json:"identificador"`
	EstablecimientoId int64     `json:"establecimiento_id"`
	ProgramId         int64     `json:"program_id"`
	Curso             int       `json:"curso"`
	Idcurso           string    `json:"idcurso"`
	Nroalumno         int       `json:"nroalumno"`
	Liberados         int       `json:"liberados"`
	Subtotal          int       `json:"subtotal"`
	Descm             int       `json:"descm"`
	Vprograma         int       `json:"vprograma"`
	Description       string    `json:"descrition"`
	Obs               string    `json:"obs"`
	Fechasalida       time.Time `json:"fechasalida"`
	Activo            int       `json:"activo"`
	State             string    `json:"state"`
	CorreoEncargado   string    `json:"correo_encargado"`
	Password          string    `json:"password"`
	FechaUltpag       time.Time `json:"fecha_ultpag"`
	FechaCierre       time.Time `json:"fecha_cierre"`
	Sendemail         int       `json:"sendemail"`
	Author            string    `json:"author"`
	Encargado         string    `json:"encargado"`
	Comision          float32   `json:"comision"`
	Tipocambio        float32   `json:"tipocambio"`
	ComisionPagada    int       `json:"comision_pagada"`
	CompanyId         int64     `json:"company_id"`
	Cuotas            int       `json:"cuotas"`
	Fechacuota        time.Time `json:"fechacuota"`
	Accesscode        string    `json:"accesscode"`
	TypeSale          string    `json:"type_sale"`
	CreatedDate       time.Time `gorm:"autoCreateTime"`
	UpdatedDate       time.Time `gorm:"autoUpdateTime"`
}

// Resp  response struct
type SaleResp struct {
	ID                string    `json:"id"`
	Fecha             time.Time `json:"fecha"`
	SellerId          int64     `json:"seller_id"`
	Identificador     string    `json:"identificador"`
	EstablecimientoId int64     `json:"establecimiento_id"`
	ProgramId         int64     `json:"program_id"`
	Curso             int       `json:"curso"`
	Idcurso           string    `json:"idcurso"`
	Nroalumno         int       `json:"nroalumno"`
	Liberados         int       `json:"liberados"`
	Subtotal          int       `json:"subtotal"`
	Descm             int       `json:"descm"`
	Vprograma         int       `json:"vprograma"`
	Description       string    `json:"descrition"`
	Obs               string    `json:"obs"`
	Fechasalida       time.Time `json:"fechasalida"`
	Activo            int       `json:"activo"`
	State             string    `json:"state"`
	CorreoEncargado   string    `json:"correo_encargado"`
	Password          string    `json:"password"`
	FechaUltpag       time.Time `json:"fecha_ultpag"`
	FechaCierre       time.Time `json:"fecha_cierre"`
	Sendemail         int       `json:"sendemail"`
	Author            string    `json:"author"`
	Encargado         string    `json:"encargado"`
	Comision          float32   `json:"comision"`
	Tipocambio        float32   `json:"tipocambio"`
	ComisionPagada    int       `json:"comision_pagada"`
	CompanyId         int64     `json:"company_id"`
	Cuotas            int       `json:"cuotas"`
	Fechacuota        time.Time `json:"fechacuota"`
	Accesscode        string    `json:"accesscode"`
	TypeSale          string    `json:"type_sale"`
	CreatedDate       time.Time `gorm:"autoCreateTime"`
	UpdatedDate       time.Time `gorm:"autoUpdateTime"`
}

func (SaleResp) TableName() string {
	return "sales" // Nombre de la tabla en la base de datos
}

type SaleListResponse struct {
	Items      []SaleResp `json:"items"`
	TotalCount int64      `json:"totalCount"`
}

// Create---Req  request struct
type CreateSaleReq struct {
	ID                string    `gorm:"primaryKey;autoIncrement"`
	Fecha             time.Time `json:"fecha"`
	SellerId          int64     `json:"seller_id"`
	Identificador     string    `json:"identificador"`
	EstablecimientoId int64     `json:"establecimiento_id"`
	ProgramId         int64     `json:"program_id"`
	Curso             int       `json:"curso"`
	Idcurso           string    `json:"idcurso"`
	Nroalumno         int       `json:"nroalumno"`
	Liberados         int       `json:"liberados"`
	Subtotal          int       `json:"subtotal"`
	Descm             int       `json:"descm"`
	Vprograma         int       `json:"vprograma"`
	Description       string    `json:"descrition"`
	Obs               string    `json:"obs"`
	Fechasalida       time.Time `json:"fechasalida"`
	Activo            int       `json:"activo"`
	State             string    `json:"state"`
	CorreoEncargado   string    `json:"correo_encargado"`
	Password          string    `json:"password"`
	FechaUltpag       time.Time `json:"fecha_ultpag"`
	FechaCierre       time.Time `json:"fecha_cierre"`
	Sendemail         int       `json:"sendemail"`
	Author            string    `json:"author"`
	Encargado         string    `json:"encargado"`
	Comision          float32   `json:"comision"`
	Tipocambio        float32   `json:"tipocambio"`
	ComisionPagada    int       `json:"comision_pagada"`
	CompanyId         int64     `json:"company_id"`
	Cuotas            int       `json:"cuotas"`
	Fechacuota        time.Time `json:"fechacuota"`
	Accesscode        string    `json:"accesscode"`
	TypeSale          string    `json:"type_sale"`
	CreatedDate       time.Time `gorm:"autoCreateTime"`
	UpdatedDate       time.Time `gorm:"autoUpdateTime"`
}

func (CreateSaleReq) TableName() string {
	return "sales" // Nombre de la tabla en la base de datos
}

type UpdateSaleReq struct {
	ID                string     `json:"-"`
	Fecha             *time.Time `json:"fecha"`
	SellerId          *int64     `json:"seller_id"`
	Identificador     *string    `json:"identificador"`
	EstablecimientoId *int64     `json:"establecimiento_id"`
	ProgramId         *int64     `json:"program_id"`
	Curso             *int       `json:"curso"`
	Idcurso           *string    `json:"idcurso"`
	Nroalumno         *int       `json:"nroalumno"`
	Liberados         *int       `json:"liberados"`
	Subtotal          *int       `json:"subtotal"`
	Descm             *int       `json:"descm"`
	Vprograma         *int       `json:"vprograma"`
	Description       *string    `json:"descrition"`
	Obs               *string    `json:"obs"`
	Fechasalida       *time.Time `json:"fechasalida"`
	Activo            *int       `json:"activo"`
	State             *string    `json:"state"`
	CorreoEncargado   *string    `json:"correo_encargado"`
	Password          *string    `json:"password"`
	FechaUltpag       *time.Time `json:"fecha_ultpag"`
	FechaCierre       *time.Time `json:"fecha_cierre"`
	Sendemail         *int       `json:"sendemail"`
	Author            *string    `json:"author"`
	Encargado         *string    `json:"encargado"`
	Comision          *float32   `json:"comision"`
	Tipocambio        *float32   `json:"tipocambio"`
	ComisionPagada    *int       `json:"comision_pagada"`
	CompanyId         *int64     `json:"company_id"`
	Cuotas            *int       `json:"cuotas"`
	Fechacuota        *time.Time `json:"fechacuota"`
	Accesscode        *string    `json:"accesscode"`
	TypeSale          *string    `json:"type_sale"`
	CreatedDate       *time.Time `gorm:"autoCreateTime"`
	UpdatedDate       *time.Time `gorm:"autoUpdateTime"`
}

func (UpdateSaleReq) TableName() string {
	return "sales" // Nombre de la tabla en la base de datos
}

type SaleReport struct {
	ID                string    `json:"id"`
	Fecha             time.Time `json:"fecha"`
	SellerId          string    `json:"seller_id"`
	Identificador     string    `json:"identificador"`
	EstablecimientoId string    `json:"establecimiento_id"`
	ProgramId         string    `json:"program_id"`
	Curso             int       `json:"curso"`
	Idcurso           string    `json:"idcurso"`
	Nroalumno         int       `json:"nroalumno"`
	Liberados         int       `json:"liberados"`
	Subtotal          int       `json:"subtotal"`
	Descm             int       `json:"descm"`
	Vprograma         int       `json:"vprograma"`
	Description       string    `json:"descrition"`
	Obs               string    `json:"obs"`
	Fechasalida       time.Time `json:"fechasalida"`
	Activo            int       `json:"activo"`
	State             string    `json:"state"`
	CorreoEncargado   string    `json:"correo_encargado"`
	Password          string    `json:"password"`
	FechaUltpag       time.Time `json:"fecha_ultpag"`
	FechaCierre       time.Time `json:"fecha_cierre"`
	Sendemail         int       `json:"sendemail"`
	Author            string    `json:"author"`
	Encargado         string    `json:"encargado"`
	Comision          float32   `json:"comision"`
	Tipocambio        float32   `json:"tipocambio"`
	ComisionPagada    int       `json:"comision_pagada"`
	CompanyId         int64     `json:"company_id"`
	Cuotas            int       `json:"cuotas"`
	Fechacuota        time.Time `json:"fechacuota"`
	Accesscode        string    `json:"accesscode"`
	TypeSale          string    `json:"type_sale"`
	CreatedDate       time.Time `gorm:"autoCreateTime"`
	UpdatedDate       time.Time `gorm:"autoUpdateTime"`
}

func (SaleReport) TableName() string {
	return "sales" // Nombre de la tabla en la base de datos
}

type SaleCursoReport struct {
	ID                string `json:"id"`
	SellerId          string `json:"seller_id"`
	Identificador     string `json:"identificador"`
	EstablecimientoId string `json:"establecimiento_id"`
	ProgramId         string `json:"program_id"`
	Curso             int    `json:"curso"`
	Idcurso           string `json:"idcurso"`
	TypeSale          string `json:"type_sale"`
}

func (SaleCursoReport) TableName() string {
	return "sales" // Nombre de la tabla en la base de datos
}

type SaleInforme struct {
	ID                    string    `json:"id"`
	Fecha                 time.Time `json:"fecha"`
	SellerId              int64     `json:"seller_id"`
	Identificador         string    `json:"identificador"`
	EstablecimientoId     int64     `json:"establecimiento_id"`
	ProgramId             int64     `json:"program_id"`
	Curso                 int       `json:"curso"`
	Idcurso               string    `json:"idcurso"`
	Nroalumno             int       `json:"nroalumno"`
	Liberados             int       `json:"liberados"`
	Subtotal              int       `json:"subtotal"`
	Descm                 int       `json:"descm"`
	Vprograma             int       `json:"vprograma"`
	Description           string    `json:"descrition"`
	Obs                   string    `json:"obs"`
	Fechasalida           time.Time `json:"fechasalida"`
	Activo                int       `json:"activo"`
	State                 string    `json:"state"`
	CorreoEncargado       string    `json:"correo_encargado"`
	Password              string    `json:"password"`
	FechaUltpag           time.Time `json:"fecha_ultpag"`
	FechaCierre           time.Time `json:"fecha_cierre"`
	Sendemail             int       `json:"sendemail"`
	Author                string    `json:"author"`
	Encargado             string    `json:"encargado"`
	Comision              float32   `json:"comision"`
	Tipocambio            float32   `json:"tipocambio"`
	ComisionPagada        int       `json:"comision_pagada"`
	CompanyId             int64     `json:"company_id"`
	Cuotas                int       `json:"cuotas"`
	TypeSale              string    `json:"type_sale"`
	Fechacuota            time.Time `json:"fechacuota"`
	CreatedDate           time.Time `gorm:"autoCreateTime"`
	UpdatedDate           time.Time `gorm:"autoUpdateTime"`
	TotalCurso            int       `json:"total_curso" gorm:"total_curso"`
	EstablecimientoNombre string    `json:"establecimiento_nombre" gorm:"establecimiento_nombre"`
	SellerName            string    `json:"seller_name" gorm:"seller_name"`
	ProgramName           string    `json:"program_name" gorm:"program_name"`
}

func (SaleInforme) TableName() string {
	return "sales" // Nombre de la tabla en la base de datos
}

type SaleInfListResponse struct {
	Items      []SaleInforme `json:"items"`
	TotalCount int64         `json:"totalCount"`
}
