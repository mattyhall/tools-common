package envvar

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetIntEnvVar(t *testing.T) {
	type test struct {
		name         string
		expectedVal  int
		expectedBool bool
		envName      string
		envValue     string
	}

	tests := []test{
		{
			name:         "ValidEnv",
			expectedVal:  10,
			expectedBool: true,
			envName:      "CB_TEST_ENVAR_VALID_CASE",
			envValue:     "10",
		},
		{
			name:    "EnvNotSet",
			envName: "CB_TEST_ENVAR_NO_ENV_CASE",
		},
		{
			name:     "EnvIsNotAnInt",
			envName:  "CB_TEST_ENVAR_NOT_AN_INT_CASE",
			envValue: "this is not an int",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := os.Setenv(test.envName, test.envValue)
			require.Nil(t, err)

			defer os.Unsetenv(test.envName)

			val, ok := GetInt(test.envName)

			require.Equal(t, test.expectedBool, ok)
			require.Equal(t, test.expectedVal, val)
		})
	}
}

func TestGetUint64EnvVar(t *testing.T) {
	type test struct {
		name         string
		expectedVal  uint64
		expectedBool bool
		envName      string
		envValue     string
	}

	tests := []test{
		{
			name:         "ValidEnv",
			expectedVal:  10,
			expectedBool: true,
			envName:      "CB_TEST_ENVAR_VALID_CASE",
			envValue:     "10",
		},
		{
			name:    "EnvNotSet",
			envName: "CB_TEST_ENVAR_NO_ENV_CASE",
		},
		{
			name:     "EnvIsNotAUint64",
			envName:  "CB_TEST_ENVAR_NOT_AN_INT_CASE",
			envValue: "this is not an int",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := os.Setenv(test.envName, test.envValue)
			require.Nil(t, err)

			defer os.Unsetenv(test.envName)

			val, ok := GetUint64(test.envName)

			require.Equal(t, test.expectedBool, ok)
			require.Equal(t, test.expectedVal, val)
		})
	}
}

func TestGetBoolEnvVar(t *testing.T) {
	type test struct {
		name          string
		value         string
		envName       string
		expectedBool  bool
		expectedValue bool
	}

	tests := []test{
		{
			name:    "NotSet",
			envName: "CB_TEST_ABCDEFG",
		},
		{
			name:          "SetToTrue",
			value:         "true",
			envName:       "CB_TEST_ABCDEFG",
			expectedBool:  true,
			expectedValue: true,
		},
		{
			name:         "SetToFalse",
			value:        "false",
			envName:      "CB_TEST_ABCDEFG",
			expectedBool: true,
		},
		{
			name:          "SetToTruey",
			value:         "t",
			envName:       "CB_TEST_ABCDEFG",
			expectedBool:  true,
			expectedValue: true,
		},
		{
			name:         "SetToFalsy",
			value:        "f",
			envName:      "CB_TEST_ABCDEFG",
			expectedBool: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := os.Setenv(test.envName, test.value)
			require.Nil(t, err)

			out, ok := GetBool(test.envName)

			require.Equal(t, test.expectedBool, ok)
			require.Equal(t, test.expectedValue, out)
		})
	}
}

func TestGetDurationEnvVar(t *testing.T) {
	type test struct {
		name             string
		value            string
		envName          string
		expectedBool     bool
		expectedDuration time.Duration
	}

	tests := []test{
		{
			name:             "SetToSeconds",
			value:            "1s",
			envName:          "CB_TEST_ABCDEFG",
			expectedBool:     true,
			expectedDuration: time.Second,
		},
		{
			name:             "SetToSecondsAndMinutes",
			value:            "33m1s",
			envName:          "CB_TEST_ABCDEFG",
			expectedBool:     true,
			expectedDuration: 33*time.Minute + 1*time.Second,
		},
		{
			name:             "UnsetEnv",
			value:            "",
			envName:          "CB_TEST_ABCDEFG_amdkasd",
			expectedBool:     false,
			expectedDuration: 0,
		},
		{
			name:             "InvalidTimeString",
			value:            "7.87.232.498",
			envName:          "CB_TEST_ABCDEFG",
			expectedBool:     false,
			expectedDuration: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := os.Setenv(test.envName, test.value)
			require.Nil(t, err)

			out, ok := GetDuration(test.envName)

			require.Equal(t, test.expectedBool, ok)
			require.Equal(t, test.expectedDuration, out)
		})
	}
}
