package routes

import (
	activitycontrollers "github.com/danielRamosMencia/consunet-api/internal/controllers/activity"
	"github.com/gofiber/fiber/v2"
)

func ActivityRoutes(router fiber.Router) {
	r := router.Group("/activities")

	r.Get("/", activitycontrollers.GetOptions)
}
