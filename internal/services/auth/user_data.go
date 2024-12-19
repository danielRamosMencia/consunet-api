package authservices

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
	"github.com/danielRamosMencia/consunet-api/internal/models/responses"
	"golang.org/x/crypto/bcrypt"
)

func UserData(ctx context.Context, req requests.Auth) (responses.UserData, string, error) {
	var userData responses.UserData
	var storedPassword string

	const query = `
	SELECT
		"id",
		"username",
		"email",
		"active",
		"password",
		"subscription_id"
	FROM
		"User"
	WHERE
		"username" = $1 AND
		"email" = $2;
	`

	row := database.Connx.QueryRowContext(
		ctx,
		query,
		req.Username,
		req.Email,
	)
	err := row.Scan(
		&userData.Id,
		&userData.Username,
		&userData.Email,
		&userData.Active,
		&storedPassword,
		&userData.Subscription_id,
	)

	switch {
	case err == sql.ErrNoRows:
		log.Println("Bad credentials")
		return userData, "Credenciales incorrectas", err
	case err != nil:
		log.Print("Error getting user data: ", err)
		return userData, "Error al obtener los datos de usuario", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(req.Password))
	if err != nil {
		log.Println("Error comparing stored password with request password: ", err)
		return userData, "Credenciales incorrectas", err
	}

	if !userData.Active {
		log.Print("Inactive user")
		return userData, "Usuario inactivo", errors.New("inactive user")
	}

	return userData, "Usuario encontrado", nil
}
