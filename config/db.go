package config

import (
  "gorm.io/gorm"
	"gorm.io/driver/postgres"
  "log"
)
  
var DB *gorm.DB

func Connect() {
  dsn := "user=postgres password=mysecretpassword dbname=mydb port=5432 sslmode=disable"
  var err error
  DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatal("Failed to connect to database:", err)
  }
}