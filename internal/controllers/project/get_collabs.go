package projectcontrollers

import (
	"context"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	projectservices "github.com/danielRamosMencia/consunet-api/internal/services/project"
	"github.com/gofiber/fiber/v2"
)

func GetCollabs(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	projectId := c.Params("project_id")

	collabsData, message, err := projectservices.SelectCollabs(ctx, projectId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  message,
			"código": "pro-err-004",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    collabsData,
		"mensaje": message,
	})
}
