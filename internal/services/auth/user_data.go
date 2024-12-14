package authservices

import (
	"context"
	"database/sql"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
	"github.com/danielRamosMencia/consunet-api/internal/models/responses"
)

func UserData(ctx context.Context, req requests.Auth) (responses.UserData, string, error) {
	var userData responses.UserData
	const query = `
	SELECT
		"id",
		"username",
		"email",
	FROM
		"User"
	WHERE
		"email" = $1 AND
		"password" = $2;
	`

	row := database.Connx.QueryRowContext(
		ctx,
		query,
		req.Email,
		req.Password,
	)
	err := row.Scan(
		&userData.Id,
		&userData.Username,
		&userData.Email,
	)

	switch {
	case err == sql.ErrNoRows:
		log.Println("Bad credentials")
		return userData, "Credenciales incorrectas", err
	case err != nil:
		log.Print("Error getting user data: ", err)
		return userData, "Error al obtener los datos de usuario", err
	}

	return userData, "Usuario encontrado", nil
}
