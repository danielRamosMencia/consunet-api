package projectservices

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
	"github.com/danielRamosMencia/consunet-api/internal/utils"
)

func Insert(ctx context.Context, userId string, req requests.CreateProject) (string, error) {
	const query = `
	INSERT INTO "Project"
		("id", "name", "code", "user_id", "connection_id")
	VALUES
		($1, $2, $3, $4, $5);
	`

	id := utils.GenerateId()

	_, err := database.Connx.ExecContext(
		ctx,
		query,
		id,
		req.Name,
		req.Code,
		userId,
		req.Connection_id,
	)
	if err != nil {
		log.Println("Error creating project: ", err)
		return "Error al crear el proyecto", err
	}

	return "Proyecto creado con éxito", nil
}
