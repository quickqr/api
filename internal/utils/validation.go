package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is required"
	case "hexcolor":
		return "should be hex color"
	case "min":
		return "should be at least " + fe.Param()
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
