package utils

import (
	"github.com/quickqr/gqr"
	"strings"
)

var (
	recoveryLevelsMap = map[string]gqr.ErrorCorrectionLevel{
		"low":    gqr.ErrorCorrectionLow,
		"medium": gqr.ErrorCorrectionMedium,
		"high":   gqr.ErrorCorrectionQuart,
		"delete": gqr.ErrorCorrectionHighest,
	}
)

func StringToRecoveryLevel(str string) (gqr.ErrorCorrectionLevel, bool) {
	c, ok := recoveryLevelsMap[strings.ToLower(str)]
	return c, ok
}
