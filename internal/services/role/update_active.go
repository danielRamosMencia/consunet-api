package roleservices

import (
	"context"
	"database/sql"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
)

var (
	responseMessage string = "Rol desactivado con éxito"
	prefix          string = "inactivar"
)

func UpdateActive(ctx context.Context, id string, req requests.ToggleActive) (string, error) {
	const query = `
	UPDATE
		"Role"
	SET
		"active" = $1,
		"updated_at" = CURRENT_TIMESTAMP
	WHERE
		"id" = $2;
	`

	if req.Active {
		prefix = "activar"
		responseMessage = "Rol activado con éxito"
	}

	result, err := database.Connx.ExecContext(
		ctx,
		query,
		req.Active,
		id,
	)
	if err != nil {
		log.Println("Error changing active status", err)
		return "Error al" + prefix + "el rol", err
	}

	row, err := result.RowsAffected()
	if err != nil {
		log.Println("Error in result update active query: ", err)
		return "Error al" + prefix + "el rol", err
	}
	if row == 0 {
		return "Rol no encontrado", sql.ErrNoRows
	}

	return responseMessage, nil
}
