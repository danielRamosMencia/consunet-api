package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func SetUpRouter(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173/", // frontend url
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Content-Type,Authorization,Accept,Origin,Access-Control-Request-Method,Access-Control-Request-Headers,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Access-Control-Allow-Methods,Access-Control-Expose-Headers,Access-Control-Max-Age,Access-Control-Allow-Credentials",
		AllowCredentials: true,
	}))
	app.Use(logger.New(logger.Config{}))
	app.Use(requestid.New())
	app.Use(helmet.New())

	v1 := app.Group("api/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "1")
		return c.Next()
	})

	panelv1 := app.Group("api/portal/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "admin-portal-v1")
		return c.Next()
	})

	// web app routes
	HealthCheckRoutes(v1)
	AuthRoutes(v1)
	ConnectionRoutes(v1)
	ProjectRoutes(v1)
	ActivityRoutes(v1)
	DeviceRoutes(v1)
	PermissionsRoutes(v1)
	UserRoutes(v1)

	// admin panel routes
	HealthCheckRoutes(panelv1)
	RoleRoutes(panelv1)
}
