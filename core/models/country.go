package models

type Country struct {
	ID          string `json:"_id,omitempty"`
	Country     string `json:"country"`
	Code        string `json:"code"`
	Countrycode string `json:"countrycode"`
}

type CountryResp struct {
	ID          string `json:"_id,omitempty"`
	Country     string `json:"country"`
	Code        string `json:"code"`
	Countrycode string `json:"countrycode"`
}

func (CountryResp) TableName() string {
	return "country" // Nombre de la tabla en la base de datos
}

type CountryListResponse struct {
	Items      []CountryResp `json:"items"`
	TotalCount int64         `json:"totalCount"`
}

type CreateCountryReq struct {
	ID          string `json:"_id,omitempty"`
	Country     string `json:"country"`
	Code        string `json:"code"`
	Countrycode string `json:"countrycode"`
}

func (CreateCountryReq) TableName() string {
	return "country" // Nombre de la tabla en la base de datos
}

type UpdateCountryReq struct {
	ID          string `json:"_id,omitempty"`
	Country     string `json:"country"`
	Code        string `json:"code"`
	Countrycode string `json:"countrycode"`
}

func (UpdateCountryReq) TableName() string {
	return "country" // Nombre de la tabla en la base de datos
}
