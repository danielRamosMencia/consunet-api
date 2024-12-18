package projectservices

import (
	"context"
	"database/sql"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
)

func DeleteCollab(ctx context.Context, collabId string) (string, error) {
	const query = `
	DELETE FROM
		"UsersOnProject"
	WHERE
		"id" = $1;
	`

	result, err := database.Connx.ExecContext(ctx, query, collabId)
	if err != nil {
		log.Println("Error removing collab from project: ", err)
		return "Error al quitar el colaborador", nil
	}

	row, err := result.RowsAffected()
	if err != nil {
		log.Println("Error in result delete query: ", err)
		return "Error al quitar el colaborador", err
	}
	if row == 0 {
		return "Colaborador en proyecto no encontrado", sql.ErrNoRows
	}

	return "Colaborador quitado con éxito", nil
}
