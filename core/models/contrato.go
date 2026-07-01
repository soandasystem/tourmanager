package models

// ContratoReq contiene todos los campos enviados por el cliente para generar el contrato
type ContratoReq struct {
	VtaDia           string `json:"vtaDia"`
	VtaMes           string `json:"vtaMes"`
	VtaAgno          string `json:"vtaAgno"`
	Rute             string `json:"rute"`
	RSocial          string `json:"rsocial"`
	NFantasia        string `json:"nfantasia"`
	RLegal           string `json:"rlegal"`
	NLegal           string `json:"nlegal"`
	EDireccion       string `json:"edireccion"`
	Colegio          string `json:"colegio"`
	Comuna           string `json:"comuna"`
	IdCurso          string `json:"idcurso"`
	Programa         string `json:"programa"`
	Reserva          string `json:"reserva"`
	NombreApod       string `json:"nombreapod"`
	NombreAlumno     string `json:"nombrealumno"`
	RutApod          string `json:"rutapod"`
	CorreoApod       string `json:"correoapod"`
	FonoApod         string `json:"fonoapod"`
	Observacion      string `json:"observacion"`
	VPrograma        string `json:"vprograma"`
	Tc               string `json:"tc"`
	Liberados        string `json:"liberados"`
	FSalida          string `json:"fsalida"`
	FSalidaMes       string `json:"fsalidames"`
	FSalidaAgno      string `json:"fsalidaaño"`
	FSalidaDia       string `json:"fsalidadia"`
	FPago            string `json:"fpago"`
	TypeSale         string `json:"type_sale"`
	TemplateFilename string `json:"template_filename"` // Nombre del archivo template DOCX en B2
}

// ContratoTempResp respuesta de Fase 1 — DOCX temporal generado
type ContratoTempResp struct {
	SessionID string `json:"session_id"` // UUID de sesión para referenciar en Fase 2
	TempFile  string `json:"temp_file"`  // Ruta del DOCX temporal
	Message   string `json:"message"`
}

// ContratoFirmaReq request de Fase 2 — firma en base64 + session_id
type ContratoFirmaReq struct {
	SessionID     string `json:"session_id"`       // UUID devuelto en Fase 1
	FirmaBase64   string `json:"firma_base64"`     // Imagen de la firma en base64 (PNG o JPEG)
	FileNameFirma string `json:"file_name_firma"` // Nombre a colocar cuando cree el pdf y antes de subirlo
}

// ContratoPDFResp respuesta de Fase 2 — PDF definitivo generado
type ContratoPDFResp struct {
	PDFFile string `json:"pdf_file"` // Ruta del PDF definitivo
	Message string `json:"message"`
}
