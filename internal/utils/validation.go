package utils

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	errHex := validate.RegisterValidation("custom_hexcolor", validateHexColor)
	if errHex != nil {
		log.Fatal("Failed to register custom_hexcolor validation tag")
	}
}

func validateHexColor(fl validator.FieldLevel) bool {
	re := regexp.MustCompile("^#([0-9A-Fa-f]{3}|[0-9A-Fa-f]{6})$")

	return re.Match([]byte(fl.Field().String()))
}

func minMaxUnits(fe validator.FieldError) string {
	if fe.Type().Name() == "string" {
		return "characters"
	}

	return ""
}

func msgForTag(fe validator.FieldError) string {

	switch fe.Tag() {
	case "required":
		return "is required"
	case "custom_hexcolor":
		return "should be valid RGB hex color"
	case "ltfield":
		return "should be less than field " + fe.Param()
	case "min":
		return fmt.Sprintf("should be at least %v %v", fe.Param(), minMaxUnits(fe))
	case "max":
		return fmt.Sprintf("should not be greater than %v %v", fe.Param(), minMaxUnits(fe))
	case "gt":
		return fmt.Sprintf("should be greater than %v %v", fe.Param(), minMaxUnits(fe))
	case "oneof":
		list := strings.Join(strings.Split(fe.Param(), " "), ", ")
		return "should be on of following values: " + list
	}

	return fe.Error() // default error
}

func ValidateStruct[T any](s T) *string {
	err := validate.Struct(s)

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		err := ve[0]

		// Getting the JSON name for the errored field
		field, _ := reflect.TypeOf(s).FieldByName(err.StructField())
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

		reason := strings.TrimSpace(fmt.Sprintf("%v %v", name, msgForTag(err)))

		return &reason
	}

	return nil
}
