package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/phillipe17macedo/Cadastro-Escola/config"
    "github.com/phillipe17macedo/Cadastro-Escola/routes"
    "time"
)

func main() {
    config.Connect()
    r := gin.Default()

    // Configurando CORS para permitir as origens específicas do frontend
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"https://66d242496274265ddcaf05af--cadastro-escola-frontend.netlify.app", "https://cadastro-escola-production.up.railway.app"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    // Middleware para lidar com requisições OPTIONS
    r.OPTIONS("/*cors", func(c *gin.Context) {
        c.AbortWithStatus(204)
    })

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

    r.Run() // Remove a especificação da porta para que o Railway aloque automaticamente
}