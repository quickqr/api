package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/quick-qr/server/internal/handlers/v1_api"
)

func Register(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/generate", v1_api.GenerateQR)
}
