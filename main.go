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

	r.GET("/professores", routes.GetProfessores)
	r.POST("/professores", routes.CreateProfessor)
	r.GET("/turmas", routes.GetTurmas)
	r.POST("/turmas", routes.CreateTurma)

	r.Run(":8080")
}
