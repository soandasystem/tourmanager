package models

import "time"

type Payments struct {
	ID              string    `json:"_id,omitempty"`
	PassengerId     int64     `json:"passenger_id"`
	Amount          float32   `json:"amount"`
	PaymentMethod   string    `json:"payment_method"`
	PaymentDate     time.Time `json:"payment_date"`
	Identifier      string    `json:"identifier"`
	Notes           string    `json:"notes"`
	TransactionRef  string    `json:"transaction_ref"`
	TransactionType string    `json:"transaction_type"`
	CardNumber      string    `json:"card_number"`
	AuthCode        string    `json:"auth_code"`
	AuthDate        time.Time `json:"auth_date"`
	PaymentToken    string    `json:"payment_token"`
	CompanyId       int64     `json:"company_id"`
	SaleId          int64     `json:"sale_id"`
	CreatedDate     time.Time `gorm:"autoCreateTime"`
	UpdatedDate     time.Time `gorm:"autoUpdateTime"`
}

// Resp  response struct
type PaymentResp struct {
	ID              string    `json:"id"`
	PassengerId     int64     `json:"passenger_id"`
	Amount          float32   `json:"amount"`
	PaymentMethod   string    `json:"payment_method"`
	PaymentDate     time.Time `json:"payment_date"`
	Identifier      string    `json:"identifier"`
	Notes           string    `json:"notes"`
	TransactionRef  string    `json:"transaction_ref"`
	TransactionType string    `json:"transaction_type"`
	CardNumber      string    `json:"card_number"`
	AuthCode        string    `json:"auth_code"`
	AuthDate        time.Time `json:"auth_date"`
	PaymentToken    string    `json:"payment_token"`
	CompanyId       int64     `json:"company_id"`
	SaleId          int64     `json:"sale_id"`
	CreatedDate     time.Time `gorm:"autoCreateTime"`
	UpdatedDate     time.Time `gorm:"autoUpdateTime"`
}

func (PaymentResp) TableName() string {
	return "payments" // Nombre de la tabla en la base de datos
}

type PaymentListResponse struct {
	Items      []PaymentResp `json:"items"`
	TotalCount int64         `json:"totalCount"`
}

// Create---Req  request struct
type CreatePaymentReq struct {
	ID              string    `gorm:"primaryKey;autoIncrement"`
	PassengerId     int64     `json:"passenger_id"`
	Amount          float32   `json:"amount"`
	PaymentMethod   string    `json:"payment_method"`
	PaymentDate     time.Time `json:"payment_date"`
	Identifier      string    `json:"identifier"`
	Notes           string    `json:"notes"`
	TransactionRef  string    `json:"transaction_ref"`
	TransactionType string    `json:"transaction_type"`
	CardNumber      string    `json:"card_number"`
	AuthCode        string    `json:"auth_code"`
	AuthDate        time.Time `json:"auth_date"`
	PeymentToken    string    `json:"payment_token"`
	CompanyId       int64     `json:"company_id"`
	SaleId          int64     `json:"sale_id"`
	CreatedDate     time.Time `gorm:"autoCreateTime"`
	UpdatedDate     time.Time `gorm:"autoUpdateTime"`
}

func (CreatePaymentReq) TableName() string {
	return "payments" // Nombre de la tabla en la base de datos
}

type UpdatePaymentReq struct {
	ID              string     `json:"-"`
	PassengerId     *int64     `json:"passenger_id"`
	Amount          *float32   `json:"amount"`
	PaymentMethod   *string    `json:"payment_method"`
	PaymentDate     *time.Time `json:"payment_date"`
	Identifier      *string    `json:"identifier"`
	Notes           *string    `json:"notes"`
	TransactionRef  *string    `json:"transaction_ref"`
	TransactionType *string    `json:"transaction_type"`
	CardNumber      *string    `json:"card_number"`
	AuthCode        *string    `json:"auth_code"`
	AuthDate        *time.Time `json:"auth_date"`
	PaymentToken    *string    `json:"payment_token"`
	CompanyId       *int64     `json:"company_id"`
	SaleId          *int64     `json:"sale_id"`
	CreatedDate     *time.Time `gorm:"autoCreateTime"`
	UpdatedDate     *time.Time `gorm:"autoUpdateTime"`
}

func (UpdatePaymentReq) TableName() string {
	return "payments" // Nombre de la tabla en la base de datos
}

type PaymentInf struct {
	ID              int64     `json:"id"`
	PassengerId     int64     `json:"passenger_id"`
	Amount          float32   `json:"amount"`
	PaymentMethod   string    `json:"payment_method"`
	PaymentDate     time.Time `json:"payment_date"`
	Identifier      string    `json:"identifier"`
	Notes           string    `json:"notes"`
	TransactionRef  string    `json:"transaction_ref"`
	TransactionType string    `json:"transaction_type"`
	CardNumber      string    `json:"card_number"`
	AuthCode        string    `json:"auth_code"`
	AuthDate        time.Time `json:"auth_date"`
	PaymentToken    string    `json:"payment_token"`
	CompanyId       int64     `json:"company_id"`
	SaleId          int64     `json:"sale_id"`
	CreatedDate     time.Time `gorm:"autoCreateTime"`
	UpdatedDate     time.Time `gorm:"autoUpdateTime"`
}

func (PaymentInf) TableName() string {
	return "payments" // Nombre de la tabla en la base de datos
}

type PaymentInfListResponse struct {
	Items      []PaymentInf `json:"items"`
	TotalCount int64        `json:"totalCount"`
}

type PaymentReport struct {
	ID              int64     `json:"id"`
	PassengerId     int64     `json:"passenger_id"`
	Amount          float32   `json:"amount"`
	PaymentMethod   string    `json:"payment_method"`
	PaymentDate     time.Time `json:"payment_date"`
	Identifier      string    `json:"identifier"`
	Notes           string    `json:"notes"`
	TransactionRef  string    `json:"transaction_ref"`
	TransactionType string    `json:"transaction_type"`
	CardNumber      string    `json:"card_number"`
	AuthCode        string    `json:"auth_code"`
	AuthDate        time.Time `json:"auth_date"`
	PaymentToken    string    `json:"payment_token"`
	CompanyId       int64     `json:"company_id"`
	SaleId          int64     `json:"sale_id"`
	CreatedDate     time.Time `gorm:"autoCreateTime"`
	UpdatedDate     time.Time `gorm:"autoUpdateTime"`
}

func (PaymentReport) TableName() string {
	return "payments" // Nombre de la tabla en la base de datos
}
