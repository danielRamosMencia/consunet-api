package helpers

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CheckActive(c *fiber.Ctx) (bool, string, error) {
	var reqBody map[string]interface{}

	err := c.BodyParser(&reqBody)
	if err != nil {
		log.Println("Error parsing body: ", err)
		return false, "Solicitud para cambiar campo activo incorrecta", err
	}

	hasActive, ok := reqBody["active"]
	if !ok {
		log.Println("Request body missing active field")
		return false, "No se encontró el campo 'active'", errors.New("request body missing active field")
	}

	active, ok := hasActive.(bool)
	if !ok {
		log.Println("Bad type for active", err)
		return false, "El campo activo no es un valor booleano", errors.ErrUnsupported
	}

	return active, "", nil
}
