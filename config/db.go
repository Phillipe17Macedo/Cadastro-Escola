package main

import (
  "gorm.io/gorm"
	"gorm.io/driver/postgres"
)
  
  dsn := "user=postgres password=mysecretpassword dbname=mydb port=5432 sslmode=disable"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
	panic("failed to connect database")
  }