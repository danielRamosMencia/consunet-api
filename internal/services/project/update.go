package projectservices

import (
	"context"
	"database/sql"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
)

func Update(ctx context.Context, id string, req requests.UpdateProject) (string, error) {
	const query = `
	UPDATE
		"Project"
	SET
		"name" = $1,
		"code" = $2,
		"connection_id" = $3,
		"updated_at" = CURRENT_TIMESTAMP
	WHERE
		"id" = $4;
	`

	result, err := database.Connx.ExecContext(
		ctx,
		query,
		req.Name,
		req.Code,
		req.Connection_id,
		id,
	)
	if err != nil {
		log.Println("Error updating project", err)
		return "Error al actualizar el proyecto", err
	}

	row, err := result.RowsAffected()
	if err != nil {
		log.Println("Error in result update query: ", err)
		return "Error al actualizar el proyecto", err
	}
	if row == 0 {
		return "Proyecto no encontrado", sql.ErrNoRows
	}

	return "Proyecto actualizado con éxito", nil
}
