package main

import (
	"log"
	"os"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/helpers"
	"github.com/danielRamosMencia/consunet-api/internal/middlewares"
	"github.com/danielRamosMencia/consunet-api/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file", err)
	}

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		ServerHeader:  "Fiber",
		AppName:       "CONSUNET-API v0.0.0",
		ErrorHandler:  middlewares.ErrorHandler,
	})

	database.ConnectDatabase()
	helpers.InitValidator()
	routes.SetUpRouter(app)

	port := os.Getenv("SERVER_PORT")
	app.Listen(port)
}
