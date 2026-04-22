package models

type Programad struct {
	ID         string  `json:"_id,omitempty"`
	ProgramaId int64   `json:"programa_id"`
	Desde      int     `json:"desde"`
	Hasta      int     `json:"hasta"`
	Valor      float32 `json:"valor"`
	Liberado   int     `json:"liberado"`
}

type ProgramadResp struct {
	ID         string  `json:"id"`
	ProgramaId int64   `json:"programa_id"`
	Desde      int     `json:"desde"`
	Hasta      int     `json:"hasta"`
	Valor      float32 `json:"valor"`
	Liberado   int     `json:"liberado"`
}

func (ProgramadResp) TableName() string {
	return "programad" // Nombre de la tabla en la base de datos
}

type ProgramadListResponse struct {
	Items      []ProgramadResp `json:"items"`
	TotalCount int64           `json:"totalCount"`
}

type CreateProgramadReq struct {
	ID         string  `json:"id" gorm:"primaryKey;autoIncrement"`
	ProgramaId int64   `json:"programa_id"`
	Desde      int     `json:"desde"`
	Hasta      int     `json:"hasta"`
	Valor      float32 `json:"valor"`
	Liberado   int     `json:"liberado"`
}

func (CreateProgramadReq) TableName() string {
	return "programad" // Nombre de la tabla en la base de datos
}

type UpdateProgramadReq struct {
	ID         string   `json:"-"`
	ProgramaId *int64   `json:"programa_id"`
	Desde      *int     `json:"desde"`
	Hasta      *int     `json:"hasta"`
	Valor      *float32 `json:"valor"`
	Liberado   *int     `json:"liberado"`
}

func (UpdateProgramadReq) TableName() string {
	return "programad" // Nombre de la tabla en la base de datos
}
