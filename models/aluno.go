type Aluno strut {
	ID uint `gorm:"primaryKey"`
	Nome string
	Matricula string
	CadTurma string
}