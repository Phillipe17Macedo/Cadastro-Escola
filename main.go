package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phillipe17macedo/Cadastro-Escola/config"
	"github.com/phillipe17macedo/Cadastro-Escola/routes"
)

func main() {
	config.Connect()
	r := gin.Default()

	r.Use(cors.Default())

	// Métodos Rota Professor
	r.GET("/professores", routes.GetProfessores)
	r.POST("/professores", routes.CreateProfessor)
	r.PUT("/professores/:id", routes.UpdateProfessor)
	r.DELETE("/professores/:id", routes.DeleteProfessor)

	// Métodos Rota Turma
	r.GET("/turmas", routes.GetTurmas)
	r.POST("/turmas", routes.CreateTurma)
	r.PUT("/turmas/:id", routes.UpdateTurma)
	r.DELETE("/turmas/:id", routes.DeleteTurma)

	// Métodos Rota Aluno
	r.GET("/alunos", routes.GetAlunos)
	r.POST("/alunos", routes.CreateAluno)
	r.PUT("/alunos/:id", routes.UpdateAluno)
	r.DELETE("/alunos/:id", routes.DeleteAluno)

	// Métodos Rota Atividades
	r.GET("/atividades", routes.GetAtividades)
	r.POST("/atividades", routes.CreateAtividade)
	r.PUT("/atividades/:id", routes.UpdateAtividade)
	r.DELETE("/atividades/:id", routes.DeleteAtividade)

	// Métodos Rota Notas
	r.GET("/notas", routes.GetNotas)
	r.POST("/notas", routes.CreateNota)
	r.PUT("/notas/:id", routes.UpdateNota)
	r.DELETE("/notas/:id", routes.DeleteNota)

	r.Run(":8080")
}
