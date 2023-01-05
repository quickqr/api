package utils

import (
	"encoding/hex"
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

func StringToRecoveryLevel(str string) (qrcode.ErrorCorrectionLevel, bool) {
	c, ok := recoveryLevelsMap[strings.ToLower(str)]
	return c, ok
}

func HexToRGBA(hexString string) color.RGBA {
	b, _ := hex.DecodeString(hexString)

	var alpha uint8 = 255
	if len(b) == 4 {
		alpha = uint8(b[3])
	}

	return color.RGBA{b[0], b[1], b[2], alpha}
}
