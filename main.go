package main

import (
	"log"
	"os"

	"go-swag/config"
	"go-swag/model"
	"go-swag/route"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)
// @title User API documentation
// @version 1.0.0

// @host localhost:5000


func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.DB = config.SetupDatabase()
	config.DB.AutoMigrate(&model.User{})

	r := route.SetupRouter()

	r.Run(":" + os.Getenv("SERVER_PORT"))
}
