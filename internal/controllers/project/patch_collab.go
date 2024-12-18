package projectcontrollers

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
	projectservices "github.com/danielRamosMencia/consunet-api/internal/services/project"
	"github.com/danielRamosMencia/consunet-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func PatchCollab(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	var updatePermission requests.UpdateCollab
	id := c.Params("id")

	err := c.BodyParser(&updatePermission)
	if err != nil {
		log.Println("Error parsing body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Campos para solicitud de actualizar el colaborador incorrectos",
			"código": "pro-err-011",
		})
	}

	err = validations.Validate.Struct(updatePermission)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":        "Error en validación/es",
			"código":       "pro-err-011",
			"validaciones": validations.ValidatorErrorsMap(err),
		})
	}

	message, err := projectservices.UpdateCollab(ctx, id, updatePermission)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  message,
			"código": "pro-err-011",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"mensaje": message,
	})
}
