package models

type Vouchers struct {
	ID        string `json:"_id,omitempty"`
	SaleId    int64  `json:"sale_id"`
	Voucher   string `json:"voucher"`
	Used      int    `json:"used"`
	CompanyId int64  `json:"company_id"`
}

// Resp  response struct
type VoucherResp struct {
	ID        string `json:"id"`
	SaleId    string `json:"sale_id"`
	Voucher   string `json:"voucher"`
	Used      int    `json:"used"`
	CompanyId int64  `json:"company_id"`
}

func (VoucherResp) TableName() string {
	return "voucher" // Nombre de la tabla en la base de datos
}

type VoucherListResponse struct {
	Items      []VoucherResp `json:"items"`
	TotalCount int64         `json:"totalCount"`
}

// Create---Req  request struct
type CreateVoucherReq struct {
	ID        string `gorm:"primaryKey;autoIncrement"`
	SaleId    int64  `json:"sale_id"`
	Voucher   string `json:"voucher"`
	Used      int    `json:"used"`
	CompanyId int64  `json:"company_id"`
}

func (CreateVoucherReq) TableName() string {
	return "voucher" // Nombre de la tabla en la base de datos
}

type UpdateVoucherReq struct {
	ID        string `json:"-"`
	SaleId    int64  `json:"sale_id"`
	Voucher   string `json:"voucher"`
	Used      int    `json:"used"`
	CompanyId int64  `json:"company_id"`
}

func (UpdateVoucherReq) TableName() string {
	return "voucher" // Nombre de la tabla en la base de datos
}
