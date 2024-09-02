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

    // Verificar se já existe uma nota para o mesmo aluno e atividade
    var existingNota models.Nota
    if err := config.DB.Where("aluno_id = ? AND atividade_id = ?", notaInput.AlunoID, notaInput.AtividadeID).First(&existingNota).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Já existe uma nota para este aluno nesta atividade"})
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

func UpdateNota(c *gin.Context) {
    var nota models.Nota
    id := c.Param("id")

    // Tenta encontrar a nota pelo ID
    if err := config.DB.First(&nota, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Nota não encontrada"})
        return
    }

    var notaInput struct {
        Valor float32 `json:"valor"`
    }

    if err := c.ShouldBindJSON(&notaInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Atualiza apenas o campo 'valor' da nota
    nota.Valor = notaInput.Valor

    // Salva as alterações
    if err := config.DB.Save(&nota).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, nota)
}

func DeleteNota(c *gin.Context) {
	var nota models.Nota
	id := c.Param("id")

	if err := config.DB.First(&nota, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nota não encontrada"})
		return
	}

	if err := config.DB.Delete(&nota).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Nota deletada com sucesso"})
}