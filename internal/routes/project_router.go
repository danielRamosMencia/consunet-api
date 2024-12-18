package routes

import (
	projectcontrollers "github.com/danielRamosMencia/consunet-api/internal/controllers/project"
	"github.com/danielRamosMencia/consunet-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func ProjectRoutes(router fiber.Router) {
	r := router.Group("/projects")

	r.Get("/", middlewares.AuthRequired, projectcontrollers.GetUserProjects)
	r.Get("/collabs/:project_id", middlewares.AuthRequired, projectcontrollers.GetCollabs)
	r.Get("/devices/:project_id", middlewares.AuthRequired, projectcontrollers.GetProjectDevices)
	r.Post("/", middlewares.AuthRequired, projectcontrollers.PostProject)
	r.Post("/devices", middlewares.AuthRequired, projectcontrollers.PostDeviceProject)
	r.Post("/users", middlewares.AuthRequired, projectcontrollers.PostUserProject)
	r.Patch("/devices/:id", middlewares.AuthRequired, projectcontrollers.PacthDeviceActivity)
	r.Delete("/devices/:id", middlewares.AuthRequired, projectcontrollers.DeleteDeviceProject)
}
