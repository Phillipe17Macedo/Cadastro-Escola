package models

type Atividade struct {
	ID uint `gorm:"primaryKey"`
	Turma Turma `gorm:"foreignKey:TurmaID"`
	TurmaID uint
	Valor float32
	Data string
}			