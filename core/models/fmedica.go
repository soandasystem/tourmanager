package models

import "time"

type Fmedicas struct {
	ID                        string    `json:"_id,omitempty"`
	SaleId                    int64     `json:"sale_id"`
	CursoId                   int64     `json:"curso_id"`
	CompanyId                 int64     `json:"company_id"`
	GrupoSanguineo            string    `json:"grupo_sanguineo"`
	Edad                      int64     `json:"edad"`
	Peso                      float64   `json:"peso"`
	Estatura                  float64   `json:"estatura"`
	Hipertension              bool      `json:"hipertension"`
	Diabetes                  bool      `json:"diabetes"`
	Asma                      bool      `json:"asma"`
	Epilepsia                 bool      `json:"epilepsia"`
	Arritmias                 bool      `json:"arritmias"`
	EnfermedadesCardiacas     bool      `json:"enfermedades_cardiacas"`
	EnfermedadesRespiratorias bool      `json:"enfermedades_respiratorias"`
	EnfermedadesRenales       bool      `json:"enfermedades_renales"`
	OtrasEnfermedades         string    `json:"otras_enfermedades"`
	BajoTratamiento           string    `json:"bajo_tratamiento"`
	Diagnostico               string    `json:"diagnostico"`
	MedicamentosTratamiento   string    `json:"medicamentos_tratamiento"`
	DosisTratamiento          string    `json:"dosis_tratamiento"`
	AlergiaMedicamentos       bool      `json:"alergia_medicamentos"`
	AlergiaAlimentos          bool      `json:"alergia_alimentos"`
	AlergiaInsectos           bool      `json:"alergia_insectos"`
	AlergiaOtras              string    `json:"alergia_otras"`
	DificultadMovilidad       string    `json:"dificultad_movilidad"`
	AsistenciaMovilidad       string    `json:"asistencia_movilidad"`
	AsistenciaEspecial        string    `json:"asistencia_especial"`
	Vegetariano               bool      `json:"vegetariano"`
	Vegano                    bool      `json:"vegano"`
	Celiaco                   bool      `json:"celiaco"`
	IntoleranciaLactosa       bool      `json:"intolerancia_lactosa"`
	DiabeticoAlim             bool      `json:"diabetico_alim"`
	OtraRestriccionAlim       string    `json:"otra_restriccion_alim"`
	ContactoNombre            string    `json:"contacto_nombre"`
	ContactoRelacion          string    `json:"contacto_relacion"`
	ContactoTelefono1         string    `json:"contacto_telefono1"`
	ContactoTelefono2         string    `json:"contacto_telefono2"`
	Observaciones             string    `json:"observaciones"`
	Author                    string    `json:"author"`
	CreatedDate               time.Time `json:"created_date"`
	UpdatedDate               time.Time `json:"updated_date"`
}

