package routes

import (
	"net/http"
    "strconv"
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

func UpdateTurma(c *gin.Context) {
    var turma models.Turma
    id := c.Param("id")

    turmaID, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Turma ID"})
        return
    }

    if err := config.DB.First(&turma, "id = ?", turmaID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Turma not found"})
        return
    }

    var turmaInput struct {
        Nome        *string  // Use ponteiro para permitir valor nulo
        Semestre    *string  // Use ponteiro para permitir valor nulo
        Ano         *int     // Use ponteiro para permitir valor nulo
        ProfessorID *uint    // Use ponteiro para permitir valor nulo
    }

    if err := c.ShouldBindJSON(&turmaInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Atualiza apenas os campos fornecidos
    if turmaInput.Nome != nil {
        turma.Nome = *turmaInput.Nome
    }
    if turmaInput.Semestre != nil {
        turma.Semestre = *turmaInput.Semestre
    } else if turma.Semestre == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Semestre is required and cannot be empty"})
        return
    }
    if turmaInput.Ano != nil {
        turma.Ano = *turmaInput.Ano
    } else if turma.Ano == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Ano is required and cannot be zero"})
        return
    }
    if turmaInput.ProfessorID != nil {
        // Verifica se o ProfessorID existe
        var professor models.Professor
        if err := config.DB.First(&professor, turmaInput.ProfessorID).Error; err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Professor ID"})
            return
        }
        turma.ProfessorID = *turmaInput.ProfessorID
    } else if turma.ProfessorID == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ProfessorID is required and cannot be zero"})
        return
    }

    if err := config.DB.Save(&turma).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, turma)
}

func DeleteTurma(c *gin.Context) {
    var turma models.Turma
    id := c.Param("id")

    if err := config.DB.First(&turma, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Turma not found"})
        return
    }

    if err := config.DB.Delete(&turma).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Turma deleted successfully"})
}