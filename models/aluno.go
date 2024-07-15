package models

type ALuno struct {
	ID uint `gorm:"primaryKey"`
	Nome string `gorm:"size:100"`
	Matricula string `gorm:"unique"`
	Turmas []Turma `gorm:"many2many:aluno_turmas;"`
}