package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/phillipe17macedo/Cadastro-Escola/config"
	"github.com/phillipe17macedo/Cadastro-Escola/models"
)

func GetAlunos(c *gin.Context) {
	var alunos []models.Aluno
	config.DB.Preload("Turmas").Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func CreateAluno(c *gin.Context) {
    var alunoInput struct {
        Nome       string
        Matricula  string
        Turmas     []uint  // Espera uma lista de IDs de turmas
    }

    if err := c.ShouldBindJSON(&alunoInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var turmas []models.Turma
    config.DB.Where("id IN ?", alunoInput.Turmas).Find(&turmas)

    aluno := models.Aluno{
        Nome:      alunoInput.Nome,
        Matricula: alunoInput.Matricula,
        Turmas:    turmas,
    }

    if err := config.DB.Create(&aluno).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, aluno)
}