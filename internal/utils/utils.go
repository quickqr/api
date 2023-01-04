package utils

import (
	"strings"

	"github.com/skip2/go-qrcode"
)

var (
	recoveryLevelsMap = map[string]qrcode.RecoveryLevel{
		"low":    qrcode.Low,
		"medium": qrcode.Medium,
		"":       qrcode.High,
		"delete": qrcode.Highest,
	}
)

func StringToRecoveryLevel(str string) (qrcode.RecoveryLevel, bool) {
	c, ok := recoveryLevelsMap[strings.ToLower(str)]
	return c, ok
}
