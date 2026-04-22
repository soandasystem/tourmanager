package models

import "time"

type Pagos struct {
	ID            string    `json:"_id,omitempty"`
	Tipocom       string    `json:"tipocom"`
	IngresoId     int64     `json:"ingreso_id"`
	Identificador string    `json:"identificador"`
	Fecha         time.Time `json:"fecha"`
	SaleId        int64     `json:"sale_id"`
	Rutalumn      string    `json:"rutalumn"`
	Transaccion   string    `json:"transaccion"`
	Tipo          string    `json:"tipo"`
	Monto         float64   `json:"monto"`
	Nrotarjeta    string    `json:"nrotarjeta"`
	Codigoauto    string    `json:"codigoauto"`
	Fechaauto     time.Time `json:"fechaauto"`
	Tipopago      string    `json:"tipopago"`
	Nrocuota      int       `json:"nrocuota"`
	Fechatransac  time.Time `json:"fechatransac"`
	Author        string    `json:"author"`
	Activo        int       `json:"activo"`
	CompanyId     int64     `json:"company_id"`
	Cuotapagada   int       `json:"cuotapagada"`
	Cuotafecha    string    `json:"cuotafecha"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

// Resp  response struct
type PagosResp struct {
	ID            string    `json:"id"`
	Tipocom       string    `json:"tipocom"`
	IngresoId     int64     `json:"ingreso_id"`
	Identificador string    `json:"identificador"`
	Fecha         time.Time `json:"fecha"`
	SaleId        int64     `json:"sale_id"`
	Rutalumn      string    `json:"rutalumn"`
	Transaccion   string    `json:"transaccion"`
	Tipo          string    `json:"tipo"`
	Monto         float64   `json:"monto"`
	Nrotarjeta    string    `json:"nrotarjeta"`
	Codigoauto    string    `json:"codigoauto"`
	Fechaauto     time.Time `json:"fechaauto"`
	Tipopago      string    `json:"tipopago"`
	Nrocuota      int       `json:"nrocuota"`
	Fechatransac  time.Time `json:"fechatransac"`
	Author        string    `json:"author"`
	Activo        int       `json:"activo"`
	CompanyId     int64     `json:"company_id"`
	Cuotapagada   int       `json:"cuotapagada"`
	Cuotafecha    string    `json:"cuotafecha"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

func (PagosResp) TableName() string {
	return "pagos" // Nombre de la tabla en la base de datos
}

type PagosListResponse struct {
	Items      []PagosResp `json:"items"`
	TotalCount int64       `json:"totalCount"`
}

// Create---Req  request struct
type CreatePagosReq struct {
	ID            string    `gorm:"primaryKey;autoIncrement"`
	Tipocom       string    `json:"tipocom"`
	IngresoId     int64     `json:"ingreso_id"`
	Identificador string    `json:"identificador"`
	Fecha         time.Time `json:"fecha"`
	SaleId        int64     `json:"sale_id"`
	Rutalumn      string    `json:"rutalumn"`
	Transaccion   string    `json:"transaccion"`
	Tipo          string    `json:"tipo"`
	Monto         float64   `json:"monto"`
	Nrotarjeta    string    `json:"nrotarjeta"`
	Codigoauto    string    `json:"codigoauto"`
	Fechaauto     time.Time `json:"fechaauto"`
	Tipopago      string    `json:"tipopago"`
	Nrocuota      int       `json:"nrocuota"`
	Fechatransac  time.Time `json:"fechatransac"`
	Author        string    `json:"author"`
	Activo        int       `json:"activo"`
	CompanyId     int64     `json:"company_id"`
	Cuotapagada   int       `json:"cuotapagada"`
	Cuotafecha    string    `json:"cuotafecha"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

func (CreatePagosReq) TableName() string {
	return "pagos" // Nombre de la tabla en la base de datos
}

type UpdatePagosReq struct {
	ID            string     `json:"-"`
	Tipocom       *string    `json:"tipocom"`
	IngresoId     *int64     `json:"ingreso_id"`
	Identificador *string    `json:"identificador"`
	Fecha         *time.Time `json:"fecha"`
	SaleId        *int64     `json:"sale_id"`
	Rutalumn      *string    `json:"rutalumn"`
	Transaccion   *string    `json:"transaccion"`
	Tipo          *string    `json:"tipo"`
	Monto         *float64   `json:"monto"`
	Nrotarjeta    *string    `json:"nrotarjeta"`
	Codigoauto    *string    `json:"codigoauto"`
	Fechaauto     *time.Time `json:"fechaauto"`
	Tipopago      *string    `json:"tipopago"`
	Nrocuota      *int       `json:"nrocuota"`
	Fechatransac  *time.Time `json:"fechatransac"`
	Author        *string    `json:"author"`
	Activo        *int       `json:"activo"`
	CompanyId     *int64     `json:"company_id"`
	Cuotapagada   *int       `json:"cuotapagada"`
	Cuotafecha    *string    `json:"cuotafecha"`
	CreatedDate   *time.Time `gorm:"autoCreateTime"`
	UpdatedDate   *time.Time `gorm:"autoUpdateTime"`
}

func (UpdatePagosReq) TableName() string {
	return "pagos" // Nombre de la tabla en la base de datos
}

type PagosReport struct {
	ID            string    `json:"-"`
	Tipocom       string    `json:"tipocom"`
	IngresoId     int64     `json:"ingreso_id"`
	Identificador string    `json:"identificador"`
	Fecha         time.Time `json:"fecha"`
	SaleId        int64     `json:"sale_id"`
	Rutalumn      string    `json:"rutalumn"`
	Transaccion   string    `json:"transaccion"`
	Tipo          string    `json:"tipo"`
	Monto         float64   `json:"monto"`
	Nrotarjeta    string    `json:"nrotarjeta"`
	Codigoauto    string    `json:"codigoauto"`
	Fechaauto     time.Time `json:"fechaauto"`
	Tipopago      string    `json:"tipopago"`
	Nrocuota      int       `json:"nrocuota"`
	Fechatransac  time.Time `json:"fechatransac"`
	Author        string    `json:"author"`
	Activo        int       `json:"activo"`
	CompanyId     int64     `json:"company_id"`
	Cuotapagada   int       `json:"cuotapagada"`
	Cuotafecha    string    `json:"cuotafecha"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

func (PagosReport) TableName() string {
	return "pagos" // Nombre de la tabla en la base de datos
}

type PagosInf struct {
	ID            int64     `json:"_id,omitempty"`
	Tipocom       string    `json:"tipocom"`
	IngresoId     int64     `json:"ingreso_id"`
	Identificador string    `json:"identificador"`
	Fecha         time.Time `json:"fecha"`
	SaleId        int64     `json:"sale_id"`
	Rutalumn      string    `json:"rutalumn"`
	Transaccion   string    `json:"transaccion"`
	Tipo          string    `json:"tipo"`
	Monto         float64   `json:"monto"`
	Nrotarjeta    string    `json:"nrotarjeta"`
	Codigoauto    string    `json:"codigoauto"`
	Fechaauto     time.Time `json:"fechaauto"`
	Tipopago      string    `json:"tipopago"`
	Nrocuota      int       `json:"nrocuota"`
	Fechatransac  time.Time `json:"fechatransac"`
	Author        string    `json:"author"`
	Activo        int       `json:"activo"`
	CompanyId     int64     `json:"company_id"`
	Cuotapagada   int       `json:"cuotapagada"`
	Cuotafecha    string    `json:"cuotafecha"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

func (PagosInf) TableName() string {
	return "pagos" // Nombre de la tabla en la base de datos
}

type PagosInfListResponse struct {
	Items      []PagosInf `json:"items"`
	TotalCount int64      `json:"totalCount"`
}
