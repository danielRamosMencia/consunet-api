package routes

import (
	permissioncontrollers "github.com/danielRamosMencia/consunet-api/internal/controllers/permission"
	"github.com/gofiber/fiber/v2"
)

func PermissionsRoutes(router fiber.Router) {
	r := router.Group("/permissions")

	r.Get("/", permissioncontrollers.GetOptions)
}
