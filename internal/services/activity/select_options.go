package activityservices

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/shared"
)

func SelectOptions(ctx context.Context) ([]shared.Options, string, error) {
	var activityOptions []shared.Options

	const query = `
	SELECT
		"id",
		"name",
		"code"
	FROM 
		"Activity"
	WHERE 
		"active" = true;
	`

	rows, err := database.Connx.QueryContext(ctx, query)
	if err != nil {
		log.Println("Error getting activities options", err)
		return nil, "Error al obtener las actividades", err
	}
	defer rows.Close()

	for rows.Next() {
		var option shared.Options

		err := rows.Scan(
			&option.Id,
			&option.Name,
			&option.Code,
		)
		if err != nil {
			log.Println("Error scanning activity", err)
			return nil, "Error escaneando una actividad", err
		}

		activityOptions = append(activityOptions, option)
	}

	return activityOptions, "Lista de actividades", nil
}
