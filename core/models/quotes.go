package models

import "time"

type Quote struct {
	ID                string    `json:"_id,omitempty"`
	Fecha             time.Time `json:"fecha"`
	Identificador     string    `json:"identificador"`
	SellerId          int64     `json:"seller_id"`
	EstablecimientoId int64     `json:"establecimiento_id"`
	Curso             int       `json:"curso"`
	Idcurso           string    `json:"idcurso"`
	Pasajeros         int       `json:"pasajeros"`
	ProgramaId        int64     `json:"programa_id"`
	Subtotal          int       `json:"subtotal"`
	Desc              int       `json:"desc"`
	Vprograma         int       `json:"vprograma"`
	Liberados         int       `json:"liberados"`
	Tipocambio        float32   `json:"tipocambio"`
	Contacto          string    `json:"contacto"`
	Contactofono      string    `json:"contactofono"`
	Contactoemail     string    `json:"contactoemail"`
	Estado            string    `json:"estado"`
	Obsestado         string    `json:"obsestado"`
	CompanyId         int64     `json:"company_id"`
	FromQuote         int64     `json:"from_quote"`
	SaleId            int64     `json:"sale_id"`
	Author            string    `json:"author"`
	TypeSale          string    `json:"type_sale"`
	CreatedDate       time.Time `gorm:"autoCreateTime"`
	UpdatedDate       time.Time `gorm:"autoCreateTime"`
}

type QuoteResp struct {
	ID                string         `json:"id"`
	Fecha             time.Time      `json:"fecha"`
	Identificador     string         `json:"identificador"`
	SellerId          int64          `json:"seller_id"`
	Users             UsersReport    `json:"users" gorm:"foreignKey:SellerId;references:ID"`
	EstablecimientoId int64          `json:"establecimiento_id"`
	Colegios          ColegiosReport `json:"colegios" gorm:"foreignKey:EstablecimientoId;references:ID"`
	Curso             int            `json:"curso"`
	Idcurso           string         `json:"idcurso"`
	Pasajeros         int            `json:"pasajeros"`
	ProgramaId        int64          `json:"programa_id"`
	Programs          ProgramsOthers `json:"programs" gorm:"foreignKey:ProgramaId;references:ID"`
	Subtotal          int            `json:"subtotal"`
	Desc              int            `json:"desc"`
	Vprograma         int            `json:"vprograma"`
	Liberados         int            `json:"liberados"`
	Tipocambio        float32        `json:"tipocambio"`
	Contacto          string         `json:"contacto"`
	Contactofono      string         `json:"contactofono"`
	Contactoemail     string         `json:"contactoemail"`
	Estado            string         `json:"estado"`
	Obsestado         string         `json:"obsestado"`
	CompanyId         int64          `json:"company_id"`
	FromQuote         int64          `json:"from_quote"`
	SaleId            int64          `json:"sale_id"`
	Author            string         `json:"author"`
	TypeSale          string         `json:"type_sale"`
	CreatedDate       time.Time      `gorm:"autoCreateTime"`
	UpdatedDate       time.Time      `gorm:"autoCreateTime"`
}

func (QuoteResp) TableName() string {
	return "quotes" // Nombre de la tabla en la base de datos
}

type QuoteListResponse struct {
	Items      []QuoteResp `json:"items"`
	TotalCount int64       `json:"totalCount"`
}

type CreateQuoteReq struct {
	ID                string    `gorm:"primaryKey;autoIncrement"`
	Fecha             time.Time `json:"fecha"`
	Identificador     string    `json:"identificador"`
	SellerId          int64     `json:"seller_id"`
	EstablecimientoId int64     `json:"establecimiento_id"`
	Curso             int       `json:"curso"`
	Idcurso           string    `json:"idcurso"`
	Pasajeros         int       `json:"pasajeros"`
	ProgramaId        int64     `json:"programa_id"`
	Subtotal          int       `json:"subtotal"`
	Desc              int       `json:"desc"`
	Vprograma         int       `json:"vprograma"`
	Liberados         int       `json:"liberados"`
	Tipocambio        float32   `json:"tipocambio"`
	Contacto          string    `json:"contacto"`
	Contactofono      string    `json:"contactofono"`
	Contactoemail     string    `json:"contactoemail"`
	Estado            string    `json:"estado"`
	Obsestado         string    `json:"obsestado"`
	CompanyId         int64     `json:"company_id"`
	FromQuote         int64     `json:"from_quote"`
	SaleId            int64     `json:"sale_id"`
	Author            string    `json:"author"`
	TypeSale          string    `json:"type_sale"`
	CreatedDate       time.Time `gorm:"autoCreateTime"`
	UpdatedDate       time.Time `gorm:"autoCreateTime"`
}

