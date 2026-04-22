package models

import "time"

type Fmedicas struct {
	ID        string    `json:"_id,omitempty"`
	SaleId    int64     `json:"sale_id"`
	Rutalumn  string    `json:"rutalumn"`
	Dato1     string    `json:"dato1"`
	Dato2     string    `json:"dato2"`
	Dato31    time.Time `json:"dato31"`
	Dato32    string    `json:"dato32"`
	Dato4     string    `json:"dato4"`
	Dato5     string    `json:"dato5"`
	Dato6     time.Time `json:"dato6"`
	Dato7     string    `json:"dato7"`
	Dato8     string    `json:"dato8"`
	Dato9     string    `json:"dato9"`
	Dato91    string    `json:"dato91"`
	Dato92    string    `json:"dato92"`
	Dato10    string    `json:"dato10"`
	Dato101   string    `json:"dato101"`
	Dato11    string    `json:"dato11"`
	Dato111   string    `json:"dato111"`
	Dato12    string    `json:"dato12"`
	Dato13    string    `json:"dato13"`
	Dato141   string    `json:"dato141"`
	Dato142   string    `json:"dato142"`
	Dato151   string    `json:"dato151"`
	Dato152   string    `json:"dato152"`
	Dato161   string    `json:"dato161"`
	Dato162   string    `json:"dato162"`
	Dato17    string    `json:"dato17"`
	Dato18    string    `json:"dato18"`
	Dato19    string    `json:"dato19"`
	Dato20    string    `json:"dato20"`
	Dato21    string    `json:"dato21"`
	Dato22    string    `json:"dato22"`
	CompanyId int64     `json:"company_id"`
}

// Resp  response struct
type FmedicaResp struct {
	ID        string    `json:"id"`
	SaleId    int64     `json:"sale_id"`
	Rutalumn  string    `json:"rutalumn"`
	Dato1     string    `json:"dato1"`
	Dato2     string    `json:"dato2"`
	Dato31    time.Time `json:"dato31"`
	Dato32    string    `json:"dato32"`
	Dato4     string    `json:"dato4"`
	Dato5     string    `json:"dato5"`
	Dato6     time.Time `json:"dato6"`
	Dato7     string    `json:"dato7"`
	Dato8     string    `json:"dato8"`
	Dato9     string    `json:"dato9"`
	Dato91    string    `json:"dato91"`
	Dato92    string    `json:"dato92"`
	Dato10    string    `json:"dato10"`
	Dato101   string    `json:"dato101"`
	Dato11    string    `json:"dato11"`
	Dato111   string    `json:"dato111"`
	Dato12    string    `json:"dato12"`
	Dato13    string    `json:"dato13"`
	Dato141   string    `json:"dato141"`
	Dato142   string    `json:"dato142"`
	Dato151   string    `json:"dato151"`
	Dato152   string    `json:"dato152"`
	Dato161   string    `json:"dato161"`
	Dato162   string    `json:"dato162"`
	Dato17    string    `json:"dato17"`
	Dato18    string    `json:"dato18"`
	Dato19    string    `json:"dato19"`
	Dato20    string    `json:"dato20"`
	Dato21    string    `json:"dato21"`
	Dato22    string    `json:"dato22"`
	CompanyId int64     `json:"company_id"`
}

func (FmedicaResp) TableName() string {
	return "fichamedicas" // Nombre de la tabla en la base de datos
}

type FmedicaListResponse struct {
	Items      []FmedicaResp `json:"items"`
	TotalCount int64         `json:"totalCount"`
}

// Create---Req  request struct
type CreateFmedicaReq struct {
	ID        string    `gorm:"primaryKey;autoIncrement"`
	SaleId    int64     `json:"sale_id"`
	Rutalumn  string    `json:"rutalumn"`
	Dato1     string    `json:"dato1"`
	Dato2     string    `json:"dato2"`
	Dato31    time.Time `json:"dato31"`
	Dato32    string    `json:"dato32"`
	Dato4     string    `json:"dato4"`
	Dato5     string    `json:"dato5"`
	Dato6     time.Time `json:"dato6"`
	Dato7     string    `json:"dato7"`
	Dato8     string    `json:"dato8"`
	Dato9     string    `json:"dato9"`
	Dato91    string    `json:"dato91"`
	Dato92    string    `json:"dato92"`
	Dato10    string    `json:"dato10"`
	Dato101   string    `json:"dato101"`
	Dato11    string    `json:"dato11"`
	Dato111   string    `json:"dato111"`
	Dato12    string    `json:"dato12"`
	Dato13    string    `json:"dato13"`
	Dato141   string    `json:"dato141"`
	Dato142   string    `json:"dato142"`
	Dato151   string    `json:"dato151"`
	Dato152   string    `json:"dato152"`
	Dato161   string    `json:"dato161"`
	Dato162   string    `json:"dato162"`
	Dato17    string    `json:"dato17"`
	Dato18    string    `json:"dato18"`
	Dato19    string    `json:"dato19"`
	Dato20    string    `json:"dato20"`
	Dato21    string    `json:"dato21"`
	Dato22    string    `json:"dato22"`
	CompanyId int64     `json:"company_id"`
}

func (CreateFmedicaReq) TableName() string {
	return "fichamedicas" // Nombre de la tabla en la base de datos
}

type UpdateFmedicaReq struct {
	ID        string     `json:"-"`
	SaleId    *int64     `json:"sale_id"`
	Rutalumn  *string    `json:"rutalumn"`
	Dato1     *string    `json:"dato1"`
	Dato2     *string    `json:"dato2"`
	Dato31    *time.Time `json:"dato31"`
	Dato32    *string    `json:"dato32"`
	Dato4     *string    `json:"dato4"`
	Dato5     *string    `json:"dato5"`
	Dato6     *time.Time `json:"dato6"`
	Dato7     *string    `json:"dato7"`
	Dato8     *string    `json:"dato8"`
	Dato9     *string    `json:"dato9"`
	Dato91    *string    `json:"dato91"`
	Dato92    *string    `json:"dato92"`
	Dato10    *string    `json:"dato10"`
	Dato101   *string    `json:"dato101"`
	Dato11    *string    `json:"dato11"`
	Dato111   *string    `json:"dato111"`
	Dato12    *string    `json:"dato12"`
	Dato13    *string    `json:"dato13"`
	Dato141   *string    `json:"dato141"`
	Dato142   *string    `json:"dato142"`
	Dato151   *string    `json:"dato151"`
	Dato152   *string    `json:"dato152"`
	Dato161   *string    `json:"dato161"`
	Dato162   *string    `json:"dato162"`
	Dato17    *string    `json:"dato17"`
	Dato18    *string    `json:"dato18"`
	Dato19    *string    `json:"dato19"`
	Dato20    *string    `json:"dato20"`
	Dato21    *string    `json:"dato21"`
	Dato22    *string    `json:"dato22"`
	CompanyId *int64     `json:"company_id"`
}

func (UpdateFmedicaReq) TableName() string {
	return "fichamedicas" // Nombre de la tabla en la base de datos
}
