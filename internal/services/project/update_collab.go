package projectservices

import (
	"context"
	"database/sql"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
)

func UpdateCollab(ctx context.Context, collabId string, req requests.UpdateCollab) (string, error) {
	const query = `
	UPDATE
		"UsersOnProject"
	SET
		"permission_id" = $1
	WHERE
		"id" = $2;
	`

	result, err := database.Connx.ExecContext(ctx,
		query,
		req.Permission_id,
		collabId,
	)
	if err != nil {
		log.Println("Error updating collab permission", err)
		return "Error al actualizar el permiso del colaborador", err
	}

	row, err := result.RowsAffected()
	if err != nil {
		log.Println("Error in result update query: ", err)
		return "Error al actualizar el permiso del colaborador", err
	}
	if row == 0 {
		log.Println("Collab not found")
		return "Colaborador no encontrado", sql.ErrNoRows
	}

	return "Permiso en colaborador actualizado con éxito", nil
}
