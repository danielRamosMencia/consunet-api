package projectservices

import (
	"context"
	"database/sql"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
)

func DeleteDeviceProject(ctx context.Context, id string) (string, error) {
	const query = `
	DELETE FROM 
		"DevicesOnProject"
	WHERE
		"id" = $1;
	`

	result, err := database.Connx.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("Error removing device from project: ", err)
		return "Error al quitar el dispositivo", nil
	}

	row, err := result.RowsAffected()
	if err != nil {
		log.Println("Error in result delete query: ", err)
		return "Error al quitar el dispositivo", err
	}
	if row == 0 {
		return "Dispositivo en proyecto no encontrado", sql.ErrNoRows
	}

	return "Dispositivo quitado con éxito", nil
}
