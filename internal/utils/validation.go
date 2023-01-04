package utils

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	validate.RegisterValidation("custom_hexcolor", validateHexColor)
}

func validateHexColor(fl validator.FieldLevel) bool {
	re := regexp.MustCompile("^([0-9A-Fa-f]{6}|[0-9A-Fa-f]{8})$")

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
		return "should be valid hex color string with length of 6 (or 8)"
	case "min":
		return fmt.Sprintf("should be at least %v %v", fe.Param(), minMaxUnits(fe))
	case "max":
		return fmt.Sprintf("should not be greater than %v %v", fe.Param(), minMaxUnits(fe))
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

		reason := fmt.Sprintf("%v %v", name, msgForTag(err))

		return &reason
	}

	return nil
}
