package config

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "log"
    "github.com/phillipe17macedo/Cadastro-Escola/models"
)

var DB *gorm.DB

func Connect() {
    // String de conexão com o MySQL
    dsn := "admin:1234@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
    
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // AutoMigrate para sincronizar os modelos com o banco de dados
    DB.AutoMigrate(&models.Professor{}, &models.Turma{}, &models.Aluno{}, &models.Atividade{})
}
