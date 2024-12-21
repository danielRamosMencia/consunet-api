package projectservices

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
	"github.com/danielRamosMencia/consunet-api/internal/utils"
)

func InsertCollab(ctx context.Context, assignedBy string, req requests.CreateUserProject) (string, error) {
	const query = `
	INSERT INTO "UsersOnProject"
		("id", "user_id", "project_id", "permission_id", "assigned_by")
	VALUES ($1, $2, $3, $4, $5);
	`

	id := utils.GenerateId()

	_, err := database.Connx.ExecContext(
		ctx,
		query,
		id,
		req.User_id,
		req.Project_id,
		req.Permission_id,
		assignedBy,
	)
	if err != nil {
		log.Println("Error adding collab to project: ", err)
		return "Error al agregar el colaborador", err
	}

	return "Colaborador agregado con éxito", nil
}
