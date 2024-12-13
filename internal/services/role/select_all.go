package roleservices

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/responses"
)

func SelectAll(ctx context.Context) ([]responses.Role, string, error) {
	var roles []responses.Role

	const sql = `
	SELECT
		"id",	 
		"name",
		"code",
		"active",
		"created_at",
		"updated_at"
	FROM 
		"Role";
	`

	rows, err := database.Connx.QueryContext(ctx, sql)
	if err != nil {
		log.Println("Error getting roles: ", err)
		return nil, "Error al obtener los roles", err
	}
	defer rows.Close()

	for rows.Next() {
		var role responses.Role
		err := rows.Scan(
			&role.Id,
			&role.Name,
			&role.Code,
			&role.Active,
			&role.Created_at,
			&role.Updated_at,
		)
		if err != nil {
			log.Println("Error scanning role:", err)
			return nil, "Error escaneando un rol", err
		}
		roles = append(roles, role)
	}

	return roles, "Lista de roles", nil
}