func (CreateQuoteReq) TableName() string {
	return "quotes" // Nombre de la tabla en la base de datos
}

type UpdateQuoteReq struct {
	ID                string     `json:"-"`
	Fecha             *time.Time `json:"fecha"`
	Identificador     *string    `json:"identificador"`
	SellerId          *int64     `json:"seller_id"`
	EstablecimientoId *int64     `json:"establecimiento_id"`
	Curso             *int       `json:"curso"`
	Idcurso           *string    `json:"idcurso"`
	Pasajeros         *int       `json:"pasajeros"`
	ProgramaId        *int64     `json:"programa_id"`
	Subtotal          *int       `json:"subtotal"`
	Desc              *int       `json:"desc"`
	Vprograma         *int       `json:"vprograma"`
	Liberados         int        `json:"liberados"`
	Tipocambio        *float32   `json:"tipocambio"`
	Contacto          *string    `json:"contacto"`
	Contactofono      *string    `json:"contactofono"`
	Contactoemail     *string    `json:"contactoemail"`
	Estado            *string    `json:"estado"`
	Obsestado         *string    `json:"obsestado"`
	CompanyId         *int64     `json:"company_id"`
	FromQuote         *int64     `json:"from_quote"`
	SaleId            *int64     `json:"sale_id"`
	Author            *string    `json:"author"`
	TypeSale          *string    `json:"type_sale"`
	CreatedDate       *time.Time `gorm:"autoCreateTime"`
	UpdatedDate       *time.Time `gorm:"autoCreateTime"`
}

func (UpdateQuoteReq) TableName() string {
	return "quotes" // Nombre de la tabla en la base de datos
}

type QuoteInforme struct {
	ID                string         `json:"id"`
	Fecha             time.Time      `json:"fecha"`
	Identificador     string         `json:"identificador"`
	SellerId          int64          `json:"seller_id"`
	Users             UsersReport    `json:"users" gorm:"foreignKey:SellerId;references:ID"`
	EstablecimientoId int64          `json:"establecimiento_id"`
	Colegios          ColegiosReport `json:"colegios" gorm:"foreignKey:EstablecimientoId;references:ID"`
	Curso             int            `json:"curso"`
	Idcurso           string         `json:"idcurso"`
	Pasajeros         int            `json:"pasajeros"`
	ProgramaId        int64          `json:"programa_id"`
	Programs          ProgramsReport `json:"programs" gorm:"foreignKey:ProgramaId;references:ID"`
	Subtotal          int            `json:"subtotal"`
	Desc              int            `json:"desc"`
	Vprograma         int            `json:"vprograma"`
	Liberados         int            `json:"liberados"`
	Tipocambio        float32        `json:"tipocambio"`
	Contacto          string         `json:"contacto"`
	Contactofono      string         `json:"contactofono"`
	Contactoemail     string         `json:"contactoemail"`
	Estado            string         `json:"estado"`
	Obsestado         string         `json:"obsestado"`
	CompanyId         int64          `json:"company_id"`
	FromQuote         int64          `json:"from_quote"`
	SaleId            int64          `json:"sale_id"`
	Author            string         `json:"author"`
	TypeSale          string         `json:"type_sale"`
	CreatedDate       time.Time      `gorm:"autoCreateTime"`
	UpdatedDate       time.Time      `gorm:"autoCreateTime"`
}

func (QuoteInforme) TableName() string {
	return "quotes" // Nombre de la tabla en la base de datos
}

type QuoteInfListResponse struct {
	Items      []QuoteInforme `json:"items"`
	TotalCount int64          `json:"totalCount"`
}
