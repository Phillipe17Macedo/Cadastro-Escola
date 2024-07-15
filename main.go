package main

import (
  "github.com/gin-gonic/gin"
  "./config"
  "./routes"
)

func main() {
  config.Connect()
  r := gin.Default()

  r.GET("/professores", routes.GetProfessores)
  r.POST("/professores", routes.CreateProfessor)
  r.GET("/turmas", routes.GetTurmas)
  r.POST("/turmas", routes.CreateTurma)

  r.Run(":8080")
}