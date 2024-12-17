package permissioncontrollers

import (
	"context"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	permissionservices "github.com/danielRamosMencia/consunet-api/internal/services/permission"
	"github.com/gofiber/fiber/v2"
)

func GetOptions(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	permissionsData, message, err :=
		permissionservices.SelectOptions(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  message,
			"código": "per-err-000",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"mensaje": message,
		"data":    permissionsData,
	})
}
