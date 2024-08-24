package models

import "time"

type Atividade struct {
	ID       uint      `gorm:"primaryKey"`
	Turma    Turma     `gorm:"foreignKey:TurmaID"`
	TurmaID  uint
	Valor    float32
	Data     time.Time `gorm:"type:date"` 
}