package usercontrollers

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	"github.com/danielRamosMencia/consunet-api/internal/helpers"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
	userservices "github.com/danielRamosMencia/consunet-api/internal/services/user"
	"github.com/danielRamosMencia/consunet-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func PatchUserSubscription(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	var addSubscription requests.UpdateUserSubscription

	err := c.BodyParser(&addSubscription)
	if err != nil {
		log.Println("Error parsing body: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Campos para solicitud de agregar subscripción incorrectos",
			"código": "user-err-001",
		})
	}

	err = validations.Validate.Struct(addSubscription)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":        "Error en validación/es",
			"código":       "user-err-001",
			"validaciones": validations.ValidatorErrorsMap(err),
		})
	}

	userData, err := helpers.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "No se pudo obtener la información de la sesión",
			"código": "user-err-001",
		})
	}

	message, err := userservices.UpdateSubscription(ctx, userData.Id, addSubscription)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  message,
			"código": "user-err-001",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"mensaje": message,
	})
}
