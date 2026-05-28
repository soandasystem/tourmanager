package models

import "time"

type Installment struct {
	ID          string    `json:"_id,omitempty"`
	PassengerID int64     `json:"passenger_id"`
	QuotaNumber int       `json:"quota_number"`
	DueDate     time.Time `json:"due_date"`
	Amount      float32   `json:"amount"`
	PaidAmount  float32   `json:"paid_amount"`
	Balance     float32   `json:"balance"`
	Status      string    `json:"status"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
}

// Resp  response struct
type InstallmentResp struct {
	ID          string    `json:"id"`
	PassengerID int64     `json:"passenger_id"`
	QuotaNumber int       `json:"quota_number"`
	DueDate     time.Time `json:"due_date"`
	Amount      float32   `json:"amount"`
	PaidAmount  float32   `json:"paid_amount"`
	Balance     float32   `json:"balance"`
	Status      string    `json:"status"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
}

func (InstallmentResp) TableName() string {
	return "installments" // Nombre de la tabla en la base de datos
}

type InstallmentListResponse struct {
	Items      []InstallmentResp `json:"items"`
	TotalCount int64             `json:"totalCount"`
}

// Create---Req  request struct
type CreateInstallmentReq struct {
	ID          string    `gorm:"primaryKey;autoIncrement"`
	PassengerID int64     `json:"passenger_id"`
	QuotaNumber int       `json:"quota_number"`
	DueDate     time.Time `json:"due_date"`
	Amount      float32   `json:"amount"`
	PaidAmount  float32   `json:"paid_amount"`
	Balance     float32   `json:"balance"`
	Status      string    `json:"status"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
}

func (CreateInstallmentReq) TableName() string {
	return "installments" // Nombre de la tabla en la base de datos
}

type UpdateInstallmentReq struct {
	ID          string     `json:"-"`
	PassengerID *int64     `json:"passenger_id"`
	QuotaNumber *int       `json:"quota_number"`
	DueDate     *time.Time `json:"due_date"`
	Amount      *float32   `json:"amount"`
	PaidAmount  *float32   `json:"paid_amount"`
	Balance     *float32   `json:"balance"`
	Status      *string    `json:"status"`
	CreatedDate *time.Time `gorm:"autoCreateTime"`
	UpdatedDate *time.Time `gorm:"autoUpdateTime"`
}

func (UpdateInstallmentReq) TableName() string {
	return "installments" // Nombre de la tabla en la base de datos
}

type InstallmentInf struct {
	ID          int64     `json:"id"`
	PassengerID int64     `json:"passenger_id"`
	QuotaNumber int       `json:"quota_number"`
	DueDate     time.Time `json:"due_date"`
	Amount      float32   `json:"amount"`
	PaidAmount  float32   `json:"paid_amount"`
	Balance     float32   `json:"balance"`
	Status      string    `json:"status"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
}

func (InstallmentInf) TableName() string {
	return "installments" // Nombre de la tabla en la base de datos
}

type InstallmentInfListResponse struct {
	Items      []InstallmentInf `json:"items"`
	TotalCount int64            `json:"totalCount"`
}

type InstallmentReport struct {
	ID          int64     `json:"id"`
	PassengerID int64     `json:"passenger_id"`
	QuotaNumber int       `json:"quota_number"`
	DueDate     time.Time `json:"due_date"`
	Amount      float32   `json:"amount"`
	PaidAmount  float32   `json:"paid_amount"`
	Balance     float32   `json:"balance"`
	Status      string    `json:"status"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
}

func (InstallmentReport) TableName() string {
	return "installments" // Nombre de la tabla en la base de datos
}
