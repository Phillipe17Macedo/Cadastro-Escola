package config

import (
	"log"

	"github.com/phillipe17macedo/Cadastro-Escola/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// String de conexão com o MySQL
	dsn := "root:teodoro10@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate para sincronizar os modelos com o banco de dados
	DB.AutoMigrate(&models.Professor{}, &models.Turma{}, &models.Aluno{}, &models.Atividade{}, &models.Nota{})
}
