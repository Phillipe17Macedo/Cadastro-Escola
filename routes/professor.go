package routes

import (
	"github.com/phillipe17macedo/Cadastro-Escola/config"
	"github.com/phillipe17macedo/Cadastro-Escola/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfessores(c *gin.Context) {
	var professores []models.Professor
	config.DB.Find(&professores)
	c.JSON(http.StatusOK, professores)
}

func CreateProfessor(c *gin.Context) {
    var professor models.Professor
    if err := c.ShouldBindJSON(&professor); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := config.DB.Create(&professor).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, professor)
}

func UpdateProfessor(c *gin.Context) {
    var professor models.Professor
    id := c.Param("id")

    // Corrigindo a consulta para usar o ID corretamente
    if err := config.DB.Where("id = ?", id).First(&professor).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Professor not found"})
        return
    }

    if err := c.ShouldBindJSON(&professor); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Save(&professor).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, professor)
}

func DeleteProfessor(c *gin.Context) {
    var professor models.Professor
    id := c.Param("id")

    // Verificar se o professor existe
    if err := config.DB.First(&professor, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Professor não encontrado"})
        return
    }

    // Verificar se existem turmas associadas ao professor
    var turmas []models.Turma
    if err := config.DB.Where("professor_id = ?", id).Find(&turmas).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar turmas associadas"})
        return
    }

    if len(turmas) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Não é possível remover o professor porque existem turmas associadas"})
        return
    }

    // Excluir o professor, pois não há turmas associadas
    if err := config.DB.Delete(&professor).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir o professor"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Professor excluído com sucesso"})
}