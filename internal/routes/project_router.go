package routes

import (
	projectcontrollers "github.com/danielRamosMencia/consunet-api/internal/controllers/project"
	"github.com/danielRamosMencia/consunet-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func ProjectRoutes(router fiber.Router) {
	r := router.Group("/projects")

	r.Get("/", middlewares.AuthRequired, projectcontrollers.GetUserProjects)
	r.Post("/", middlewares.AuthRequired, projectcontrollers.PostProject)
	r.Put("/:id", middlewares.AuthRequired, projectcontrollers.PutProject)
	r.Delete("/:id", middlewares.AuthRequired, projectcontrollers.DeleteProject)

	r.Get("/collabs/:project_id", middlewares.AuthRequired, projectcontrollers.GetCollabs)
	r.Post("/collabs", middlewares.AuthRequired, projectcontrollers.PostCollab)
	r.Delete("/collabs/:id", middlewares.AuthRequired, projectcontrollers.DeleteCollab)
	r.Patch("/collabs/:id", middlewares.AuthRequired, projectcontrollers.PatchCollab)

	r.Get("/devices/:project_id", middlewares.AuthRequired, projectcontrollers.GetProjectDevices)
	r.Post("/devices", middlewares.AuthRequired, projectcontrollers.PostDeviceProject)
	r.Delete("/devices/:id", middlewares.AuthRequired, projectcontrollers.DeleteDeviceProject)
	r.Patch("/devices/:id", middlewares.AuthRequired, projectcontrollers.PacthDeviceActivity)

}
