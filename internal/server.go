package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gitlab.com/quick-qr/server/internal/handlers"
	"time"
)

func RunServer(port string) error {
	app := fiber.New(fiber.Config{StrictRouting: false})

	app.Use(cors.New())

	app.Static("/docs/", "./docs/public", fiber.Static{CacheDuration: time.Hour})
	handlers.Register(app)

	return app.Listen(":" + port)
}
