package models

import (
	"time"
)

type Atividade struct {
	ID      uint      `gorm:"primaryKey"`
	Nome    string    `gorm:"size:100"`
	Valor   float32   // Valor da atividade (pontuação)
	Data    time.Time // Data da atividade
	TurmaID uint      // Chave estrangeira para a turma
	Turma   Turma     `gorm:"foreignKey:TurmaID"`
}