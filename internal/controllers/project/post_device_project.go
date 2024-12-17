package projectcontrollers

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	"github.com/danielRamosMencia/consunet-api/internal/helpers"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
	projectservices "github.com/danielRamosMencia/consunet-api/internal/services/project"
	"github.com/danielRamosMencia/consunet-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func PostDeviceProject(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	var addDevice requests.CreateDeviceProject

	err := c.BodyParser(&addDevice)
	if err != nil {
		log.Println("Error parsing body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Campos para solicitud de nuevo proyecto incorrectos",
			"código": "pro-err-002",
		})
	}

	err = validations.Validate.Struct(addDevice)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":        "Error en validación/es",
			"código":       "pro-err-002",
			"validaciones": validations.ValidatorErrorsMap(err),
		})
	}

	userData, err := helpers.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "No se pudo obtener la información de la sesión",
			"código": "pro-err-002",
		})
	}

	message, err := projectservices.InsertDeviceProject(ctx, userData.Email, addDevice)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  message,
			"código": "pro-err-002",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"mensaje": message,
	})
}
