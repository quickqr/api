package utils

import (
	"errors"
	"fmt"
	"github.com/yeqown/go-qrcode/v2"
	"image/color"
	"strings"
)

var (
	recoveryLevelsMap = map[string]qrcode.ErrorCorrectionLevel{
		"low":    qrcode.ErrorCorrectionLow,
		"medium": qrcode.ErrorCorrectionMedium,
		"high":   qrcode.ErrorCorrectionQuart,
		"delete": qrcode.ErrorCorrectionHighest,
	}
)

type ConvertibleBoolean bool

func (bit *ConvertibleBoolean) UnmarshalJSON(data []byte) error {
	asString := string(data)
	if asString == "1" || asString == "true" {
		*bit = true
	} else if asString == "0" || asString == "false" {
		*bit = false
	} else {
		return errors.New(fmt.Sprintf("Boolean unmarshal error: invalid input %s", asString))
	}

	return nil
}

// TODO: Move this code into own library for qr codes.

// HexToRGBA converts # prefixed hex colors of length 3, 4 or 8 (without #) to color.RGBA
func HexToRGBA(hexString string) (c color.RGBA) {
	c.A = 0xff

	// TODO: Some error handling?
	switch len(hexString) {
	case 9:
		fmt.Sscanf(hexString, "#%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	case 7:
		fmt.Sscanf(hexString, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		fmt.Sscanf(hexString, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		// TODO: throw an error
	}

	return
}

func StringToRecoveryLevel(str string) (qrcode.ErrorCorrectionLevel, bool) {
	c, ok := recoveryLevelsMap[strings.ToLower(str)]
	return c, ok
}
