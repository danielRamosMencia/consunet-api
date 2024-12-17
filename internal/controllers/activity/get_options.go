package activitycontrollers

import (
	"context"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	activityservices "github.com/danielRamosMencia/consunet-api/internal/services/activity"
	"github.com/gofiber/fiber/v2"
)

func GetOptions(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	activitiesData, message, err :=
		activityservices.SelectOptions(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  message,
			"código": "conn-err-000",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"mensaje": message,
		"data":    activitiesData,
	})
}
