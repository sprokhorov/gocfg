package config

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// EnvLookuperMock implements EnvLookuper.
type EnvLookuperMock struct {
	vars map[string]string
}

// LookupEnv lookups environment variables from internal map.
func (elm *EnvLookuperMock) LookupEnv(key string) (string, bool) {
	v, b := elm.vars[key]
	return v, b
}

func TestConfigString(t *testing.T) {
	env := &EnvLookuperMock{
		vars: map[string]string{
			"LOG_LEVEL":   "DEBUG",
			"API_URL":     "https://api.example.com",
			"API_VERSION": "/v5",
		},
	}

	testcases := []struct {
		name     string
		variable *Variable
		err      error
		value    string
	}{
		{"basic test case", &Variable{Name: "LOG_LEVEL"}, nil, "DEBUG"},
		{"default value", &Variable{Name: "LOG_FORMAT", Default: "JSON"}, nil, "JSON"},
		{"wrong default value type", &Variable{Name: "LOG_FORMAT", Default: 1}, NewParseErrors(errors.New("variable 'LOG_FORMAT' has a wrong default value type")), ""},
		{"default value conflict", &Variable{Name: "LOG_LEVEL", Default: "WARN"}, nil, "DEBUG"},
		{"missing required var", &Variable{Name: "LOGG_LVL", Required: true}, NewParseErrors(errors.New("'LOGG_LVL' variable is missing")), ""},
		{"unformated variable name", &Variable{Name: "api-version"}, nil, "/v5"},
	}

	for _, tc := range testcases {
		cfg := New()
		cfg.SetEnvLookuper(env)
		var value string
		cfg.SetString(&value, tc.variable)
		err := cfg.Parse()
		assert.Equal(t, err, tc.err, fmt.Sprintf("Test case: %s, Error: '%+v'", tc.name, err))
		assert.Equal(t, value, tc.value, fmt.Sprintf("Test case: %s", tc.name))
	}
}

func TestConfigInt(t *testing.T) {
	env := &EnvLookuperMock{
		vars: map[string]string{
			"REQUEST_TIMEOUT": "30",
			"API_URL":         "https://api.example.com",
			"API_VERSION":     "5",
		},
	}

	testcases := []struct {
		name     string
		variable *Variable
		err      error
		value    int
	}{
		{"basic test case", &Variable{Name: "REQUEST_TIMEOUT"}, nil, 30},
		{"default value", &Variable{Name: "BATCH_SIZE", Default: 500}, nil, 500},
		{"wrong default value type", &Variable{Name: "BATCH_SIZE", Default: "1"}, NewParseErrors(errors.New("variable 'BATCH_SIZE' has a wrong default value type")), 0},
		{"default value conflict", &Variable{Name: "REQUEST_TIMEOUT", Default: 5}, nil, 30},
		{"missing required var", &Variable{Name: "BATCH_SIZE", Required: true}, NewParseErrors(errors.New("'BATCH_SIZE' variable is missing")), 0},
		{"unformated variable name", &Variable{Name: "api-version"}, nil, 5},
	}

	for _, tc := range testcases {
		cfg := New()
		cfg.SetEnvLookuper(env)
		var value int
		cfg.SetInt(&value, tc.variable)
		err := cfg.Parse()
		assert.Equal(t, err, tc.err, fmt.Sprintf("Test case: %s, Error: '%+v'", tc.name, err))
		assert.Equal(t, value, tc.value, fmt.Sprintf("Test case: %s", tc.name))
	}
}

func TestConfigInt64(t *testing.T) {
	env := &EnvLookuperMock{
		vars: map[string]string{
			"REQUEST_TIMEOUT": "30",
			"API_VERSION":     "5",
		},
	}

	testcases := []struct {
		name     string
		variable *Variable
		err      error
		value    int64
	}{
		{"basic test case", &Variable{Name: "REQUEST_TIMEOUT"}, nil, 30},
		{"default value", &Variable{Name: "BATCH_SIZE", Default: int64(500)}, nil, 500},
		{"wrong default value type", &Variable{Name: "BATCH_SIZE", Default: "1"}, NewParseErrors(errors.New("variable 'BATCH_SIZE' has a wrong default value type")), 0},
		{"default value conflict", &Variable{Name: "REQUEST_TIMEOUT", Default: 5}, nil, 30},
		{"missing required var", &Variable{Name: "BATCH_SIZE", Required: true}, NewParseErrors(errors.New("'BATCH_SIZE' variable is missing")), 0},
		{"unformated variable name", &Variable{Name: "api-version"}, nil, 5},
	}

	for _, tc := range testcases {
		cfg := New()
		cfg.SetEnvLookuper(env)
		var value int64
		cfg.SetInt64(&value, tc.variable)
		err := cfg.Parse()
		assert.Equal(t, err, tc.err, fmt.Sprintf("Test case: %s, Error: '%+v'", tc.name, err))
		assert.Equal(t, value, tc.value, fmt.Sprintf("Test case: %s", tc.name))
	}
}

