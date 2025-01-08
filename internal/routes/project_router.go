package routes

import (
	projectcontrollers "github.com/danielRamosMencia/consunet-api/internal/controllers/project"
	"github.com/danielRamosMencia/consunet-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func ProjectRoutes(router fiber.Router) {
	r := router.Group("/projects")

	r.Use(middlewares.AuthRequired)

	r.Get("/", projectcontrollers.GetUserProjects)
	r.Get("/:id", projectcontrollers.GetProject)
	r.Post("/", projectcontrollers.PostProject)
	r.Put("/:id", projectcontrollers.PutProject)
	r.Delete("/:id", projectcontrollers.DeleteProject)

	r.Get("/collabs/:project_id", projectcontrollers.GetCollabs)
	r.Post("/collabs", projectcontrollers.PostCollab)
	r.Delete("/collabs/:id", projectcontrollers.DeleteCollab)
	r.Patch("/collabs/:id", projectcontrollers.PatchCollab)

	r.Get("/devices/:project_id", projectcontrollers.GetProjectDevices)
	r.Post("/devices", projectcontrollers.PostDeviceProject)
	r.Delete("/devices/:id", projectcontrollers.DeleteDeviceProject)
	r.Patch("/devices/:id", projectcontrollers.PacthDeviceActivity)
}
