package models

type Turma struct {
	ID          uint      `gorm:"primaryKey"`
	Nome        string    `gorm:"size:100"`
	Semestre    string
	Ano         int
	Professor   Professor `gorm:"foreignKey:ProfessorID"`
	ProfessorID uint
	Alunos      []Aluno   `gorm:"many2many:aluno_turmas;"`
	Atividades  []Atividade
}