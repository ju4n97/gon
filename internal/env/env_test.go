package env_test

import (
	"os"
	"testing"

	"github.com/ju4n97/gon/internal/env"
	"github.com/stretchr/testify/assert"
)

const (
	testKeyString  = "TEST_STRING_KEY"
	testKeyBool    = "TEST_BOOL_KEY"
	testKeyInt     = "TEST_INT_KEY"
	testKeyUint    = "TEST_UINT_KEY"
	testKeyFloat32 = "TEST_FLOAT32_KEY"
	testKeyFloat64 = "TEST_FLOAT64_KEY"
	testKeySlice   = "TEST_SLICE_KEY"

	testKeyStringValue  = "test_string_value"
	testKeyBoolValue    = "true"
	testKeyIntValue     = "97"
	testKeyUintValue    = "97"
	testKeyFloat32Value = "3.14"
	testKeyFloat64Value = "3.14"
	testKeySliceValue   = "1,2,3"
)

const (
	testKeyExpectedStringValue  = "test_string_value"
	testKeyExpectedBoolValue    = true
	testKeyExpectedIntValue     = int(97)
	testKeyExpectedUintValue    = uint(97)
	testKeyExpectedFloat32Value = float32(3.14)
	testKeyExpectedFloat64Value = float64(3.14)
)

var testKeyExpectedSliceValue = []string{"1", "2", "3"}

func TestMain(m *testing.M) {
	// Setup
	os.Setenv(testKeyString, testKeyStringValue)
	os.Setenv(testKeyBool, testKeyBoolValue)
	os.Setenv(testKeyInt, testKeyIntValue)
	os.Setenv(testKeyUint, testKeyUintValue)
	os.Setenv(testKeyFloat32, testKeyFloat32Value)
	os.Setenv(testKeyFloat64, testKeyFloat64Value)
	os.Setenv(testKeySlice, testKeySliceValue)

	// Run tests
	exitCode := m.Run()

	// Teardown
	os.Unsetenv(testKeyString)
	os.Unsetenv(testKeyBool)
	os.Unsetenv(testKeyInt)
	os.Unsetenv(testKeyUint)
	os.Unsetenv(testKeyFloat32)
	os.Unsetenv(testKeyFloat64)
	os.Unsetenv(testKeySlice)

	os.Exit(exitCode)
}

func TestNew(t *testing.T) {
	t.Parallel()

	testEnv := env.New(testKeyString)

	assert.NotNil(t, testEnv)
	_, err := testEnv.ToString()
	assert.NoError(t, err)
}

func TestNewWithoutVariable(t *testing.T) {
	t.Parallel()

	expectedKey := "SOME_NON_EXISTING_KEY"

	testEnv := env.New(expectedKey)

	os.Unsetenv(expectedKey) // Ensure it doesn't exist.

	assert.NotNil(t, testEnv)
	_, err := testEnv.ToString()
	assert.Error(t, err)
}

func TestWithDefault(t *testing.T) {
	t.Parallel()

	expectedKey := "NON_EXISTING_KEY"
	expectedDefaultValue := "test_default_value"

	os.Unsetenv(expectedKey) // Ensure it doesn't exist.

	testEnv := env.New(expectedKey).WithDefault(expectedDefaultValue)

	assert.NotNil(t, testEnv)
	assert.Equal(t, expectedDefaultValue, testEnv.MustToString())
}

func TestValues(t *testing.T) {
	t.Parallel()

	strValue, err := env.New(testKeyString).ToString()
	assert.NoError(t, err)
	assert.Equal(t, testKeyExpectedStringValue, strValue)

	boolValue, err := env.New(testKeyBool).ToBool()
	assert.NoError(t, err)
	assert.Equal(t, testKeyExpectedBoolValue, boolValue)

	intValue, err := env.New(testKeyInt).ToInt()
	assert.NoError(t, err)
	assert.Equal(t, testKeyExpectedIntValue, intValue)

	uintValue, err := env.New(testKeyUint).ToUint()
	assert.NoError(t, err)
	assert.Equal(t, testKeyExpectedUintValue, uintValue)

	float32Value, err := env.New(testKeyFloat32).ToFloat32()
	assert.NoError(t, err)
	assert.Equal(t, testKeyExpectedFloat32Value, float32Value)

	float64Value, err := env.New(testKeyFloat64).ToFloat64()
	assert.NoError(t, err)
	assert.Equal(t, testKeyExpectedFloat64Value, float64Value)

	sliceValue, err := env.New(testKeySlice).ToStringSlice(",")
	assert.NoError(t, err)
	assert.Equal(t, testKeyExpectedSliceValue, sliceValue)

	// Must functions
	assert.Equal(t, testKeyExpectedStringValue, env.New(testKeyString).MustToString())
	assert.Equal(t, testKeyExpectedBoolValue, env.New(testKeyBool).MustToBool())
	assert.Equal(t, testKeyExpectedIntValue, env.New(testKeyInt).MustToInt())
	assert.Equal(t, testKeyExpectedUintValue, env.New(testKeyUint).MustToUint())
	assert.Equal(t, testKeyExpectedFloat32Value, env.New(testKeyFloat32).MustToFloat32())
	assert.Equal(t, testKeyExpectedFloat64Value, env.New(testKeyFloat64).MustToFloat64())
	assert.Equal(t, testKeyExpectedSliceValue, env.New(testKeySlice).MustToStringSlice(","))
}

func TestErrors(t *testing.T) {
	t.Parallel()

	expectedKey := "NON_EXISTING_KEY"

	os.Unsetenv(expectedKey) // Ensure it doesn't exist.

	testEnv := env.New(expectedKey)

	_, err := testEnv.ToBool()
	assert.Error(t, err)

	_, err = testEnv.ToInt()
	assert.Error(t, err)

	_, err = testEnv.ToUint()
	assert.Error(t, err)

	_, err = testEnv.ToFloat32()
	assert.Error(t, err)

	_, err = testEnv.ToFloat64()
	assert.Error(t, err)

	_, err = testEnv.ToStringSlice(",")
	assert.Error(t, err)

	// Must functions
	assert.Panics(t, func() { testEnv.MustToBool() })
	assert.Panics(t, func() { testEnv.MustToInt() })
	assert.Panics(t, func() { testEnv.MustToUint() })
	assert.Panics(t, func() { testEnv.MustToFloat32() })
	assert.Panics(t, func() { testEnv.MustToFloat64() })
	assert.Panics(t, func() { testEnv.MustToStringSlice(",") })
}
