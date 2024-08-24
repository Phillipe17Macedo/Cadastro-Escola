package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/phillipe17macedo/Cadastro-Escola/config"
	"github.com/phillipe17macedo/Cadastro-Escola/models"
)

func GetTurmas(c *gin.Context) {
	var turmas []models.Turma
	config.DB.Preload("Professor").Find(&turmas) // Preload para carregar o relacionamento com Professor
	c.JSON(http.StatusOK, turmas)
}

func CreateTurma(c *gin.Context) {
    var turmaInput struct {
        Nome        string
        Semestre    string
        Ano         int  // Ajustar para int
        ProfessorID uint  // Ajustar para uint
    }

    // Bind JSON para a struct turmaInput
    if err := c.ShouldBindJSON(&turmaInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "details": err})
        return
    }

    // Cria a nova turma com os dados recebidos
    turma := models.Turma{
        Nome:        turmaInput.Nome,
        Semestre:    turmaInput.Semestre,
        Ano:         turmaInput.Ano,
        ProfessorID: turmaInput.ProfessorID,
    }

    // Salva a nova turma no banco de dados
    if err := config.DB.Create(&turma).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, turma)
}
