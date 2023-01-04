package utils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"image/color"
	"strings"

	"github.com/skip2/go-qrcode"
)

var (
	recoveryLevelsMap = map[string]qrcode.RecoveryLevel{
		"low":    qrcode.Low,
		"medium": qrcode.Medium,
		"high":   qrcode.High,
		"delete": qrcode.Highest,
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

func StringToRecoveryLevel(str string) (qrcode.RecoveryLevel, bool) {
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
