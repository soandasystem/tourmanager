package models

type InitFlowPaymentReq struct {
	Monto         int    `json:"mpagar"`
	Identificador string `json:"identificador"`
	ValorCuota    int    `json:"valorcuota"`
	NroCuotas     int    `json:"nrocuotas"`
	FechaInicial  string `json:"fechainicial"`
	CompanyID     string `json:"company_id"`
	SaleID        string `json:"sale_id"`
	UserRut       string `json:"user_rut"`
}

type InitFlowPaymentResp struct {
	RedirectURL string `json:"redirect_url"`
}