func TestConfigFloat32(t *testing.T) {
	env := &EnvLookuperMock{
		vars: map[string]string{
			"CPU_REQUEST":            "1.2",
			"LOAD_AVERAGE_THRESHOLD": "0.8",
		},
	}

	testcases := []struct {
		name     string
		variable *Variable
		err      error
		value    float32
	}{
		{"basic test case", &Variable{Name: "CPU_REQUEST"}, nil, 1.2},
		{"default value", &Variable{Name: "CPU_LIMIT", Default: float32(1.8)}, nil, 1.8},
		{"wrong default value type", &Variable{Name: "CPU_LIMIT", Default: "1"}, NewParseErrors(errors.New("variable 'CPU_LIMIT' has a wrong default value type")), 0},
		{"default value conflict", &Variable{Name: "LOAD_AVERAGE_THRESHOLD", Default: 2.2}, nil, 0.8},
		{"missing required var", &Variable{Name: "CPU_LIMIT", Required: true}, NewParseErrors(errors.New("'CPU_LIMIT' variable is missing")), 0},
		{"unformated variable name", &Variable{Name: "cpu-request"}, nil, 1.2},
	}

	for _, tc := range testcases {
		cfg := New()
		cfg.SetEnvLookuper(env)
		var value float32
		cfg.SetFloat32(&value, tc.variable)
		err := cfg.Parse()
		assert.Equal(t, err, tc.err, fmt.Sprintf("Test case: %s, Error: '%+v'", tc.name, err))
		assert.Equal(t, value, tc.value, fmt.Sprintf("Test case: %s", tc.name))
	}
}

func TestConfigFloat64(t *testing.T) {
	env := &EnvLookuperMock{
		vars: map[string]string{
			"CPU_REQUEST":            "1.2",
			"LOAD_AVERAGE_THRESHOLD": "0.8",
		},
	}

	testcases := []struct {
		name     string
		variable *Variable
		err      error
		value    float64
	}{
		{"basic test case", &Variable{Name: "CPU_REQUEST"}, nil, 1.2},
		{"default value", &Variable{Name: "CPU_LIMIT", Default: float64(1.8)}, nil, 1.8},
		{"wrong default value type", &Variable{Name: "CPU_LIMIT", Default: "1"}, NewParseErrors(errors.New("variable 'CPU_LIMIT' has a wrong default value type")), 0},
		{"default value conflict", &Variable{Name: "LOAD_AVERAGE_THRESHOLD", Default: 2.2}, nil, 0.8},
		{"missing required var", &Variable{Name: "CPU_LIMIT", Required: true}, NewParseErrors(errors.New("'CPU_LIMIT' variable is missing")), 0},
		{"unformated variable name", &Variable{Name: "cpu-request"}, nil, 1.2},
	}

	for _, tc := range testcases {
		cfg := New()
		cfg.SetEnvLookuper(env)
		var value float64
		cfg.SetFloat64(&value, tc.variable)
		err := cfg.Parse()
		assert.Equal(t, err, tc.err, fmt.Sprintf("Test case: %s, Error: '%+v'", tc.name, err))
		assert.Equal(t, value, tc.value, fmt.Sprintf("Test case: %s", tc.name))
	}
}

func TestConfigBool(t *testing.T) {
	env := &EnvLookuperMock{
		vars: map[string]string{
			"TRACING_ENABLED": "true",
			"USE_GCS_STORAGE": "false",
		},
	}

	testcases := []struct {
		name     string
		variable *Variable
		err      error
		value    bool
	}{
		{"basic test case", &Variable{Name: "TRACING_ENABLED"}, nil, true},
		{"default value", &Variable{Name: "USE_S3_STORAGE", Default: true}, nil, true},
		{"wrong default value type", &Variable{Name: "USE_S3_STORAGE", Default: "NO"}, NewParseErrors(errors.New("variable 'USE_S3_STORAGE' has a wrong default value type")), false},
		{"default value conflict", &Variable{Name: "USE_GCS_STORAGE", Default: true}, nil, false},
		{"missing required var", &Variable{Name: "USE_S3_STORAGE", Required: true}, NewParseErrors(errors.New("'USE_S3_STORAGE' variable is missing")), false},
		{"unformated variable name", &Variable{Name: "tracing-enabled"}, nil, true},
	}

	for _, tc := range testcases {
		cfg := New()
		cfg.SetEnvLookuper(env)
		var value bool
		cfg.SetBool(&value, tc.variable)
		err := cfg.Parse()
		assert.Equal(t, err, tc.err, fmt.Sprintf("Test case: %s, Error: '%+v'", tc.name, err))
		assert.Equal(t, value, tc.value, fmt.Sprintf("Test case: %s", tc.name))
	}
}
