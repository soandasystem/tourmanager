package models

import "time"

type PaymentInstallment struct {
	ID            string    `json:"_id,omitempty"`
	PaymentID     int64     `json:"payment_id"`
	InstallmentID int64     `json:"installment_id"`
	AppliedAmount float32   `json:"applied_amount"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

// Resp  response struct
type PaymentInstallmentResp struct {
	ID            string    `json:"id"`
	PaymentID     int64     `json:"payment_id"`
	InstallmentID int64     `json:"installment_id"`
	AppliedAmount float32   `json:"applied_amount"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

func (PaymentInstallmentResp) TableName() string {
	return "payment_installments" // Nombre de la tabla en la base de datos
}

type PaymentInstallmentListResponse struct {
	Items      []PaymentInstallmentResp `json:"items"`
	TotalCount int64                    `json:"totalCount"`
}

// Create---Req  request struct
type CreatePaymentInstallmentReq struct {
	ID            string    `gorm:"primaryKey;autoIncrement"`
	PaymentID     int64     `json:"payment_id"`
	InstallmentID int64     `json:"installment_id"`
	AppliedAmount float32   `json:"applied_amount"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

func (CreatePaymentInstallmentReq) TableName() string {
	return "payment_installments" // Nombre de la tabla en la base de datos
}

type UpdatePaymentInstallmentReq struct {
	ID            string     `json:"-"`
	PaymentID     *int64     `json:"payment_id"`
	InstallmentID *int64     `json:"installment_id"`
	AppliedAmount *float32   `json:"applied_amount"`
	CreatedDate   *time.Time `gorm:"autoCreateTime"`
	UpdatedDate   *time.Time `gorm:"autoUpdateTime"`
}

func (UpdatePaymentInstallmentReq) TableName() string {
	return "payment_installments" // Nombre de la tabla en la base de datos
}

type PaymentInstallmentInf struct {
	ID            int64     `json:"id"`
	PaymentID     int64     `json:"payment_id"`
	InstallmentID int64     `json:"installment_id"`
	AppliedAmount float32   `json:"applied_amount"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

func (PaymentInstallmentInf) TableName() string {
	return "payment_installments" // Nombre de la tabla en la base de datos
}

type PaymentInstallmentInfListResponse struct {
	Items      []PaymentInstallmentInf `json:"items"`
	TotalCount int64                   `json:"totalCount"`
}

type PaymentInstallmentReport struct {
	ID            int64     `json:"id"`
	PaymentID     int64     `json:"payment_id"`
	InstallmentID int64     `json:"installment_id"`
	AppliedAmount float32   `json:"applied_amount"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

func (PaymentInstallmentReport) TableName() string {
	return "payment_installments" // Nombre de la tabla en la base de datos
}
