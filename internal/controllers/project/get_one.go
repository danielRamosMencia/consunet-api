package projectcontrollers

import (
	"context"
	"database/sql"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	projectservices "github.com/danielRamosMencia/consunet-api/internal/services/project"
	"github.com/gofiber/fiber/v2"
)

func GetProject(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	id := c.Params("id")

	projectData, message, err := projectservices.SelectOne(ctx, id)
	switch {
	case err == sql.ErrNoRows:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  message,
			"código": "pro-err-012",
		})
	case err != nil:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  message,
			"código": "pro-err-012",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"mensaje": message,
		"data":    projectData,
	})
}
