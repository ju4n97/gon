package env

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// EnvironmentVariable is a struct that holds the key and value of an environment variable.
// It provides methods to convert the value to different types.
type EnvironmentVariable struct {
	key   string
	value interface{}
}

// New creates a new `EnvironmentVariable` struct with the given key.
// If the environment variable is not found in the operating system, it will
// return a new `EnvironmentVariable` struct with `value` set to `nil`.
func New(key string) *EnvironmentVariable {
	foundValue, isFound := os.LookupEnv(key)
	if !isFound {
		return &EnvironmentVariable{key: key, value: nil}
	}

	return &EnvironmentVariable{key: key, value: foundValue}
}

// WithDefault modifies the existing `EnvironmentVariable` struct with the given key
// and a default value. If the environment variable is not found in the
// operating system, it will return a modified `EnvironmentVariable` struct with
// `value` set to `defaultValue`.
func (e *EnvironmentVariable) WithDefault(defaultValue interface{}) *EnvironmentVariable {
	foundValue, isFound := os.LookupEnv(e.key)
	if !isFound {
		return &EnvironmentVariable{key: e.key, value: defaultValue}
	}

	return &EnvironmentVariable{key: e.key, value: foundValue}
}

// ToString converts the value of the environment variable to a string.
func (e *EnvironmentVariable) ToString() (string, error) {
	if e.value == nil {
		return "", fmt.Errorf("no value found for %s", e.key)
	}

	return fmt.Sprintf("%v", e.value), nil
}

// MustToString is like `ToString`, but it panics if the value is not found.
func (e *EnvironmentVariable) MustToString() string {
	value, err := e.ToString()
	if err != nil {
		log.Panicf("failed to convert %s to string: %s", e.key, err)
	}

	return value
}

// ToBool converts the value of the environment variable to a boolean value.
func (e *EnvironmentVariable) ToBool() (bool, error) {
	value, err := e.ToString()
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(value)
}

// MustToBool is like `ToBool`, but it panics if the value is not found or there's an error in the conversion.
func (e *EnvironmentVariable) MustToBool() bool {
	value, err := strconv.ParseBool(e.MustToString())
	if err != nil {
		log.Panicf("failed to parse %s To bool: %s", e.key, err)
	}

	return value
}

// ToInt converts the value of the environment variable to an integer value.
func (e *EnvironmentVariable) ToInt() (int, error) {
	value, err := e.ToString()
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(value)
}

// MustToInt is like `ToInt`, but it panics if the value is not found or there's an error in the conversion.
func (e *EnvironmentVariable) MustToInt() int {
	value, err := strconv.Atoi(e.MustToString())
	if err != nil {
		log.Panicf("failed to convert %s to int: %s", e.key, err)
	}

	return value
}

// ToUint converts the value of the environment variable to an unsigned integer value.
func (e *EnvironmentVariable) ToUint() (uint, error) {
	value, err := e.ToString()
	if err != nil {
		return 0, err
	}

	parsedValue, err := strconv.ParseUint(value, 10, 64)

	return uint(parsedValue), err
}

// MustToUint is like `ToUint`, but it panics if the value is not found or there's an error in the conversion.
func (e *EnvironmentVariable) MustToUint() uint {
	value, err := strconv.ParseUint(e.MustToString(), 10, 64)
	if err != nil {
		log.Panicf("failed to parse %s To uint: %s", e.key, err)
	}

	return uint(value)
}

// ToFloat32 converts the value of the environment variable to a float32 value.
func (e *EnvironmentVariable) ToFloat32() (float32, error) {
	value, err := e.ToString()
	if err != nil {
		return 0, err
	}

	parsedValue, err := strconv.ParseFloat(value, 32)

	return float32(parsedValue), err
}

// MustToFloat32 is like `ToFloat32`, but it panics if the value is not found or there's an error in the conversion.
func (e *EnvironmentVariable) MustToFloat32() float32 {
	value, err := strconv.ParseFloat(e.MustToString(), 32)
	if err != nil {
		log.Panicf("failed to parse %s To float32: %s", e.key, err)
	}

	return float32(value)
}

// ToFloat64 converts the value of the environment variable to a float64 value.
func (e *EnvironmentVariable) ToFloat64() (float64, error) {
	value, err := e.ToString()
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(value, 64)
}

// MustToFloat64 is like `ToFloat64`, but it panics if the value is not found or there's an error in the conversion.
func (e *EnvironmentVariable) MustToFloat64() float64 {
	value, err := strconv.ParseFloat(e.MustToString(), 64)
	if err != nil {
		log.Panicf("failed to parse %s To float64: %s", e.key, err)
	}

	return value
}

// ToStringSlice converts the value of the environment variable to a slice of strings.
func (e *EnvironmentVariable) ToStringSlice(separator string) ([]string, error) {
	strVal, err := e.ToString()
	if err != nil {
		return nil, err
	}

	return strings.Split(strVal, separator), nil
}

// MustToStringSlice is like `ToStringSlice`, but it panics if the value is not found or there's an error in the conversion.
func (e *EnvironmentVariable) MustToStringSlice(separator string) []string {
	strVal := e.MustToString()
	return strings.Split(strVal, separator)
}
