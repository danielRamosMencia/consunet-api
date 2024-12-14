package routes

import (
	authcontrollers "github.com/danielRamosMencia/consunet-api/internal/controllers/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router) {
	r := router.Group("/auth")

	r.Post("/sign-in", authcontrollers.Login)
	r.Post("/sign-out", authcontrollers.Logout)
}
