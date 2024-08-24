package models

type Professor struct {
	ID        uint    `gorm:"primaryKey"`
	Nome      string  `gorm:"size:100"`
	Email     string  `gorm:"unique"`
	CPF       string  `gorm:"unique"`
	Turmas    []Turma `gorm:"foreignKey:ProfessorID"`
}