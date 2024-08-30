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
        c.JSON(http.StatusNotFound, gin.H{"error": "Professor not found"})
        return
    }

    // Excluir as turmas associadas ao professor
    if err := config.DB.Where("professor_id = ?", id).Delete(&models.Turma{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete associated turmas"})
        return
    }

    // Excluir o professor
    if err := config.DB.Delete(&professor).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Professor deleted successfully"})
}