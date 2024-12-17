package projectservices

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/responses"
)

func SelectUserProjects(ctx context.Context, userId string) ([]responses.UserProjects, string, error) {
	var userProjects []responses.UserProjects
	const query = `
	SELECT 
		"id",
		"name",
		"code"
	FROM 
		"Project"
	WHERE 
		"user_id" = $1;
	`

	rows, err := database.Connx.QueryContext(ctx, query, userId)
	if err != nil {
		log.Println("Error getting user projects: ", err)
		return nil, "Error obteniendo los proyectos del usuario", err
	}
	defer rows.Close()

	for rows.Next() {
		var project responses.UserProjects
		err := rows.Scan(
			&project.Id,
			&project.Name,
			&project.Code,
		)
		if err != nil {
			log.Println("Error scanning user project", err)
			return nil, "Error escaneado los proyectos", err
		}
		userProjects = append(userProjects, project)
	}

	return userProjects, "Proyectos del usuario", nil
}
