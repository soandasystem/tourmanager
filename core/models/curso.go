package models

import (
	"encoding/json"
	"time"
)

type Curso struct {
	ID             string    `json:"_id,omitempty"`
	SaleId         int8      `json:"sale_id"`
	Rutalumno      string    `json:"rutalumno"`
	Nombrealumno   string    `json:"nombrealumno"`
	Fechanac       time.Time `json:"fechanac"`
	Rutapod        string    `json:"rutapod"`
	Nombreapod     string    `json:"nombreapod"`
	Dircalle       string    `json:"dircalle"`
	Dirnumero      string    `json:"dirnumero"`
	Nrodepto       string    `json:"nrodepto"`
	RegionId       int64     `json:"region_id"`
	ComunaId       int64     `json:"comuna_id"`
	Fono           string    `json:"fono"`
	Celular        string    `json:"celular"`
	Correo         string    `json:"correo"`
	Vpagar         float64   `json:"vpagar"`
	Descto         float64   `json:"descto"`
	Apagar         float64   `json:"apagar"`
	Liberado       int       `json:"liberado"`
	Enviado        string    `json:"enviado"`
	Estado         string    `json:"estado"`
	Password       string    `json:"password"`
	AceptaContrato int       `json:"acepta_contrato"`
	Signature      string    `json:"signature"`
	Author         string    `json:"author"`
	Pasaporte      string    `json:"pasaporte"`
	CompanyId      int64     `json:"company_id"`
	CreatedDate    time.Time `json:"created_date"`
	UpdatedDate    time.Time `json:"updated_date"`
}

// Resp  response struct
type CursoResp struct {
	ID             string    `json:"id"`
	SaleId         int8      `json:"sale_id"`
	Rutalumno      string    `json:"rutalumno"`
	Nombrealumno   string    `json:"nombrealumno"`
	Fechanac       time.Time `json:"fechanac"`
	Rutapod        string    `json:"rutapod"`
	Nombreapod     string    `json:"nombreapod"`
	Dircalle       string    `json:"dircalle"`
	Dirnumero      string    `json:"dirnumero"`
	Nrodepto       string    `json:"nrodepto"`
	RegionId       int64     `json:"region_id"`
	ComunaId       int64     `json:"comuna_id"`
	Fono           string    `json:"fono"`
	Celular        string    `json:"celular"`
	Correo         string    `json:"correo"`
	Vpagar         float64   `json:"vpagar"`
	Descto         float64   `json:"descto"`
	Apagar         float64   `json:"apagar"`
	Liberado       int       `json:"liberado"`
	Enviado        string    `json:"enviado"`
	Estado         string    `json:"estado"`
	Password       string    `json:"password"`
	AceptaContrato int       `json:"acepta_contrato"`
	Signature      string    `json:"signature"`
	Author         string    `json:"author"`
	Pasaporte      string    `json:"pasaporte"`
	CompanyId      int64     `json:"company_id"`
	CreatedDate    time.Time `json:"created_date"`
	UpdatedDate    time.Time `json:"updated_date"`
}

func (CursoResp) TableName() string {
	return "cursos" // Nombre de la tabla en la base de datos
}

type CursoListResponse struct {
	Items      []CursoResp `json:"items"`
	TotalCount int64       `json:"totalCount"`
}

// Create---Req  request struct
type CreateCursoReq struct {
	ID             string    `gorm:"primaryKey;autoIncrement"`
	SaleId         int8      `json:"sale_id"`
	Rutalumno      string    `json:"rutalumno"`
	Nombrealumno   string    `json:"nombrealumno"`
	Fechanac       time.Time `json:"fechanac"`
	Rutapod        string    `json:"rutapod"`
	Nombreapod     string    `json:"nombreapod"`
	Dircalle       string    `json:"dircalle"`
	Dirnumero      string    `json:"dirnumero"`
	Nrodepto       string    `json:"nrodepto"`
	RegionId       int64     `json:"region_id"`
	ComunaId       int64     `json:"comuna_id"`
	Fono           string    `json:"fono"`
	Celular        string    `json:"celular"`
	Correo         string    `json:"correo"`
	Vpagar         float64   `json:"vpagar"`
	Descto         float64   `json:"descto"`
	Apagar         float64   `json:"apagar"`
	Liberado       int       `json:"liberado"`
	Enviado        string    `json:"enviado"`
	Estado         string    `json:"estado"`
	Password       string    `json:"password"`
	AceptaContrato int       `json:"acepta_contrato"`
	Signature      string    `json:"signature"`
	Author         string    `json:"author"`
	Pasaporte      string    `json:"pasaporte"`
	CompanyId      int64     `json:"company_id"`
	CreatedDate    time.Time `gorm:"autoCreateTime"`
	UpdatedDate    time.Time `gorm:"autoUpdateTime"`
}

