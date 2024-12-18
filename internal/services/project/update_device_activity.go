package projectservices

import (
	"context"
	"database/sql"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
)

func UpdateDeviceActivity(ctx context.Context, deviceProjectId string, req requests.UpdateDeviceActivity) (string, error) {
	const query = `
	UPDATE 
		"DevicesOnProject"
	SET
		"activity_id" = $1
	WHERE
		"id" = $2;
	`

	result, err := database.Connx.ExecContext(
		ctx,
		query,
		req.Activity_id,
		deviceProjectId,
	)
	if err != nil {
		log.Println("Error updating the device activity", err)
		return "Error al actualizar la actividad del dispositivo", err
	}

	row, err := result.RowsAffected()
	if err != nil {
		log.Println("Error in result update query:", err)
		return "Error al actualizar la actividad del dispositivo", err
	}
	if row == 0 {
		return "Dispositivo en proyecto no encontrado", sql.ErrNoRows
	}

	return "Actividad de dispositivo actualizada con éxito", nil
}
