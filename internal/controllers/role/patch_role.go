package rolecontrollers

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	roleservices "github.com/danielRamosMencia/consunet-api/internal/services/role"
	"github.com/danielRamosMencia/consunet-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func PatchRole(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	id := c.Params("id")

	changeActive, res, err := validations.CheckActive(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  res,
			"código": "role-err-005",
		})
	}

	log.Println("changeActive === ", changeActive)

	message, err := roleservices.UpdateActive(ctx, id, changeActive)
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
