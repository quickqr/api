package v1_api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/quick-qr/server/internal/utils"
)

type generateBody struct {
	Data            string `json:"data" validate:"required"`
	BackgroundColor string `json:"backgroundColor" validate:"hexcolor"`
	ForegroundColor string `json:"foregroundColor" validate:"hexcolor"`
	Size            int    `json:"size,string" validate:"min=128"`
	RecoveryLevel   string `json:"recoveryLevel" validate:"oneof=low medium high highest"`
}

func GenerateQR(c *fiber.Ctx) error {
	payload := generateBody{
		BackgroundColor: "#ffffff",
		ForegroundColor: "#000000",
		Size:            256,
		RecoveryLevel:   "medium",
	}

	if err := c.BodyParser(&payload); err != nil {
		fmt.Println(err)

		return c.Status(400).JSON(errorResponse{err.Error()})
	}

	err := utils.ValidateStruct(payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)

	}

	return c.JSON(payload)
}
