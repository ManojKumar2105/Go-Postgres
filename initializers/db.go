package initializers

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(config *Config) {
	// db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost@5432/postgres"), &gorm.Config{})

	// dsn := "host=localhost user=postgres password=password123 dbname=golang-gorm port=6500 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db

	if err != nil {
		panic(err)
	}
	fmt.Println("? Connected Successfully to the Database")
}
