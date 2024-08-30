package routes

import (
	"net/http"
	"fmt"
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
		Nome      string
		Matricula string
		Turmas    []uint // Espera uma lista de IDs de turmas
	}

	if err := c.ShouldBindJSON(&alunoInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Turmas recebidas:", alunoInput.Turmas)

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

func UpdateAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Param("id")

	// Garantindo que o ID seja encontrado corretamente
	if err := config.DB.Preload("Turmas").First(&aluno, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno not found"})
		return
	}

	var alunoInput struct {
		Nome      string
		Matricula string
		Turmas    []uint
	}

	if err := c.ShouldBindJSON(&alunoInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var turmas []models.Turma
	if len(alunoInput.Turmas) > 0 {
		config.DB.Where("id IN ?", alunoInput.Turmas).Find(&turmas)
	}

	aluno.Nome = alunoInput.Nome
	aluno.Matricula = alunoInput.Matricula
	aluno.Turmas = turmas

	if err := config.DB.Save(&aluno).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeleteAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Param("id")

	// Inicia uma transação
	tx := config.DB.Begin()

	if err := tx.First(&aluno, "id = ?", id).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno not found"})
		return
	}

	// Excluir manualmente as associações na tabela aluno_turmas
	if err := tx.Exec("DELETE FROM aluno_turmas WHERE aluno_id = ?", id).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete associated records"})
		return
	}

	// Agora, exclua o aluno
	if err := tx.Delete(&aluno).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Confirma a transação
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aluno deleted successfully"})
}