// Resp  response struct
type FmedicaResp struct {
	ID                        string    `json:"id"`
	SaleId                    int64     `json:"sale_id"`
	CursoId                   int64     `json:"curso_id"`
	CompanyId                 int64     `json:"company_id"`
	GrupoSanguineo            string    `json:"grupo_sanguineo"`
	Edad                      int64     `json:"edad"`
	Peso                      float64   `json:"peso"`
	Estatura                  float64   `json:"estatura"`
	Hipertension              bool      `json:"hipertension"`
	Diabetes                  bool      `json:"diabetes"`
	Asma                      bool      `json:"asma"`
	Epilepsia                 bool      `json:"epilepsia"`
	Arritmias                 bool      `json:"arritmias"`
	EnfermedadesCardiacas     bool      `json:"enfermedades_cardiacas"`
	EnfermedadesRespiratorias bool      `json:"enfermedades_respiratorias"`
	EnfermedadesRenales       bool      `json:"enfermedades_renales"`
	OtrasEnfermedades         string    `json:"otras_enfermedades"`
	BajoTratamiento           string    `json:"bajo_tratamiento"`
	Diagnostico               string    `json:"diagnostico"`
	MedicamentosTratamiento   string    `json:"medicamentos_tratamiento"`
	DosisTratamiento          string    `json:"dosis_tratamiento"`
	AlergiaMedicamentos       bool      `json:"alergia_medicamentos"`
	AlergiaAlimentos          bool      `json:"alergia_alimentos"`
	AlergiaInsectos           bool      `json:"alergia_insectos"`
	AlergiaOtras              string    `json:"alergia_otras"`
	DificultadMovilidad       string    `json:"dificultad_movilidad"`
	AsistenciaMovilidad       string    `json:"asistencia_movilidad"`
	AsistenciaEspecial        string    `json:"asistencia_especial"`
	Vegetariano               bool      `json:"vegetariano"`
	Vegano                    bool      `json:"vegano"`
	Celiaco                   bool      `json:"celiaco"`
	IntoleranciaLactosa       bool      `json:"intolerancia_lactosa"`
	DiabeticoAlim             bool      `json:"diabetico_alim"`
	OtraRestriccionAlim       string    `json:"otra_restriccion_alim"`
	ContactoNombre            string    `json:"contacto_nombre"`
	ContactoRelacion          string    `json:"contacto_relacion"`
	ContactoTelefono1         string    `json:"contacto_telefono1"`
	ContactoTelefono2         string    `json:"contacto_telefono2"`
	Observaciones             string    `json:"observaciones"`
	Author                    string    `json:"author"`
	CreatedDate               time.Time `json:"created_date"`
	UpdatedDate               time.Time `json:"updated_date"`
}

func (FmedicaResp) TableName() string {
	return "fmedicas" // Nombre de la tabla en la base de datos
}

type FmedicaListResponse struct {
	Items      []FmedicaResp `json:"items"`
	TotalCount int64         `json:"totalCount"`
}

// Create---Req  request struct
type CreateFmedicaReq struct {
	ID                        string    `gorm:"primaryKey;autoIncrement"`
	SaleId                    int64     `json:"sale_id"`
	CursoId                   int64     `json:"curso_id"`
	CompanyId                 int64     `json:"company_id"`
	GrupoSanguineo            string    `json:"grupo_sanguineo"`
	Edad                      int64     `json:"edad"`
	Peso                      float64   `json:"peso"`
	Estatura                  float64   `json:"estatura"`
	Hipertension              bool      `json:"hipertension"`
	Diabetes                  bool      `json:"diabetes"`
	Asma                      bool      `json:"asma"`
	Epilepsia                 bool      `json:"epilepsia"`
	Arritmias                 bool      `json:"arritmias"`
	EnfermedadesCardiacas     bool      `json:"enfermedades_cardiacas"`
	EnfermedadesRespiratorias bool      `json:"enfermedades_respiratorias"`
	EnfermedadesRenales       bool      `json:"enfermedades_renales"`
	OtrasEnfermedades         string    `json:"otras_enfermedades"`
	BajoTratamiento           string    `json:"bajo_tratamiento"`
	Diagnostico               string    `json:"diagnostico"`
	MedicamentosTratamiento   string    `json:"medicamentos_tratamiento"`
	DosisTratamiento          string    `json:"dosis_tratamiento"`
	AlergiaMedicamentos       bool      `json:"alergia_medicamentos"`
	AlergiaAlimentos          bool      `json:"alergia_alimentos"`
	AlergiaInsectos           bool      `json:"alergia_insectos"`
	AlergiaOtras              string    `json:"alergia_otras"`
	DificultadMovilidad       string    `json:"dificultad_movilidad"`
	AsistenciaMovilidad       string    `json:"asistencia_movilidad"`
	AsistenciaEspecial        string    `json:"asistencia_especial"`
	Vegetariano               bool      `json:"vegetariano"`
	Vegano                    bool      `json:"vegano"`
	Celiaco                   bool      `json:"celiaco"`
	IntoleranciaLactosa       bool      `json:"intolerancia_lactosa"`
	DiabeticoAlim             bool      `json:"diabetico_alim"`
	OtraRestriccionAlim       string    `json:"otra_restriccion_alim"`
	ContactoNombre            string    `json:"contacto_nombre"`
	ContactoRelacion          string    `json:"contacto_relacion"`
	ContactoTelefono1         string    `json:"contacto_telefono1"`
	ContactoTelefono2         string    `json:"contacto_telefono2"`
	Observaciones             string    `json:"observaciones"`
	Author                    string    `json:"author"`
	CreatedDate               time.Time `json:"created_date"`
	UpdatedDate               time.Time `json:"updated_date"`
}

