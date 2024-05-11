package validators

import (
	"strconv"

	"github.com/go-playground/validator/v10"
)

func ValidatePort(fl validator.FieldLevel) bool {
	port, err := strconv.Atoi(fl.Field().String())
	if err != nil {
		return false
	}

	return port > 0 && port < 65535
}
