package routes

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/phillipe17macedo/Cadastro-Escola/config"
	"github.com/phillipe17macedo/Cadastro-Escola/models"
)

func GetAtividades(c *gin.Context) {
	var atividades []models.Atividade
	config.DB.Preload("Turma").Find(&atividades)
	c.JSON(http.StatusOK, atividades)
}

func CreateAtividade(c *gin.Context) {
	var atividadeInput struct {
		Nome    string
		Valor   float32
		Data    string // Data como string no formato "YYYY-MM-DD"
		TurmaID uint
	}

	if err := c.ShouldBindJSON(&atividadeInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Converte a data de string para time.Time
	data, err := time.Parse("2006-01-02", atividadeInput.Data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data deve estar no formato YYYY-MM-DD"})
		return
	}

	// Verifica se a soma dos valores das atividades da turma não ultrapassa 100 pontos
	var somaValores float32
	err = config.DB.Model(&models.Atividade{}).Where("turma_id = ?", atividadeInput.TurmaID).Select("IFNULL(sum(valor), 0)").Scan(&somaValores).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if somaValores+atividadeInput.Valor > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A soma dos valores das atividades ultrapassa 100 pontos"})
		return
	}

	atividade := models.Atividade{
		Nome:    atividadeInput.Nome,
		Valor:   atividadeInput.Valor,
		Data:    data,
		TurmaID: atividadeInput.TurmaID,
	}

	if err := config.DB.Create(&atividade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, atividade)
}