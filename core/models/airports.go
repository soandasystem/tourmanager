package models

type Airports struct {
	ID        string  `json:"_id,omitempty"`
	Icao      string  `json:"icao,omitempty"`
	Iata      string  `json:"iata,omitempty"`
	Name      string  `json:"name,omitempty"`
	City      string  `json:"city,omitempty"`
	State     string  `json:"state,omitempty"`
	Country   string  `json:"country,omitempty"`
	Elevation int     `json:"elevation,omitempty"`
	Lat       float32 `json:"lat,omitempty"`
	Lon       float32 `json:"lon,omitempty"`
	Tz        string  `json:"tz,omitempty"`
}

type AirportsResp struct {
	ID        string  `json:"_id,omitempty"`
	Icao      string  `json:"icao,omitempty"`
	Iata      string  `json:"iata,omitempty"`
	Name      string  `json:"name,omitempty"`
	City      string  `json:"city,omitempty"`
	State     string  `json:"state,omitempty"`
	Country   string  `json:"country,omitempty"`
	Elevation int     `json:"elevation,omitempty"`
	Lat       float32 `json:"lat,omitempty"`
	Lon       float32 `json:"lon,omitempty"`
	Tz        string  `json:"tz,omitempty"`
}

func (AirportsResp) TableName() string {
	return "airports" // Nombre de la tabla en la base de datos
}

type AirportsListResponse struct {
	Items      []AirportsResp `json:"items"`
	TotalCount int64          `json:"totalCount"`
}

type CreateAirportsReq struct {
	ID        string  `json:"_id,omitempty"`
	Icao      string  `json:"icao,omitempty"`
	Iata      string  `json:"iata,omitempty"`
	Name      string  `json:"name,omitempty"`
	City      string  `json:"city,omitempty"`
	State     string  `json:"state,omitempty"`
	Country   string  `json:"country,omitempty"`
	Elevation int     `json:"elevation,omitempty"`
	Lat       float32 `json:"lat,omitempty"`
	Lon       float32 `json:"lon,omitempty"`
	Tz        string  `json:"tz,omitempty"`
}

func (CreateAirportsReq) TableName() string {
	return "airports" // Nombre de la tabla en la base de datos
}

type UpdateAirportsReq struct {
	ID        string   `json:"_id,omitempty"`
	Icao      *string  `json:"icao,omitempty"`
	Iata      *string  `json:"iata,omitempty"`
	Name      *string  `json:"name,omitempty"`
	City      *string  `json:"city,omitempty"`
	State     *string  `json:"state,omitempty"`
	Country   *string  `json:"country,omitempty"`
	Elevation *int     `json:"elevation,omitempty"`
	Lat       *float32 `json:"lat,omitempty"`
	Lon       *float32 `json:"lon,omitempty"`
	Tz        *string  `json:"tz,omitempty"`
}

func (UpdateAirportsReq) TableName() string {
	return "airports" // Nombre de la tabla en la base de datos
}

type AirportsInf struct {
	City      string  `json:"city,omitempty"`
	ID        string  `json:"_id,omitempty"`
	Icao      string  `json:"icao,omitempty"`
	Iata      string  `json:"iata,omitempty"`
	Name      string  `json:"name,omitempty"`
	State     string  `json:"state,omitempty"`
	Country   string  `json:"country,omitempty"`
	Elevation int     `json:"elevation,omitempty"`
	Lat       float32 `json:"lat,omitempty"`
	Lon       float32 `json:"lon,omitempty"`
	Tz        string  `json:"tz,omitempty"`
}

func (AirportsInf) TableName() string {
	return "airports" // Nombre de la tabla en la base de datos
}

type AirportsInfResponse struct {
	Items      []AirportsInf `json:"items"`
	TotalCount int64         `json:"totalCount"`
}
