package projectservices

import (
	"context"
	"database/sql"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
)

func Delete(ctx context.Context, id string) (string, error) {
	const query = `
	DELETE FROM 
		"Project"
	WHERE
		"id" = $1;
	`

	result, err := database.Connx.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("Error deleting project: ", err)
		return "Error al eliminar el proyecto", nil
	}

	row, err := result.RowsAffected()
	if err != nil {
		log.Println("Error in result delete query: ", err)
		return "Error al eliminar el proyecto", err
	}
	if row == 0 {
		return "Proyecto no encontrado", sql.ErrNoRows
	}

	return "Proyecto eliminado con éxito", nil
}
