package models

import "time"

type Payments struct {
	ID            string    `json:"_id,omitempty"`
	PassengerID   int64     `json:"passenger_id"`
	Amount        float32   `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	PaymentDate   time.Time `json:"payment_date"`
	Reference     string    `json:"reference"`
	Notes         string    `json:"notes"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

// Resp  response struct
type PaymentResp struct {
	ID            string    `json:"id"`
	PassengerID   int64     `json:"passenger_id"`
	Amount        float32   `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	PaymentDate   time.Time `json:"payment_date"`
	Reference     string    `json:"reference"`
	Notes         string    `json:"notes"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
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
	ID            string    `gorm:"primaryKey;autoIncrement"`
	PassengerID   int64     `json:"passenger_id"`
	Amount        float32   `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	PaymentDate   time.Time `json:"payment_date"`
	Reference     string    `json:"reference"`
	Notes         string    `json:"notes"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

func (CreatePaymentReq) TableName() string {
	return "payments" // Nombre de la tabla en la base de datos
}

type UpdatePaymentReq struct {
	ID            string     `json:"-"`
	PassengerID   *int64     `json:"passenger_id"`
	Amount        *float32   `json:"amount"`
	PaymentMethod *string    `json:"payment_method"`
	PaymentDate   *time.Time `json:"payment_date"`
	Reference     *string    `json:"reference"`
	Notes         *string    `json:"notes"`
	CreatedDate   *time.Time `gorm:"autoCreateTime"`
	UpdatedDate   *time.Time `gorm:"autoUpdateTime"`
}

func (UpdatePaymentReq) TableName() string {
	return "payments" // Nombre de la tabla en la base de datos
}

type PaymentInf struct {
	ID            int64     `json:"id"`
	PassengerID   int64     `json:"passenger_id"`
	Amount        float32   `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	PaymentDate   time.Time `json:"payment_date"`
	Reference     string    `json:"reference"`
	Notes         string    `json:"notes"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

func (PaymentInf) TableName() string {
	return "payments" // Nombre de la tabla en la base de datos
}

type PaymentInfListResponse struct {
	Items      []PaymentInf `json:"items"`
	TotalCount int64        `json:"totalCount"`
}

type PaymentReport struct {
	ID            int64     `json:"id"`
	PassengerID   int64     `json:"passenger_id"`
	Amount        float32   `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	PaymentDate   time.Time `json:"payment_date"`
	Reference     string    `json:"reference"`
	Notes         string    `json:"notes"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

func (PaymentReport) TableName() string {
	return "payments" // Nombre de la tabla en la base de datos
}
