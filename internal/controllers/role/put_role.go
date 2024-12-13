package rolecontrollers

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
	roleservices "github.com/danielRamosMencia/consunet-api/internal/services/role"
	"github.com/danielRamosMencia/consunet-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func PutRole(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	var updateRole requests.UpdateRole
	id := c.Params("id")

	err := c.BodyParser(&updateRole)
	if err != nil {
		log.Println("Error parsing body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Campos para solicitud de actualizar rol incorrectos",
			"código": "role-err-003",
		})
	}

	err = validations.Validate.Struct(updateRole)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":        "Error en validación/es",
			"código":       "role-err-003",
			"validaciones": validations.ValidatorErrorsMap(err),
		})
	}

	message, err := roleservices.Update(ctx, id, updateRole)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  message,
			"código": "role-err-003",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"mensaje": message,
	})
}
