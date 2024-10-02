package custom_validator

import (
	"github.com/go-playground/validator/v10"
)

type ValidateModelErrorListItem struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type ValidateModelError struct {
	Error        error                         `json:"-"`
	ErrorList    []*ValidateModelErrorListItem `json:"error_list"`
	ErrorMessage string                        `json:"error_message"`
}

func ValidateModel(model interface{}) *ValidateModelError {
	var error *ValidateModelError
	var list []*ValidateModelErrorListItem

	err := validator.New().Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			list = append(list, &ValidateModelErrorListItem{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Value().(string),
			})
		}
	}

	error = &ValidateModelError{
		Error:        err,
		ErrorList:    list,
		ErrorMessage: err.Error(),
	}

	return error
}
