package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gitlab.com/quick-qr/server/internal/handlers"
)

func RunServer(port string) error {
	app := fiber.New()

	app.Get("/docs/*", swagger.HandlerDefault) // default

	handlers.Register(app)

	return app.Listen(":" + port)
}
