package projectcontrollers

import (
	"context"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	"github.com/danielRamosMencia/consunet-api/internal/helpers"
	projectservices "github.com/danielRamosMencia/consunet-api/internal/services/project"
	"github.com/gofiber/fiber/v2"
)

func GetUserProjects(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	userData, err := helpers.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "No se pudo obtener la información de la sesión",
			"código": "pro-err-000",
		})
	}

	userProjects, message, err := projectservices.SelectUserProjects(ctx, userData.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  message,
			"código": "pro-err-000",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    userProjects,
		"mensaje": message,
	})
}
