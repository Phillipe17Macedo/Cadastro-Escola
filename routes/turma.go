package routes

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "./config"
  "./models"
)

func GetTurmas(c *gin.Context) {
  var turmas []models.Turma
  config.DB.Preload("Professor").Find(&turmas)
  c.JSON(http.StatusOK, turmas)
}

func CreateTurma(c *gin.Context) {
  var turma models.Turma
  if err := c.ShouldBindJSON(&turma); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  config.DB.Create(&turma)
  c.JSON(http.StatusOK, turma)
}