package internal

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/quick-qr/server/internal/handlers"
)

func RunServer(port string) error {
	app := fiber.New()

	handlers.Register(app)

	return app.Listen(":" + port)
}
