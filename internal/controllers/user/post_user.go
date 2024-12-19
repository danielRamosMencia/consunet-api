package usercontrollers

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
	userservices "github.com/danielRamosMencia/consunet-api/internal/services/user"
	"github.com/danielRamosMencia/consunet-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func PostUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	var newUser requests.CreateUser

	err := c.BodyParser(&newUser)
	if err != nil {
		log.Println("Error parsing body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Campos para solicitud de nuevo usuario incorrectos",
			"código": "user-err-000",
		})
	}

	err = validations.Validate.Struct(newUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":        "Error en validación/es",
			"código":       "user-err-000",
			"validaciones": validations.ValidatorErrorsMap(err),
		})
	}

	message, err := userservices.Insert(ctx, newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  message,
			"código": "user-err-000",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"mensaje": message,
	})
}
