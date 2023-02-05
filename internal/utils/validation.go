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
	_ = validate.RegisterValidation("custom_hexcolor", validateHexColor)
}

func validateHexColor(fl validator.FieldLevel) bool {
	re := regexp.MustCompile("^#([0-9A-Fa-f]{3}|[0-9A-Fa-f]{6}|[0-9A-Fa-f]{8})$")

	return re.Match([]byte(fl.Field().String()))
}

func minMaxUnits(fe validator.FieldError) string {
	if fe.Type().Name() == "string" {
		return "characters"
	}
	if fe.Type().String() == "[]string" {
		return "items"
	}

	return ""
}

func getJsonName(rt reflect.Type, s string) string {
	field, _ := rt.FieldByName(s)
	return strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
}

func msgForTag(rt reflect.Type, fe validator.FieldError) string {
	// Getting the JSON name for the errored field
	name := getJsonName(rt, fe.StructField())
	msg := ""

	switch fe.Tag() {
	case "required":
		msg = "is required"
		break
	case "custom_hexcolor":
		msg = "hex color should have length of 3, 6 or 8 prefixed with #"
		break
	case "ltfield":
		msg = "should be less than field " + getJsonName(rt, fe.Param())
		break
	case "min":
		msg = fmt.Sprintf("should be at least %v %v", fe.Param(), minMaxUnits(fe))
		break
	case "max":
		msg = fmt.Sprintf("should not be greater than %v %v", fe.Param(), minMaxUnits(fe))
		break
	case "gt":
		msg = fmt.Sprintf("should be greater than %v %v", fe.Param(), minMaxUnits(fe))
		break
	case "oneof":
		list := strings.Join(strings.Split(fe.Param(), " "), ", ")
		msg = "should be one of following values: " + list
		break
	default:
		msg = fe.Error()
	}

	return fmt.Sprintf("%v %v", name, msg)
}

func ValidateStruct[T any](s T) *string {
	err := validate.Struct(s)

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		err := ve[0]

		reason := strings.TrimSpace(msgForTag(reflect.TypeOf(s), err))

		return &reason
	}

	return nil
}
