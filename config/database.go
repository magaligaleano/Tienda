package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(127.0.0.1:3306)/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	return db
}
