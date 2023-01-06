package v1_api

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"gitlab.com/quick-qr/server/internal/utils"
	"gopkg.in/mcuadros/go-defaults.v1"
	"image"
	_ "image/jpeg"
	"log"
	"regexp"
)

type generateBody struct {
	// Data that will be encoded inside the QR code
	Data string `json:"data" query:"data" validate:"required,max=2953" example:"Hello, world"`
	// Color of the background for the image
	BackgroundColor string `json:"backgroundColor" query:"backgroundColor" validate:"custom_hexcolor" example:"#ffffff" default:"#ffffff"`
	// Color of QR blocks
	ForegroundColor string `json:"foregroundColor" query:"foregroundColor" validate:"custom_hexcolor" example:"#000000" default:"#000000"`
	// Defines the size of the produced image in pixels
	Size int `json:"size" query:"size" validate:"min=128" example:"512" default:"512"`

	RecoveryLevel string `json:"recoveryLevel" query:"recoveryLevel" validate:"oneof=low medium high highest" example:"medium" default:"medium"`
	// Defines size of the quiet zone for the QR code. With bigger border size, the actual size of QR code makes smaller
	BorderSize int `json:"borderSize" query:"borderSize" validate:"ltfield=Size" example:"30" default:"30"`
	// Image to put at the center of QR code
	Logo      *string `json:"logo" query:"logo" example:"base64 string or URL to image"`
	LogoScale float32 `json:"logoScale" query:"logoScale" validate:"gt=0,max=0.25" example:"0.2" default:"0.2"`
}

func (b *generateBody) getLogoData() ([]byte, *httpError) {
	if b.Logo == nil {
		return nil, &httpError{fiber.StatusBadRequest, "No Image data supplied"}
	}

	urlRE := regexp.MustCompile("^(https?:\\/\\/)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b(?:[-a-zA-Z0-9()@:%_\\+.~#?&\\/=]*)$")

	if urlRE.Match([]byte(*b.Logo)) {
		// TODO: fetch image
	}

	decoded, err := base64.StdEncoding.DecodeString(*b.Logo)

	if err != nil {
		return nil, &httpError{fiber.StatusBadRequest, "Invalid format for base64 encoded logo."}
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

func generateFromRequest(req generateBody) ([]byte, *httpError) {
	lvl, _ := utils.StringToRecoveryLevel(req.RecoveryLevel)

	qr, err := qrcode.NewWith(req.Data, qrcode.WithErrorCorrectionLevel(lvl))

	if err != nil {
		return nil, &httpError{500, err.Error()}
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
		b, logoErr := req.getLogoData()
		if logoErr != nil {
			return nil, logoErr
		}

		logo, _, err := image.Decode(bytes.NewReader(b))

		if err != nil {
			return nil, &httpError{400, "Unable to read logo image."}
		}

		options = append(options, standard.WithLogoImage(logo))
	}

	w := standard.NewWithWriter(
		&BufferWriteCloser{bufio.NewWriter(&png)},
		options...,
	)
	saveErr := qr.Save(w)

	if saveErr != nil {
		log.Printf("Failed to export QR code with request: %v", err)
		return nil, &httpError{500, "Some error happened when trying to export QR code."}
	}

	return png.Bytes(), nil
}

func GET_GenerateQR() {

}

// GetGenerateQR godoc
//
//		@Summary		Generate customizable QR code with GET request
//		@Description.markdown generate-qr
//		@Accept			json
//	 	@Param 			q query v1_api.generateBody true "Configuration for the QR code"
//		@Produce		png
//		@Failure		400	{object}	v1_api.errorResponse
//		@Success		201	{string}  string	"Will return generated QR code as PNG"
//		@Router			/v1/generate [get]
func GetGenerateQR(c *fiber.Ctx) error {
	return proccedWithRequestData(c, c.QueryParser)
}

// PostGenerateQR godoc
//
//	@Summary		Generate customizable QR code
//	@Description	This path is alternative to `GET /v1/generate`, all params need to be supplied in body. Refer to `GET` version for any documentation
//	@Param			request	body	v1_api.generateBody	true	"Configuration for QR code generator. Default values are showed below"
//	@Accept			json
//	@Produce		png
//	@Failure		400	{object}	v1_api.errorResponse
//	@Success		201	{string}  string	"Will return generated QR code as PNG"
//	@Router			/v1/generate [post]
func PostGenerateQR(c *fiber.Ctx) error {
	return proccedWithRequestData(c, c.BodyParser)
}

func proccedWithRequestData(c *fiber.Ctx, parser func(out interface{}) error) error {
	payload := new(generateBody)
	defaults.SetDefaults(payload)

	if err := parser(payload); err != nil {
		errMsg := err.Error()

		if jsonError, ok := err.(*fiber.UnmarshalTypeError); ok {
			errMsg = fmt.Sprintf("%v should be of type %v, but received %v", jsonError.Field, jsonError.Type, jsonError.Value)
		}

		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{errMsg})
	}

	if err := utils.ValidateStruct(*payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{*err})
	}

	img, err := generateFromRequest(*payload)

	if err != nil {
		return c.Status(err.Status).JSON(errorResponse{err.Message})
	}

	c.Set("Content-Type", "image/png")
	return c.Status(fiber.StatusCreated).SendString(string(img))
}
