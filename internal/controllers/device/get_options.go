package devicecontrollers

import (
	"context"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	deviceservices "github.com/danielRamosMencia/consunet-api/internal/services/device"
	"github.com/gofiber/fiber/v2"
)

func GetOptions(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	devicesData, message, err :=
		deviceservices.SelectOptions(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  message,
			"código": "devs-err-000",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"mensaje": message,
		"data":    devicesData,
	})
}
