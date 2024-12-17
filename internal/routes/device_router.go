package routes

import (
	devicecontrollers "github.com/danielRamosMencia/consunet-api/internal/controllers/device"
	"github.com/gofiber/fiber/v2"
)

func DeviceRoutes(router fiber.Router) {
	r := router.Group("/devices")

	r.Get("/", devicecontrollers.GetOptions)
}