func (CreateFmedicaReq) TableName() string {
	return "fmedicas" // Nombre de la tabla en la base de datos
}

type UpdateFmedicaReq struct {
	ID                        string     `json:"-"`
	SaleId                    *int64     `json:"sale_id"`
	CursoId                   *int64     `json:"curso_id"`
	CompanyId                 *int64     `json:"company_id"`
	GrupoSanguineo            *string    `json:"grupo_sanguineo"`
	Edad                      *int64     `json:"edad"`
	Peso                      *float64   `json:"peso"`
	Estatura                  *float64   `json:"estatura"`
	Hipertension              *bool      `json:"hipertension"`
	Diabetes                  *bool      `json:"diabetes"`
	Asma                      *bool      `json:"asma"`
	Epilepsia                 *bool      `json:"epilepsia"`
	Arritmias                 *bool      `json:"arritmias"`
	EnfermedadesCardiacas     *bool      `json:"enfermedades_cardiacas"`
	EnfermedadesRespiratorias *bool      `json:"enfermedades_respiratorias"`
	EnfermedadesRenales       *bool      `json:"enfermedades_renales"`
	OtrasEnfermedades         *string    `json:"otras_enfermedades"`
	BajoTratamiento           *string    `json:"bajo_tratamiento"`
	Diagnostico               *string    `json:"diagnostico"`
	MedicamentosTratamiento   *string    `json:"medicamentos_tratamiento"`
	DosisTratamiento          *string    `json:"dosis_tratamiento"`
	AlergiaMedicamentos       *bool      `json:"alergia_medicamentos"`
	AlergiaAlimentos          *bool      `json:"alergia_alimentos"`
	AlergiaInsectos           *bool      `json:"alergia_insectos"`
	AlergiaOtras              *string    `json:"alergia_otras"`
	DificultadMovilidad       *string    `json:"dificultad_movilidad"`
	AsistenciaMovilidad       *string    `json:"asistencia_movilidad"`
	AsistenciaEspecial        *string    `json:"asistencia_especial"`
	Vegetariano               *bool      `json:"vegetariano"`
	Vegano                    *bool      `json:"vegano"`
	Celiaco                   *bool      `json:"celiaco"`
	IntoleranciaLactosa       *bool      `json:"intolerancia_lactosa"`
	DiabeticoAlim             *bool      `json:"diabetico_alim"`
	OtraRestriccionAlim       *string    `json:"otra_restriccion_alim"`
	ContactoNombre            *string    `json:"contacto_nombre"`
	ContactoRelacion          *string    `json:"contacto_relacion"`
	ContactoTelefono1         *string    `json:"contacto_telefono1"`
	ContactoTelefono2         *string    `json:"contacto_telefono2"`
	Observaciones             *string    `json:"observaciones"`
	Author                    *string    `json:"author"`
	CreatedDate               *time.Time `json:"created_date"`
	UpdatedDate               *time.Time `json:"updated_date"`
}

func (UpdateFmedicaReq) TableName() string {
	return "fmedicas" // Nombre de la tabla en la base de datos
}
