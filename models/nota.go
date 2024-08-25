package models

type Nota struct {
	ID          uint      `gorm:"primaryKey"`
	Valor       float32   // Valor da nota
	AlunoID     uint      // Chave estrangeira para o aluno
	Aluno       Aluno     `gorm:"foreignKey:AlunoID"`
	AtividadeID uint      // Chave estrangeira para a atividade
	Atividade   Atividade `gorm:"foreignKey:AtividadeID"`
}