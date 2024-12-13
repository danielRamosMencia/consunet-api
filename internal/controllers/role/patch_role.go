package rolecontrollers

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	"github.com/danielRamosMencia/consunet-api/internal/helpers"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
	roleservices "github.com/danielRamosMencia/consunet-api/internal/services/role"
	"github.com/gofiber/fiber/v2"
)

func PatchRole(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	id := c.Params("id")

	changeActive, res, err := helpers.CheckActive(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  res,
			"código": "role-err-005",
		})
	}

	log.Println("changeActive === ", changeActive)

	toggleValue := requests.ToggleActive{
		Active: changeActive,
	}

	message, err := roleservices.UpdateActive(ctx, id, toggleValue)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  message,
			"código": "role-err-005",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"mensaje": message,
	})
}
