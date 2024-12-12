package roleservices

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
	"github.com/danielRamosMencia/consunet-api/internal/utils"
)

func Insert(ctx context.Context, req requests.CreateRole) (string, error) {
	const query = `
	INSERT INTO "Role"
		("id", "name", "code", "active")
	VALUES
		($1, $2, $3, $4);
	`

	id := utils.GenerateId()

	if req.Active == nil {
		defaultTrue := true
		req.Active = &defaultTrue
	}

	result, err := database.Connx.ExecContext(
		ctx,
		query,
		id,
		req.Name,
		req.Code,
		*req.Active,
	)
	if err != nil {
		log.Println("Error creating role: ", err)
		return "Error al crear el rol", err
	}

	log.Print(result)
	return "Rol creado con éxito", nil
}
