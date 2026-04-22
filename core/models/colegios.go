package models

import "github.com/shopspring/decimal"

type Establecimiento struct {
	ID        string `json:"_id,omitempty"`
	Codigo    string `json:"codigo"`
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
	Comuna    string `json:"comuna"`
	Latitud   int16  `json:"latitiud"`
	Longitud  int16  `json:"longitud"`
	RegionId  int64  `json:"region_id"`
	ComunaId  int64  `json:"comuna_id"`
	CompanyId int64  `json:"company_id"`
}

type Colegios struct {
	ID        string `json:"_id,omitempty"`
	Codigo    string `json:"codigo"`
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
	Comuna    string `json:"comuna"`
	Latitud   int16  `json:"latitiud"`
	Longitud  int16  `json:"longitud"`
	RegionId  int64  `json:"region_id"`
	ComunaId  int64  `json:"comuna_id"`
	CompanyId int64  `json:"company_id"`
}

func (Colegios) TableName() string {
	return "establecimientos" // Nombre de la tabla en la base de datos
}

// Resp  response struct
type ColegiosResp struct {
	ID        string `json:"id"`
	Codigo    string `json:"codigo"`
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
	Comuna    string `json:"comuna"`
	Latitud   int16  `json:"latitiud"`
	Longitud  int16  `json:"longitud"`
	RegionId  int64  `json:"region_id"`
	ComunaId  int64  `json:"comuna_id"`
	CompanyId int64  `json:"company_id"`
}

func (ColegiosResp) TableName() string {
	return "establecimientos" // Nombre de la tabla en la base de datos
}

type ColegiosListResponse struct {
	Items      []ColegiosResp `json:"items"`
	TotalCount int64          `json:"totalCount"`
}

// Create---Req  request struct
type CreateColegiosReq struct {
	ID        string          `gorm:"primaryKey;autoIncrement"`
	Codigo    string          `json:"codigo"`
	Nombre    string          `json:"nombre"`
	Direccion string          `json:"direccion"`
	Comuna    string          `json:"comuna"`
	Latitud   decimal.Decimal `json:"latitiud"`
	Longitud  decimal.Decimal `json:"longitud"`
	RegionId  int64           `json:"region_id"`
	ComunaId  int64           `json:"comuna_id"`
	CompanyId int64           `json:"company_id"`
}

func (CreateColegiosReq) TableName() string {
	return "establecimientos" // Nombre de la tabla en la base de datos
}

type UpdateColegiosReq struct {
	ID        string  `json:"-"`
	Codigo    *string `json:"codigo"`
	Nombre    *string `json:"nombre"`
	Direccion *string `json:"direccion"`
	Comuna    *string `json:"comuna"`
	Latitud   *int16  `json:"latitiud"`
	Longitud  *int16  `json:"longitud"`
	RegionId  *int64  `json:"region_id"`
	ComunaId  *int64  `json:"comuna_id"`
	CompanyId *int64  `json:"company_id"`
}

func (UpdateColegiosReq) TableName() string {
	return "establecimientos" // Nombre de la tabla en la base de datos
}

type ColegiosReport struct {
	ID     int64  `json:"id"`
	Codigo string `json:"codigo"`
	Nombre string `json:"nombre"`
}

func (ColegiosReport) TableName() string {
	return "establecimientos" // Nombre de la tabla en la base de datos
}
