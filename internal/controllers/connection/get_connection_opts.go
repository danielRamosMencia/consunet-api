package connectioncontrollers

import (
	"context"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	connectionservices "github.com/danielRamosMencia/consunet-api/internal/services/connection"
	"github.com/gofiber/fiber/v2"
)

func GetConnectionOptions(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	connectionsData, message, err := connectionservices.SelectOptions(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  message,
			"código": "conn-err-000",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"mensaje": message,
		"data":    connectionsData,
	})
}
