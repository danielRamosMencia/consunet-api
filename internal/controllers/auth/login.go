package authcontrollers

import (
	"context"
	"database/sql"

	"github.com/danielRamosMencia/consunet-api/internal/configs"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
	authservices "github.com/danielRamosMencia/consunet-api/internal/services/auth"
	"github.com/danielRamosMencia/consunet-api/internal/utils"
	"github.com/danielRamosMencia/consunet-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeOut)
	defer cancel()

	var authRequest requests.Auth

	err := c.BodyParser(&authRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Campos para solicitud de inicio de sesión incorrectos",
			"code":    "error-log-000",
		})
	}

	err = validations.Validate.Struct(authRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error en validaciones",
			"code":    "error-log-000",
		})
	}

	userData, message, err := authservices.UserData(ctx, authRequest)
	switch {
	case err == sql.ErrNoRows:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": message,
			"code":    "error-log-000",
		})
	case err != nil:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": message,
			"code":    "error-log-000",
		})
	}

	token, maxAge, err := utils.GenerateJwt(userData)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Inautorizado",
			"code":    "error-log-000",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HTTPOnly: true,
		MaxAge:   int(maxAge),
		SameSite: "Lax", // "None" if the request is cross-origin
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
		"user":  userData,
	})
}
