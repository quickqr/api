package v1_api

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"gitlab.com/quick-qr/server/internal/utils"
	"image"
	_ "image/jpeg"
	"regexp"
)

type generateBody struct {
	Data            string  `json:"data" validate:"required,max=2953" example:"Some data to encode"`
	BackgroundColor string  `json:"backgroundColor" validate:"custom_hexcolor" example:"#ffffff"`
	ForegroundColor string  `json:"foregroundColor" validate:"custom_hexcolor" example:"#000000"`
	Size            int     `json:"size" validate:"min=128" example:"512"`
	RecoveryLevel   string  `json:"recoveryLevel" validate:"oneof=low medium high highest" example:"medium"`
	BorderSize      int     `json:"borderSize" validate:"ltfield=Size" example:"30"`
	Logo            *string `json:"logo" example:"base64 string or URL to image"`
	LogoScale       float32 `json:"logoScale" validate:"gt=0,max=0.25" example:"0.2"`
}

func (b *generateBody) getLogoData() ([]byte, error) {
	if b.Logo == nil {
		return nil, errors.New("No image data supplied")
	}

	urlRE := regexp.MustCompile("^(https?:\\/\\/)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b(?:[-a-zA-Z0-9()@:%_\\+.~#?&\\/=]*)$")

	if urlRE.Match([]byte(*b.Logo)) {
		fmt.Println("url")
		// TODO: fetch image
	}

	decoded, err := base64.StdEncoding.DecodeString(*b.Logo)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Failed to read base64 data.")
	}

	return decoded, nil
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

	options := []standard.ImageOption{
		standard.WithBgColorRGBHex(req.BackgroundColor),
		standard.WithFgColorRGBHex(req.ForegroundColor),
		standard.WithLogoScale(req.LogoScale),
		standard.WithImageSize(uint(req.Size)),
	}

	if req.BorderSize >= 0 {
		options = append(options, standard.WithBorderWidth(req.BorderSize))
	}

	if req.Logo != nil {
		b, err := req.getLogoData()
		if err != nil {
			return nil, err
		}

		logo, _, err := image.Decode(bytes.NewReader(b))

		if err != nil {
			return nil, err
		}

		options = append(options, standard.WithLogoImage(logo))
	}

	w := standard.NewWithWriter(
		&BufferWriteCloser{bufio.NewWriter(&png)},
		options...,
	)
	saveErr := qr.Save(w)

	return png.Bytes(), saveErr
}

// GenerateQR godoc
//
//	@Summary		Generate customizable QR code
//	@Description.markdown	generate-qr
//	@Param			request	body	v1_api.generateBody	true	"Configuration for QR code generator. Default values are showed below"
//	@Accept			json
//	@Produce		png
//	@Failure		400	{object}	v1_api.errorResponse
//	@Success		201	{object}	string	"Will return generated QR code as PNG"
//	@Router			/v1/generate [post]
func GenerateQR(c *fiber.Ctx) error {
	payload := generateBody{
		BackgroundColor: "#ffffff",
		ForegroundColor: "#000000",
		// values less than 0 does not count
		BorderSize:    -1,
		LogoScale:     0.2,
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

	// TODO: Return struct with status code to differentiate 5xx and 4xx instead of single status code below
	img, err := generateFromRequest(payload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{err.Error()})
	}

	c.Set("Content-Type", "image/png")
	return c.Status(fiber.StatusCreated).SendString(string(img))
}
