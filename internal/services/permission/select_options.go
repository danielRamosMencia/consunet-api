package permissionservices

import (
	"context"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/shared"
)

func SelectOptions(ctx context.Context) ([]shared.Options, string, error) {
	var permissionOptions []shared.Options

	const query = `
	SELECT
		"id",
		"name",
		"code"
	FROM
		"Permission"
	WHERE
		"active" = true;
	`

	rows, err := database.Connx.QueryContext(ctx, query)
	if err != nil {
		log.Println("Error getting permissions: ", err)
		return nil, "Error al obtener los permisos", err
	}
	defer rows.Close()

	for rows.Next() {
		var option shared.Options

		err := rows.Scan(
			&option.Id,
			&option.Name,
			&option.Code,
		)
		if err != nil {
			log.Println("Error scanning permission: ", err)
			return nil, "Error al escanear un permiso", err
		}

		permissionOptions = append(permissionOptions, option)
	}

	return permissionOptions, "Lista de permisos", nil
}
