package helpers

import (
	"errors"

	"github.com/danielRamosMencia/consunet-api/internal/models/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetClaims(c *fiber.Ctx) (responses.UserData, error) {

	userClaims, ok := c.Locals("user_claims").(jwt.MapClaims)
	if !ok {
		return responses.UserData{}, errors.New("error getting jwt claims")
	}

	return responses.UserData{
		Id:       userClaims["id"].(string),
		Username: userClaims["username"].(string),
		Email:    userClaims["email"].(string),
		Active:   userClaims["active"].(bool),
	}, nil

}
