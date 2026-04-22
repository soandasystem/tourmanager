package models

import "time"

type Ingreso struct {
	ID            string    `json:"_id,omitempty"`
	Tipocomp      string    `json:"tipocomp"`
	Fecha         time.Time `json:"fecha"`
	Identificador string    `json:"identificador"`
	SaleId        int64     `json:"sale_id"`
	CursoId       int64     `json:"curso_id"`
	Rutapo        string    `json:"rutapo"`
	Rutalum       string    `json:"rutalum"`
	Fpago         string    `json:"fpago"`
	Monto         float32   `json:"monto"`
	Activo        int       `json:"activo"`
	StatusPago    string    `json:"status_pago"`
	Author        string    `json:"author"`
	CompanyId     int64     `json:"company_id"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
	TokenFlow     string    `json:"token_flow"`
	Nrocuotas     int       `json:"nrocuotas"`
	Valorcuota    float32   `json:"valorcuota"`
	Fechainicial  time.Time `json:"fechainicial"`
	Reserva       int       `json:"reserva"`
}

// Resp  response struct
type IngresoResp struct {
	ID            string    `json:"id"`
	Tipocomp      string    `json:"tipocomp"`
	Fecha         time.Time `json:"fecha"`
	Identificador string    `json:"identificador"`
	SaleId        int64     `json:"sale_id"`
	CursoId       int64     `json:"curso_id"`
	Rutapo        string    `json:"rutapo"`
	Rutalum       string    `json:"rutalum"`
	Fpago         string    `json:"fpago"`
	Monto         float32   `json:"monto"`
	Activo        int       `json:"activo"`
	StatusPago    string    `json:"status_pago"`
	Author        string    `json:"author"`
	CompanyId     int64     `json:"company_id"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
	TokenFlow     string    `json:"token_flow"`
	Nrocuotas     int       `json:"nrocuotas"`
	Valorcuota    float32   `json:"valorcuota"`
	Fechainicial  time.Time `json:"fechainicial"`
	Reserva       int       `json:"reserva"`
}

func (IngresoResp) TableName() string {
	return "ingresos" // Nombre de la tabla en la base de datos
}

type IngresoListResponse struct {
	Items      []IngresoResp `json:"items"`
	TotalCount int64         `json:"totalCount"`
}

// Create---Req  request struct
type CreateIngresoReq struct {
	ID            string    `gorm:"primaryKey;autoIncrement"`
	Tipocomp      string    `json:"tipocomp"`
	Fecha         time.Time `json:"fecha"`
	Identificador string    `json:"identificador"`
	SaleId        int64     `json:"sale_id"`
	CursoId       int64     `json:"curso_id"`
	Rutapo        string    `json:"rutapo"`
	Rutalum       string    `json:"rutalum"`
	Fpago         string    `json:"fpago"`
	Monto         float32   `json:"monto"`
	Activo        int       `json:"activo"`
	StatusPago    string    `json:"status_pago"`
	Author        string    `json:"author"`
	CompanyId     int64     `json:"company_id"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
	TokenFlow     string    `json:"token_flow"`
	Nrocuotas     int       `json:"nrocuotas"`
	Valorcuota    float32   `json:"valorcuota"`
	Fechainicial  time.Time `json:"fechainicial"`
	Reserva       int       `json:"reserva"`
}

func (CreateIngresoReq) TableName() string {
	return "ingresos" // Nombre de la tabla en la base de datos
}

type UpdateIngresoReq struct {
	ID            string     `json:"-"`
	Tipocomp      *string    `json:"tipocomp"`
	Fecha         *time.Time `json:"fecha"`
	Identificador *string    `json:"identificador"`
	SaleId        *int64     `json:"sale_id"`
	CursoId       *int64     `json:"curso_id"`
	Rutapo        *string    `json:"rutapo"`
	Rutalum       *string    `json:"rutalum"`
	Fpago         *string    `json:"fpago"`
	Monto         *float32   `json:"monto"`
	Activo        *int       `json:"activo"`
	StatusPago    *string    `json:"status_pago"`
	Author        *string    `json:"author"`
	CompanyId     *int64     `json:"company_id"`
	CreatedDate   *time.Time `gorm:"autoCreateTime"`
	UpdatedDate   *time.Time `gorm:"autoUpdateTime"`
	TokenFlow     string     `json:"token_flow"`
	Nrocuotas     int        `json:"nrocuotas"`
	Valorcuota    float32    `json:"valorcuota"`
	Fechainicial  time.Time  `json:"fechainicial"`
	Reserva       int        `json:"reserva"`
}

func (UpdateIngresoReq) TableName() string {
	return "ingresos" // Nombre de la tabla en la base de datos
}

type IngresoInf struct {
	ID            int64         `json:"id"`
	Tipocomp      string        `json:"tipocomp"`
	Fecha         time.Time     `json:"fecha"`
	Identificador string        `json:"identificador"`
	SaleId        int64         `json:"sale_id"`
	Sale          SaleReport    `json:"sale" gorm:"foreignKey:SaleId;references:ID"`
	CursoId       int64         `json:"curso_id"`
	Curso         CursoReport   `json:"curso" gorm:"foreignKey:CursoId;references:ID"`
	Rutapo        string        `json:"rutapo"`
	Rutalum       string        `json:"rutalum"`
	Fpago         string        `json:"fpago"`
	Monto         float32       `json:"monto"`
	Activo        int           `json:"activo"`
	StatusPago    string        `json:"status_pago"`
	Author        string        `json:"author"`
	CompanyId     int64         `json:"company_id"`
	Pago          []PagosReport `json:"pago" gorm:"foreignKey:IngresoId;references:ID"`
	CreatedDate   time.Time     `gorm:"autoCreateTime"`
	TokenFlow     string        `json:"token_flow"`
	Nrocuotas     int           `json:"nrocuotas"`
	Valorcuota    float32       `json:"valorcuota"`
	Fechainicial  time.Time     `json:"fechainicial"`
	UpdatedDate   time.Time     `gorm:"autoUpdateTime"`
	Reserva       int           `json:"reserva"`
}

func (IngresoInf) TableName() string {
	return "ingresos" // Nombre de la tabla en la base de datos
}

type IngresoInfListResponse struct {
	Items      []IngresoInf `json:"items"`
	TotalCount int64        `json:"totalCount"`
}

type IngresoReport struct {
	ID            int64     `json:"id"`
	Tipocomp      string    `json:"tipocomp"`
	Fecha         time.Time `json:"fecha"`
	Identificador string    `json:"identificador"`
	SaleId        int64     `gorm:"column:sale_id"`
	CursoId       int64     `json:"curso_id"`
	Rutapo        string    `json:"rutapo"`
	Rutalum       string    `json:"rutalum"`
	Fpago         string    `json:"fpago"`
	Monto         float32   `json:"monto"`
	Activo        int       `json:"activo"`
	StatusPago    string    `json:"status_pago"`
	Author        string    `json:"author"`
	CompanyId     int64     `json:"company_id"`
	TokenFlow     string    `json:"token_flow"`
	Nrocuotas     int       `json:"nrocuotas"`
	Valorcuota    float32   `json:"valorcuota"`
	Fechainicial  time.Time `json:"fechainicial"`
	Reserva       int       `json:"reserva"`
}

func (IngresoReport) TableName() string {
	return "ingresos" // Nombre de la tabla en la base de datos
}
