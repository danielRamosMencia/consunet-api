package projectservices

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/responses"
)

func SelectCollabs(ctx context.Context, projectId string) ([]responses.Collabs, string, error) {
	var collabs []responses.Collabs
	const query = `
	SELECT 
		"UP"."id",
		"U"."id" AS "collab_id",
		"U"."username" AS "collab",
		"U"."email",
		"P"."name" AS "permission"
	FROM 
		"UsersOnProject" AS "UP"
	JOIN "User" AS "U" ON "UP"."user_id" = "U"."id"
	JOIN "Permission" AS "P" ON "UP"."permission_id" = "P"."id"
	WHERE 
		"project_id" = $1;
	`

	rows, err := database.Connx.QueryContext(ctx, query, projectId)
	if err != nil {
		log.Println("Error getting project collabs", err)
		return nil, "Error al obtener los colaboradores", err
	}
	defer rows.Close()

	for rows.Next() {
		var collab responses.Collabs

		err := rows.Scan(
			&collab.Id,
			&collab.Collab_id,
			&collab.Collab,
			&collab.Email,
			&collab.Permission,
		)
		if err != nil {
			log.Println("Error scanning collab: ", err)
			return nil, "Error escaneando un colaborador", err
		}

		collabs = append(collabs, collab)
	}

	if len(collabs) == 0 {
		return collabs, "No se encontraron colaboradores", nil
	}

	return collabs, "Lista de colaboradores", nil
}
