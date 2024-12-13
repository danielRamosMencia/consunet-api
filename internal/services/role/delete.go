package roleservices

import (
	"context"
	"database/sql"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
)

func Delete(ctx context.Context, id string) (string, error) {
	const query = `
	DELETE FROM
		"Role"
	WHERE
		"id" = $1;
	`

	result, err := database.Connx.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("Error deleting role: ", err)
		return "Error al eliminar el rol", err
	}

	row, err := result.RowsAffected()
	if err != nil {
		log.Println("Error in result delete query: ", err)
		return "Error al eliminar el rol", err
	}
	if row == 0 {
		return "Rol no encontrado", sql.ErrNoRows
	}

	return "Rol eliminado con éxito", nil
}
