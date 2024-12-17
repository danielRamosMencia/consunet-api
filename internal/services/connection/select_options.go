package connectionservices

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/responses"
)

func SelectOptions(ctx context.Context) ([]responses.ConnectionOption, string, error) {
	var connections []responses.ConnectionOption
	const query = `
	SELECT
		"id",
		"name",
		"code"
	FROM
		"Connection"
	WHERE
		"active" = true;
	`

	rows, err := database.Connx.QueryContext(ctx, query)
	if err != nil {
		log.Println("Error getting connection options", err)
		return nil, "Error al obtener las conexiones", err
	}
	defer rows.Close()

	for rows.Next() {
		var connection responses.ConnectionOption
		err := rows.Scan(
			&connection.Id,
			&connection.Name,
			&connection.Code,
		)
		if err != nil {
			log.Println("Error scanning connection: ", err)
			return nil, "Error escaneando una conexión", err
		}
		connections = append(connections, connection)
	}

	return connections, "Lista de conexiones", nil
}
