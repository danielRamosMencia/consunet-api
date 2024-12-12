package rolecontrollers

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	"github.com/danielRamosMencia/consunet-api/internal/helpers"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
	roleservices "github.com/danielRamosMencia/consunet-api/internal/services/role"
	"github.com/gofiber/fiber/v2"
)

func CreateRole(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	var newRole requests.CreateRole

	err := c.BodyParser(&newRole)
	if err != nil {
		log.Println("Error parsing body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Campos para solicitud de nuevo rol incorrectos",
			"código": "role-err-002",
		})
	}

	err = helpers.Validate.Struct(newRole)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":        "Error en validación/es",
			"código":       "role-err-002",
			"validaciones": helpers.ValidatorErrorsMap(err),
		})
	}

	message, err := roleservices.Insert(ctx, newRole)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  message,
			"código": "role-err-002",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"mensaje": message,
	})
}
