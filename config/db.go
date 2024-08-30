package config

import (
	"log"
	"github.com/phillipe17macedo/Cadastro-Escola/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// URL completa fornecida pelo Railway
	dsn := "root:ehrByRQgwxDRoiaAvYcuBCCxkTxwvKem@tcp(autorack.proxy.rlwy.net:26077)/railway?charset=utf8mb4&parseTime=True&loc=Local"

	// Tenta abrir a conexão com o banco de dados
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate para sincronizar os modelos com o banco de dados (opcional)
	DB.AutoMigrate(&models.Professor{}, &models.Turma{}, &models.Aluno{}, &models.Atividade{}, &models.Nota{})
}