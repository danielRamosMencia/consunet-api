package projectcontrollers

import (
	"context"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	projectservices "github.com/danielRamosMencia/consunet-api/internal/services/project"
	"github.com/gofiber/fiber/v2"
)

func GetProjectDevices(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	projectId := c.Params("project_id")

	projectDevicesData, message, err := projectservices.SelectProjectDevices(ctx, projectId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": message,
			"code":  "pro-err-005",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    projectDevicesData,
		"message": message,
	})
}
