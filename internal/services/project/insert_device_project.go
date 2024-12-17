package projectservices

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
	"github.com/danielRamosMencia/consunet-api/internal/utils"
)

func InsertDeviceProject(ctx context.Context, assignedBy string, req requests.CreateDeviceProject) (string, error) {
	const query = `
	INSERT INTO "DevicesOnProject"
		("id", "device_id", "project_id", "activity_id", "assigned_by")
	VALUES ($1, $2, $3, $4, $5);
	`

	id := utils.GenerateId()

	_, err := database.Connx.ExecContext(
		ctx,
		query,
		id,
		req.Device_id,
		req.Project_id,
		req.Activity_id,
		assignedBy,
	)

	if err != nil {
		log.Println("Error adding device to project: ", err)
		return "Error al agregar dispositivo", err
	}

	return "Dispositivo agregado correctamente", nil
}
