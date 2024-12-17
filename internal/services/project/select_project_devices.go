package projectservices

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/responses"
)

func SelectProjectDevices(ctx context.Context, projectId string) ([]responses.ProjectDevices, string, error) {
	var projectDevices []responses.ProjectDevices

	const query = `
	SELECT
		"DP"."id",
		"D"."name" AS "device",
		"A"."name" AS "activity",
		"A"."max_consumption",
		"A"."min_consumption"
	FROM 
		"DevicesOnProject" AS "DP"
	JOIN "Device" AS "D" ON "DP"."device_id" = "D"."id"
	JOIN "Activity" AS "A" ON "DP"."activity_id" = "A"."id"
	WHERE
		"DP"."project_id" = $1;
	`

	rows, err := database.Connx.QueryContext(ctx, query, projectId)
	if err != nil {
		log.Println("Error getting project devices: ", err)
		return nil, "Error al obtener los dispositivos del proyecto", err
	}
	defer rows.Close()

	for rows.Next() {
		var device responses.ProjectDevices

		err := rows.Scan(
			&device.Id,
			&device.Device,
			&device.Activity,
			&device.Max_consumption,
			&device.Min_consumption,
		)
		if err != nil {
			log.Println("Error scanning a device in project: ", err)
		}

		projectDevices = append(projectDevices, device)
	}

	if len(projectDevices) == 0 {
		return nil, "No hay dispositivos agregados al proyecto", nil
	}

	return projectDevices, "Lista de dispositivos en proyecto", nil

}
