package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phillipe17macedo/Cadastro-Escola/config"
	"github.com/phillipe17macedo/Cadastro-Escola/models"
)

func GetNotas(c *gin.Context) {
	var notas []models.Nota
	config.DB.Preload("Aluno").Preload("Atividade").Find(&notas)
	c.JSON(http.StatusOK, notas)
}

func CreateNota(c *gin.Context) {
	var notaInput struct {
		Valor       float32
		AlunoID     uint
		AtividadeID uint
	}

	if err := c.ShouldBindJSON(&notaInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buscar a atividade para validar o valor máximo da nota
	var atividade models.Atividade
	if err := config.DB.First(&atividade, notaInput.AtividadeID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Atividade não encontrada"})
		return
	}

	// Verificar se a nota excede o valor máximo permitido pela atividade
	if notaInput.Valor > atividade.Valor {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A nota excede o valor máximo permitido pela atividade"})
		return
	}

	nota := models.Nota{
		Valor:       notaInput.Valor,
		AlunoID:     notaInput.AlunoID,
		AtividadeID: notaInput.AtividadeID,
	}

	if err := config.DB.Create(&nota).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nota)
}