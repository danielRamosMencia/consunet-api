package routes

import (
	usercontrollers "github.com/danielRamosMencia/consunet-api/internal/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	r := router.Group("/users")

	r.Post("/", usercontrollers.PostUser)
}
