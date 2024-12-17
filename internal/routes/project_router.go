package routes

import (
	projectcontrollers "github.com/danielRamosMencia/consunet-api/internal/controllers/project"
	"github.com/danielRamosMencia/consunet-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func ProjectRoutes(router fiber.Router) {
	r := router.Group("/projects")

	r.Get("/user", middlewares.AuthRequired, projectcontrollers.GetUserProjects)
	r.Post("/", middlewares.AuthRequired, projectcontrollers.PostProject)
	r.Post("/device", middlewares.AuthRequired, projectcontrollers.PostDeviceProject)
	r.Post("/user", middlewares.AuthRequired, projectcontrollers.PostUserProject)
}
