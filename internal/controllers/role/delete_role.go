package rolecontrollers

import (
	"context"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	roleservices "github.com/danielRamosMencia/consunet-api/internal/services/role"
	"github.com/gofiber/fiber/v2"
)

func DeleteRole(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	id := c.Params("id")

	message, err := roleservices.Delete(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  message,
			"código": "role-err-004",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"mensaje": message,
	})
}
