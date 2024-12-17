package routes

import (
	connectioncontrollers "github.com/danielRamosMencia/consunet-api/internal/controllers/connection"
	"github.com/danielRamosMencia/consunet-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func ConnectionRoutes(router fiber.Router) {
	r := router.Group("/connections")

	r.Get("/", middlewares.AuthRequired, connectioncontrollers.GetOptions)
}

func ConnectionRoutesAdm(router fiber.Router) {
	// TODO
}
