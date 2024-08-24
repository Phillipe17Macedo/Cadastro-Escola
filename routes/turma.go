package routes

import (
  "net/http"
  "github.com/gin-gonic/gin"
	"github.com/phillipe17macedo/Cadastro-Escola/config"
	"github.com/phillipe17macedo/Cadastro-Escola/models"
)

func GetTurmas(c *gin.Context) {
  var turmas map[string]models.Turma
  ref := config.FirebaseDB.NewRef("turmas")
  if err := ref.Get(context.Background(), &turmas); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve turmas"})
      return
  }

  c.JSON(http.StatusOK, gin.H{"data": turmas})
}

func CreateTurma(c *gin.Context) {
  var turma models.Turma
  if err := c.ShouldBindJSON(&turma); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
  }

  ref := config.FirebaseDB.NewRef("turmas")
  newRef, err := ref.Push(context.Background(), turma)
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create turma"})
      return
  }

  c.JSON(http.StatusOK, gin.H{"data": newRef.Key})
}