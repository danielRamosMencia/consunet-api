package projectcontrollers

import (
	"context"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	projectservices "github.com/danielRamosMencia/consunet-api/internal/services/project"
	"github.com/gofiber/fiber/v2"
)

func DeleteDeviceProject(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	id := c.Params("id")

	message, err := projectservices.DeleteDeviceProject(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  message,
			"código": "pro-err-006",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"mensaje": message,
	})
}
