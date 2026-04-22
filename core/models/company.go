package models

import (
	"time"
)

type Company struct {
	ID                  string    `json:"_id,omitempty"`
	Rut                 string    `json:"rut"`
	Razonsocial         string    `json:"razonsocial"`
	Nomfantasia         string    `json:"nomfantasia"`
	Direccion           string    `json:"direccion"`
	ComunaId            int64     `json:"comuna_id"`
	RegionId            int64     `json:"region_id"`
	Rutreplegal         string    `json:"rutreplegal"`
	Nomreplegal         string    `json:"nomreplegal"`
	Nombrecontacto1     string    `json:"nombrecontacto1"`
	Fonocontacto1       string    `json:"fonocontacto1"`
	Emailcontacto1      string    `json:"emailcontacto1"`
	Nombrecontacto2     string    `json:"nombrecontacto2"`
	Fonocontacto2       string    `json:"fonocontacto2"`
	Emailcontacto2      string    `json:"emailcontacto2"`
	Iniciooperacion     time.Time `json:"iniciooperacion"`
	Contrato            string    `json:"contrato"`
	ContratoVg          string    `json:"contrato_vg"`
	Website             string    `json:"website"`
	Author              string    `json:"author"`
	Identificador       string    `json:"identificador"`
	Active              int       `json:"active"`
	Email               string    `json:"email"`
	SchemaName          string    `json:"schema_name"`
	PlancodeId          int       `json:"plancode_id"`
	Additionaluser      int       `json:"additionaluser"`
	Maxusers            int       `json:"maxusers"`
	Maxquote            int       `json:"maxquote"`
	Maxsales            int       `json:"maxsales"`
	Terminoscondiciones int       `json:"terminoscondiciones"`
	Politicasdeuso      int       `json:"politicasdeuso"`
	CreatedDate         time.Time `gorm:"autoCreateTime"`
	UpdatedDate         time.Time `gorm:"autoUpdateTime"`
}

// Resp  response struct
type CompanyResp struct {
	ID                  string    `json:"id"`
	Rut                 string    `json:"rut"`
	Razonsocial         string    `json:"razonsocial"`
	Nomfantasia         string    `json:"nomfantasia"`
	Direccion           string    `json:"direccion"`
	ComunaId            int64     `json:"comuna_id"`
	RegionId            int64     `json:"region_id"`
	Rutreplegal         string    `json:"rutreplegal"`
	Nomreplegal         string    `json:"nomreplegal"`
	Nombrecontacto1     string    `json:"nombrecontacto1"`
	Fonocontacto1       string    `json:"fonocontacto1"`
	Emailcontacto1      string    `json:"emailcontacto1"`
	Nombrecontacto2     string    `json:"nombrecontacto2"`
	Fonocontacto2       string    `json:"fonocontacto2"`
	Emailcontacto2      string    `json:"emailcontacto2"`
	Iniciooperacion     time.Time `json:"iniciooperacion"`
	Contrato            string    `json:"contrato"`
	ContratoVg          string    `json:"contrato_vg"`
	Website             string    `json:"website"`
	Author              string    `json:"author"`
	Identificador       string    `json:"identificador"`
	Active              int       `json:"active"`
	Email               string    `json:"email"`
	SchemaName          string    `json:"schema_name"`
	PlancodeId          int       `json:"plancode_id"`
	Additionaluser      int       `json:"additionaluser"`
	Maxusers            int       `json:"maxusers"`
	Maxquote            int       `json:"maxquote"`
	Maxsales            int       `json:"maxsales"`
	Terminoscondiciones int       `json:"terminoscondiciones"`
	Politicasdeuso      int       `json:"politicasdeuso"`
	CreatedDate         time.Time `gorm:"autoCreateTime"`
	UpdatedDate         time.Time `gorm:"autoUpdateTime"`
}

func (CompanyResp) TableName() string {
	return "company" // Nombre de la tabla en la base de datos
}

type CompanyListResponse struct {
	Items      []CompanyResp `json:"items"`
	TotalCount int64         `json:"totalCount"`
}

