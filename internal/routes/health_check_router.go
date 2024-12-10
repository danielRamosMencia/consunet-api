package routes

import "github.com/gofiber/fiber/v2"

func HealthCheckRoutes(router fiber.Router) {
	r := router.Group("/health-check")

	r.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Server ready to Go",
		})
	})
}
