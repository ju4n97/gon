package validators_test

import (
	"testing"

	v10 "github.com/go-playground/validator/v10"
	"github.com/mesatechlabs/gokit/tools/validators"
)

func TestValidatePort(t *testing.T) {
	validate := v10.New(v10.WithRequiredStructEnabled())
	validate.RegisterValidation("port", validators.ValidatePort)

	tests := map[string]bool{
		"8080":  true,
		"3000":  true,
		"300":   true,
		"4200":  true,
		"5000":  true,
		"1234":  true,
		"40000": true,

		"65536": false,
		"-1":    false,
		"foo":   false,
		"true":  false,
		"false": false,
		"300.0": false,
		"300.1": false,
		"300.2": false,
	}

	for value, valid := range tests {
		err := validate.Var(value, "port")

		if err != nil && valid {
			t.Errorf("expected %v to be valid, got error %v", value, err)
		} else if !valid && err == nil {
			t.Errorf("expected %v to be invalid, got no error", value)
		}
	}
}
