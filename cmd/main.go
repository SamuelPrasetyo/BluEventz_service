package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/SamuelPrasetyo/BluEventz_service/config"
	"github.com/SamuelPrasetyo/BluEventz_service/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	app := fiber.New()
	routes.MainRoutes(app)

	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
