package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func SetUpRouter(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{}))
	app.Use(requestid.New())
	app.Use(helmet.New())

	v1 := app.Group("api/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "1")
		return c.Next()
	})

	HealthCheckRoutes(v1)
}
