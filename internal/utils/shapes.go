package utils

import (
	"fmt"
	"github.com/quickqr/gqr/export/image/shapes"
	"strings"
)

var (
	roundnessLevels = map[string]float64{
		"square":  0,
		"rounded": 0.3,
		"circle":  0.5,
	}
)

// StringToFinderShape converts string to shapes.FinderDrawConfig, assuming the input is already valid
func StringToFinderShape(str string) shapes.FinderDrawConfig {
	c, _ := roundnessLevels[strings.ToLower(str)]
	return shapes.RoundedFinderShape(c)
}

// StringToModuleDrawer converts string to shapes.ModuleDrawer, assuming the input is already valid
func StringToModuleDrawer(str string) shapes.ModuleDrawer {
	fmt.Println(str)
	switch str {
	case "fluid":
		return shapes.RoundedModuleShape(0.5, true)
	case "hline", "vline":
		fmt.Println(str)
		return shapes.LineModuleShape(0.5, str == "vline")
	}

	c, _ := roundnessLevels[strings.ToLower(str)]
	return shapes.RoundedModuleShape(c, false)
}
