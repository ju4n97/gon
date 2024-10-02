package custom_validator

import (
	"reflect"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func ValidatePort(fl validator.FieldLevel) bool {
	var port int
	var err error

	switch fl.Field().Kind() {
	case reflect.String:
		port, err = strconv.Atoi(fl.Field().String())
		if err != nil {
			return false
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		port = int(fl.Field().Int())
	default:
		return false
	}

	if err != nil {
		return false
	}

	return port > 0 && port < 65535
}
