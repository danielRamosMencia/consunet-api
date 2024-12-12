package rolecontrollers

import (
	"context"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	roleservices "github.com/danielRamosMencia/consunet-api/internal/services/role"
	"github.com/gofiber/fiber/v2"
)

func GetRoles(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	rolesData, message, err := roleservices.SelectAll(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  message,
			"código": "role-err-000",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"mensaje": message,
		"data":    rolesData,
	})
}
