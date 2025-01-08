package projectservices

import (
	"context"
	"database/sql"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/responses"
)

func SelectOne(ctx context.Context, id string) (responses.Project, string, error) {
	const query = `
	SELECT
		"P"."id",
		"P"."code",
		"P"."name",
		"C"."name" AS "connection_name",
		"P"."created_at",
		"P"."updated_at"
	FROM "Project" AS "P"
	JOIN "Connection" AS "C" ON "P"."connection_id" = "C"."id"
	WHERE "P"."id" = $1;
	`

	var project responses.Project

	row := database.Connx.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&project.Id,
		&project.Name,
		&project.Code,
		&project.ConnectionName,
		&project.CreatedAt,
		&project.UpdatedAt,
	)

	switch {
	case err == sql.ErrNoRows:
		log.Println("Project with id: " + id + " not found")
		return project, "Proyecto no encontrado", err
	case err != nil:
		log.Println("Error getting role: ", err)
		return project, "Error al obtener el proyecto", err
	}

	return project, "Proyecto encontrado", nil
}
