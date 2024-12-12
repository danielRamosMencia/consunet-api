package rolecontrollers

import (
	"context"
	"database/sql"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	roleservices "github.com/danielRamosMencia/consunet-api/internal/services/role"
	"github.com/gofiber/fiber/v2"
)

func GetRole(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	id := c.Params("id")

	roleData, message, err := roleservices.SelectOne(ctx, id)
	switch {
	case err == sql.ErrNoRows:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  message,
			"código": "role-err-001",
		})
	case err != nil:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  message,
			"código": "role-err-001",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"mensaje": message,
		"data":    roleData,
	})
}
