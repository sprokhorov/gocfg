package config

import (
	"os"
	"strings"
)

// valueType represents valid config types
type valueType byte

// supported Variable types
const (
	STRING valueType = iota
	INT
	INT64
	FLOAT32
	FLOAT64
	BOOL
)

// EnvLookuper represents app environment variables. It helps
// to avoid manipulations with real environment and mock it in tests.
type EnvLookuper interface {
	LookupEnv(key string) (string, bool)
}

// EnvLookuperImpl implements EnvLookuper interface.
type EnvLookuperImpl struct{}

// LookupEnv wraps os.LookupEnv() function.
func (eli *EnvLookuperImpl) LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

// Variable represents environment variable and rules of it's validation.
type Variable struct {
	Default        interface{}
	Name           string
	Required       bool
	ValidationFunc func(value interface{}) error
	pointer        interface{}
	valueType      valueType
}

// Config manages variables lookup and validation.
type Config struct {
	variables []*Variable
	env       EnvLookuper
}

// New returns new Config object.
func New() *Config {
	return &Config{env: &EnvLookuperImpl{}}
}

// SetEnvLookuper sets default EnvLookuper. Check tests for examples.
func (c *Config) SetEnvLookuper(l EnvLookuper) {
	c.env = l
}

// setVariable adds variable to config.
func (c *Config) setVariable(setting *Variable) {
	formatEnvVarName(&setting.Name)
	c.variables = append(c.variables, setting)
}

// SetString adds variable to config. It requiers a pointer to the
// go variable to assign value after parsing.
func (c *Config) SetString(pointer *string, setting *Variable) {
	setting.valueType = STRING
	setting.pointer = pointer
	c.setVariable(setting)
}

// SetInt adds variable to config. It requiers a pointer to the
// go variable to assign value after parsing.
func (c *Config) SetInt(pointer *int, setting *Variable) {
	setting.valueType = INT
	setting.pointer = pointer
	c.setVariable(setting)
}

// SetInt64 adds variable to config. It requiers a pointer to the
// go variable to assign value after parsing.
func (c *Config) SetInt64(pointer *int64, setting *Variable) {
	setting.valueType = INT64
	setting.pointer = pointer
	c.setVariable(setting)
}

// SetFloat32 adds variable to config. It requiers a pointer to the
// go variable to assign value after parsing.
func (c *Config) SetFloat32(pointer *float32, setting *Variable) {
	setting.valueType = FLOAT32
	setting.pointer = pointer
	c.setVariable(setting)
}

// SetFloat64 adds variable to config. It requiers a pointer to the
// go variable to assign value after parsing.
func (c *Config) SetFloat64(pointer *float64, setting *Variable) {
	setting.valueType = FLOAT64
	setting.pointer = pointer
	c.setVariable(setting)
}

// SetBool adds variable to config. It requiers a pointer to the
// go variable to assign value after parsing.
func (c *Config) SetBool(pointer *bool, setting *Variable) {
	setting.valueType = BOOL
	setting.pointer = pointer
	c.setVariable(setting)
}

func formatEnvVarName(name *string) {
	*name = strings.ReplaceAll(*name, "-", "_")
	*name = strings.ToUpper(*name)
}
