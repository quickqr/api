package v1_api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"gitlab.com/quick-qr/server/internal/utils"
)

type generateBody struct {
	Data            string `json:"data" validate:"required,max=2953" example:"Some data to encode"`
	BackgroundColor string `json:"backgroundColor" validate:"custom_hexcolor" example:"ffffff"`
	ForegroundColor string `json:"foregroundColor" validate:"custom_hexcolor" example:"000000"`
	Size            int    `json:"size" validate:"min=128" example:"512"`
	RecoveryLevel   string `json:"recoveryLevel" validate:"oneof=low medium high highest" example:"medium"`
	//DisableBorder   bool   `json:"disableBorder" example:"false"`
}

type BufferWriteCloser struct {
	*bufio.Writer
}

func (bwc *BufferWriteCloser) Close() error {
	// Noop
	return nil
}

func generateFromRequest(req generateBody) ([]byte, error) {
	lvl, _ := utils.StringToRecoveryLevel(req.RecoveryLevel)

	qr, err := qrcode.NewWith(req.Data, qrcode.WithErrorCorrectionLevel(lvl))

	if err != nil {
		return nil, err
	}

	var png bytes.Buffer

	w := standard.NewWithWriter(
		&BufferWriteCloser{bufio.NewWriter(&png)},
		standard.WithBgColorRGBHex(req.BackgroundColor),
		standard.WithFgColorRGBHex(req.ForegroundColor),
	)
	saveErr := qr.Save(w)

	return png.Bytes(), saveErr
}

// GenerateQR godoc
//
//	@Summary		Get user list
//	@Description.markdown	generate-qr
//	@Param			request	body	v1_api.generateBody	true	"Configuration for QR code generator. Default values are showed below"
//	@Accept			json
//	@Produce		png
//	@Failure		400	{object}	errorResponse
//	@Success		201	{object}	string	"Will return generated QR code as PNG"
//	@Router			/v1/generate [post]
func GenerateQR(c *fiber.Ctx) error {
	payload := generateBody{
		BackgroundColor: "#ffffff",
		ForegroundColor: "#000000",
		//DisableBorder:   false,
		Size:          512,
		RecoveryLevel: "medium",
	}

	if err := c.BodyParser(&payload); err != nil {
		errMsg := err.Error()

		if jsonError, ok := err.(*json.UnmarshalTypeError); ok {
			errMsg = fmt.Sprintf("%v should be of type %v, but received %v", jsonError.Field, jsonError.Type, jsonError.Value)
		}

		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{errMsg})
	}

	if err := utils.ValidateStruct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{*err})
	}

	img, err := generateFromRequest(payload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{err.Error()})
	}

	c.Set("Content-Type", "image/png")
	return c.Status(fiber.StatusCreated).SendString(string(img))
}
