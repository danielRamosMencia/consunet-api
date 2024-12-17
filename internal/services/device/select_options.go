package deviceservices

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/shared"
)

func SelectOptions(ctx context.Context) ([]shared.Options, string, error) {
	var deviceOptions []shared.Options

	const query = `
	SELECT
		"id",
		"name",
		"code"
	FROM 
		"Device"
	WHERE 
		"active" = true;
	`

	rows, err := database.Connx.QueryContext(ctx, query)
	if err != nil {
		log.Panicln("Error getting devices: ", err)
		return deviceOptions, "Error obteniendo los dispositivos", err
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
			log.Println("Error scanning option: ", err)
			return deviceOptions, "Error escaneando un dispositivo", err
		}

		deviceOptions = append(deviceOptions, option)
	}

	return deviceOptions, "Lista de dispositivos", nil
}
