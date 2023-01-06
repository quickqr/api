package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gitlab.com/quick-qr/server/internal/handlers"
	"time"
)

func RunServer(port string) error {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Redirect("/docs/index.html")
	})

	app.Static("/docs/", "./docs/public", fiber.Static{CacheDuration: time.Hour})
	handlers.Register(app)

	return app.Listen(":" + port)
}
