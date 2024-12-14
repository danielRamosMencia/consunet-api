package utils

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/danielRamosMencia/consunet-api/internal/models/responses"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(data responses.UserData) (string, int64, error) {
	secret := os.Getenv("JWT_SECRET")
	rawDuration := os.Getenv("JWT_TIME")

	duration, err := strconv.Atoi(rawDuration)
	if err != nil {
		log.Panic("Cannot convert string to int: ", err)
	}

	now := time.Now()
	exp := now.Add(time.Hour * time.Duration(duration)).Unix()
	iat := now.Unix()
	nbf := now.Unix()
	maxAge := exp - iat

	log.Print(maxAge, "maxAge")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = data.Id
	claims["username"] = data.Username
	claims["email"] = data.Email
	claims["exp"] = exp
	claims["iat"] = iat
	claims["nbf"] = nbf

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Panic("Error signin token: ", err)
		return "", maxAge, err
	}

	return signedToken, maxAge, err
}
