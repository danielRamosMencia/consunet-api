package roleservices

import (
	"context"
	"database/sql"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
)

func Update(ctx context.Context, id string, req requests.CreateRole) (string, error) {
	const query = `
	UPDATE
		"Role"
	SET
		"name" = $1,
		"code" = $2,
		"active" = $3,
		"updated_at" = CURRENT_TIMESTAMP
	WHERE
		"id" = $4
	`

	if req.Active == nil {
		defaultTrue := true
		req.Active = &defaultTrue
	}

	result, err := database.Connx.ExecContext(
		ctx,
		query,
		req.Name,
		req.Code,
		req.Active,
		id,
	)
	if err != nil {
		log.Println("Error updating role: ", err)
		return "Error al actualizar el rol", err
	}

	row, err := result.RowsAffected()
	if err != nil {
		log.Println("Error in result update query: ", err)
		return "Error al actualizar el rol", err
	}
	if row == 0 {
		return "Rol no encontrado", sql.ErrNoRows
	}

	return "Rol actualizado con éxito", nil
}
