package roleservices

import (
	"context"
	"database/sql"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/responses"
)

func SelectOne(ctx context.Context, id string) (responses.Role, string, error) {
	var role responses.Role

	const query = `
	SELECT 
		"id",
		"name",
		"code",
		"active",
		"created_at",
		"updated_at"
	FROM 
		"Role"
	WHERE 
		"id" = $1;
	`

	row := database.Connx.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&role.Id,
		&role.Name,
		&role.Code,
		&role.Active,
		&role.Created_at,
		&role.Updated_at,
	)

	switch {
	case err == sql.ErrNoRows:
		log.Println("Role with id: " + id + " not found")
		return role, "Rol no encontrado", err
	case err != nil:
		log.Println("Error getting role: ", err)
		return role, "Error al obtener el rol", err
	}

	return role, "Rol encontrado", nil
}