func (CreateCursoReq) TableName() string {
	return "cursos" // Nombre de la tabla en la base de datos
}

type UpdateCursoReq struct {
	ID             string     `json:"-"`
	SaleId         *int8      `json:"sale_id"`
	Rutalumno      *string    `json:"rutalumno"`
	Nombrealumno   *string    `json:"nombrealumno"`
	Fechanac       *time.Time `json:"fechanac"`
	Rutapod        *string    `json:"rutapod"`
	Nombreapod     *string    `json:"nombreapod"`
	Dircalle       *string    `json:"dircalle"`
	Dirnumero      *string    `json:"dirnumero"`
	Nrodepto       *string    `json:"nrodepto"`
	RegionId       *int64     `json:"region_id"`
	ComunaId       *int64     `json:"comuna_id"`
	Fono           *string    `json:"fono"`
	Celular        *string    `json:"celular"`
	Correo         *string    `json:"correo"`
	Vpagar         *float64   `json:"vpagar"`
	Descto         *float64   `json:"descto"`
	Apagar         *float64   `json:"apagar"`
	Liberado       *int       `json:"liberado"`
	Enviado        *string    `json:"enviado"`
	Estado         *string    `json:"estado"`
	Password       *string    `json:"password"`
	AceptaContrato *int       `json:"acepta_contrato"`
	Signature      *string    `json:"signature"`
	Author         *string    `json:"author"`
	Pasaporte      *string    `json:"pasaporte"`
	CompanyId      *int64     `json:"company_id"`
	CreatedDate    *time.Time `gorm:"autoCreateTime"`
	UpdatedDate    *time.Time `gorm:"autoUpdateTime"`
}

func (UpdateCursoReq) TableName() string {
	return "cursos" // Nombre de la tabla en la base de datos
}

// Método MarshalJSON para personalizar el formato JSON de las fechas
func (c CursoResp) MarshalJSON() ([]byte, error) {
	type Alias CursoResp
	return json.Marshal(&struct {
		Alias
		Fechanac    string `json:"fechanac"`
		CreatedDate string `json:"created_date"`
		UpdatedDate string `json:"updated_date"`
	}{
		Alias:       (Alias)(c),
		Fechanac:    c.Fechanac.Format("2006-01-02"),
		CreatedDate: c.CreatedDate.Format("2006-01-02"),
		UpdatedDate: c.UpdatedDate.Format("2006-01-02"),
	})
}

type CursoReport struct {
	ID           string    `json:"-"`
	Rutalumno    string    `json:"rutalumno"`
	Nombrealumno string    `json:"nombrealumno"`
	Fechanac     time.Time `json:"fechanac"`
	Rutapod      string    `json:"rutapod"`
	Nombreapod   string    `json:"nombreapod"`
	Pasaporte    string    `json:"pasaporte"`
}

func (CursoReport) TableName() string {
	return "cursos" // Nombre de la tabla en la base de datos
}

type CursoInf struct {
	ID             int64           `json:"id"`
	SaleId         int8            `json:"sale_id"`
	Sale           SaleCursoReport `json:"sale" gorm:"foreignKey:SaleId;references:ID"`
	Rutalumno      string          `json:"rutalumno"`
	Nombrealumno   string          `json:"nombrealumno"`
	Fechanac       time.Time       `json:"fechanac"`
	Rutapod        string          `json:"rutapod"`
	Nombreapod     string          `json:"nombreapod"`
	Dircalle       string          `json:"dircalle"`
	Dirnumero      string          `json:"dirnumero"`
	Nrodepto       string          `json:"nrodepto"`
	RegionId       int64           `json:"region_id"`
	ComunaId       int64           `json:"comuna_id"`
	Fono           string          `json:"fono"`
	Celular        string          `json:"celular"`
	Correo         string          `json:"correo"`
	Vpagar         float64         `json:"vpagar"`
	Descto         float64         `json:"descto"`
	Apagar         float64         `json:"apagar"`
	Liberado       int             `json:"liberado"`
	Enviado        string          `json:"enviado"`
	Estado         string          `json:"estado"`
	Password       string          `json:"password"`
	AceptaContrato int             `json:"acepta_contrato"`
	Signature      string          `json:"signature"`
	Author         string          `json:"author"`
	Pasaporte      string          `json:"pasaporte"`
	CompanyId      int64           `json:"company_id"`
	Ingreso        []IngresoReport `json:"pago" gorm:"foreignKey:CursoId;references:ID"`
	CreatedDate    time.Time       `json:"created_date"`
	UpdatedDate    time.Time       `json:"updated_date"`
}

func (CursoInf) TableName() string {
	return "cursos" // Nombre de la tabla en la base de datos
}

type CursoInfListResponse struct {
	Items      []CursoInf `json:"items"`
	TotalCount int64      `json:"totalCount"`
}
