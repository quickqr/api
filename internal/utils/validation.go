package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorResponse struct {
	Reason string `json:"reason"`
}

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

func ValidateStruct[T any](s T) *ErrorResponse {
	err := validate.Struct(s)

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		err := ve[0]

		var response ErrorResponse
		response.Reason = fmt.Sprintf("%v %v", err.Field(), msgForTag(err))

		return &response

	}

	return nil
}
