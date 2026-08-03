package models

type InitFlowPaymentReq struct {
	Monto         int    `json:"mpagar"`
	Identificador string `json:"identificador"`
	ValorCuota    int    `json:"valorcuota"`
	NroCuotas     int    `json:"nrocuotas"`
	FechaInicial  string `json:"fechainicial"`
	CompanyID     int64  `json:"company_id"`
	SaleID        int64  `json:"sale_id"`
	CursoID       int64  `json:"curso_id"`
	UserRut       string `json:"user_rut"`
}

type InitFlowPaymentResp struct {
	RedirectURL string `json:"redirect_url"`
}
