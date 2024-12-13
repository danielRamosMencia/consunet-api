package roleservices

import (
	"context"
	"database/sql"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
)

var (
	responseMessage string
	prefix          string
)

func UpdateActive(ctx context.Context, id string, req bool) (string, error) {
	const query = `
	UPDATE
		"Role"
	SET
		"active" = $1,
		"updated_at" = CURRENT_TIMESTAMP
	WHERE
		"id" = $2;
	`

	if req {
		prefix = "activar"
		responseMessage = "Rol activado con éxito"
	} else {
		prefix = "inactivar"
		responseMessage = "Rol desactivado con éxito"
	}

	result, err := database.Connx.ExecContext(
		ctx,
		query,
		req,
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