// Create---Req  request struct
type CreateCompanyReq struct {
	ID                  string    `gorm:"primaryKey;autoIncrement"`
	Rut                 string    `json:"rut"`
	Razonsocial         string    `json:"razonsocial"`
	Nomfantasia         string    `json:"nomfantasia"`
	Direccion           string    `json:"direccion"`
	ComunaId            int64     `json:"comuna_id"`
	RegionId            int64     `json:"region_id"`
	Rutreplegal         string    `json:"rutreplegal"`
	Nomreplegal         string    `json:"nomreplegal"`
	Nombrecontacto1     string    `json:"nombrecontacto1"`
	Fonocontacto1       string    `json:"fonocontacto1"`
	Emailcontacto1      string    `json:"emailcontacto1"`
	Nombrecontacto2     string    `json:"nombrecontacto2"`
	Fonocontacto2       string    `json:"fonocontacto2"`
	Emailcontacto2      string    `json:"emailcontacto2"`
	Iniciooperacion     time.Time `json:"iniciooperacion"`
	Contrato            string    `json:"contrato"`
	ContratoVg          string    `json:"contrato_vg"`
	Website             string    `json:"website"`
	Author              string    `json:"author"`
	Identificador       string    `json:"identificador"`
	Active              int       `json:"active"`
	Email               string    `json:"email"`
	SchemaName          string    `json:"schema_name"`
	PlancodeId          int       `json:"plancode_id"`
	Additionaluser      int       `json:"additionaluser"`
	Maxusers            int       `json:"maxusers"`
	Maxquote            int       `json:"maxquote"`
	Maxsales            int       `json:"maxsales"`
	Terminoscondiciones int       `json:"terminoscondiciones"`
	Politicasdeuso      int       `json:"politicasdeuso"`
	CreatedDate         time.Time `gorm:"autoCreateTime"`
	UpdatedDate         time.Time `gorm:"autoUpdateTime"`
}

func (CreateCompanyReq) TableName() string {
	return "company" // Nombre de la tabla en la base de datos
}

type UpdateCompanyReq struct {
	ID                  string     `json:"-"`
	Rut                 *string    `json:"rut"`
	Razonsocial         *string    `json:"razonsocial"`
	Nomfantasia         *string    `json:"nomfantasia"`
	Direccion           *string    `json:"direccion"`
	ComunaId            *int64     `json:"comuna_id"`
	RegionId            *int64     `json:"region_id"`
	Rutreplegal         *string    `json:"rutreplegal"`
	Nomreplegal         *string    `json:"nomreplegal"`
	Nombrecontacto1     *string    `json:"nombrecontacto1"`
	Fonocontacto1       *string    `json:"fonocontacto1"`
	Emailcontacto1      *string    `json:"emailcontacto1"`
	Nombrecontacto2     *string    `json:"nombrecontacto2"`
	Fonocontacto2       *string    `json:"fonocontacto2"`
	Emailcontacto2      *string    `json:"emailcontacto2"`
	Iniciooperacion     *time.Time `json:"iniciooperacion"`
	Contrato            *string    `json:"contrato"`
	ContratoVg          *string    `json:"contrato_vg"`
	Website             *string    `json:"website"`
	Author              *string    `json:"author"`
	Identificador       *string    `json:"identificador"`
	Active              *int       `json:"active"`
	Email               *string    `json:"email"`
	SchemaName          *string    `json:"schema_name"`
	PlancodeId          *int       `json:"plancode_id"`
	Additionaluser      *int       `json:"additionaluser"`
	Maxusers            *int       `json:"maxusers"`
	Maxquote            *int       `json:"maxquote"`
	Maxsales            *int       `json:"maxsales"`
	Terminoscondiciones *int       `json:"terminoscondiciones"`
	Politicasdeuso      *int       `json:"politicasdeuso"`
	CreatedDate         time.Time  `gorm:"autoCreateTime"`
	UpdatedDate         time.Time  `gorm:"autoUpdateTime"`
}

func (UpdateCompanyReq) TableName() string {
	return "company" // Nombre de la tabla en la base de datos
}
