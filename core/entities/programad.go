package entities

type Programad struct {
	ID         string  `json:"id" gorm:"primaryKey;autoIncrement"`
	ProgramaId int64   `json:"programa_id"`
	Desde      int     `json:"desde"`
	Hasta      int     `json:"hasta"`
	Valor      float32 `json:"valor"`
	Liberado   int     `json:"liberado"`
}
