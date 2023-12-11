package main

import (
	"github.com/Manojkumar2105/gin-gorm-setup/initializers"
	"github.com/Manojkumar2105/gin-gorm-setup/routes"
	"github.com/gin-gonic/gin"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

func initialize() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	initializers.Connect(&config)
}

func main() {
	router := gin.New()

	initialize()

	routes.Helloroute(router)

	router.Run(":4000")
}

// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 TimeZone=Asia/Kolkata"
// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// if err != nil {
// 	panic("failed to connect database")
// }

// fmt.Println(db)
