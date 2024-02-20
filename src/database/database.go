package database

import (
	"os"
	models "poj/src/database/models"
	"poj/src/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	// valid envs
	err := utils.ValidateEnv([]string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT"})
	if err != nil {
		panic(err)
	}

	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	return db
}

func Init() *gorm.DB {
	db := Connect()
	db.AutoMigrate(&models.User{})
	return db
}
