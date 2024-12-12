package routes

import (
	rolecontrollers "github.com/danielRamosMencia/consunet-api/internal/controllers/role"
	"github.com/gofiber/fiber/v2"
)

func RoleRoutes(router fiber.Router) {
	r := router.Group("/roles")

	r.Get("/", rolecontrollers.GetRoles)
	r.Get("/:id", rolecontrollers.GetRole)
}
