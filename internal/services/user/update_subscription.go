package userservices

import (
	"context"
	"database/sql"
	"log"

	"github.com/danielRamosMencia/consunet-api/internal/database"
	"github.com/danielRamosMencia/consunet-api/internal/models/requests"
)

func UpdateSubscription(ctx context.Context, userId string, req requests.UpdateUserSubscription) (string, error) {
	const query = `
	UPDATE 
		"User"
	SET
		"subscription_id" = $1
	WHERE 
		"id" = $2;
	`

	result, err := database.Connx.ExecContext(
		ctx,
		query,
		req.Subscription_id,
		userId,
	)
	if err != nil {
		log.Println("Error updating user subscription: ", err)
		return "Error al añadir la subscripción al usuario", err
	}

	row, err := result.RowsAffected()
	if err != nil {
		log.Println("Error in result update query: ", err)
		return "Error al añadir la subscripción al usuario", err
	}
	if row == 0 {
		log.Println("User not found")
		return "Usuario para agrear subscripción no encontrado", sql.ErrNoRows
	}

	return "Subscripción añadida con éxito", nil
}
