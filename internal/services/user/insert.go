package userservices

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
	"github.com/danielRamosMencia/consunet-api/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func Insert(ctx context.Context, req requests.CreateUser) (string, error) {
	// Get the default role
	var roleId string

	const roleQuery = `
	SELECT 
		"id"
	FROM
		"Role"
	WHERE
		"code" = 'CLI';
	`

	row := database.Connx.QueryRowContext(ctx, roleQuery)
	err := row.Scan(
		&roleId,
	)
	if err != nil {
		log.Println("Error getting default role", err)
		return "Error al registrar el usuario", err
	}

	// Create the user
	const query = `
	INSERT INTO "User"
		("id", "username", "email", "password", "role_id", "subscription_id")
	VALUES
		($1, $2, $3, $4, $5, $6);
	`

	id := utils.GenerateId()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password: ", err)
		return "Error al registrar el usuario", err
	}

	_, err = database.Connx.ExecContext(
		ctx,
		query,
		id,
		req.Username,
		req.Email,
		hashedPassword,
		roleId,
		req.Subscription_id,
	)
	if err != nil {
		log.Println("Error creating user: ", err)
		return "Error al registrar el usuario", err
	}

	return "El usuario se ha registrado con éxito", nil
}
