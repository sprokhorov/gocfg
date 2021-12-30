package gocfg

import (
	"fmt"
	"strconv"
)

// parseString lookups for the environment variable and assign it
// to the pointer. It return error if:
// - variable was not defined but it's required
// - default value has a wrong type
// - variable values has a wrong type
func (c *Config) parseString(setting *Variable) error {
	v, ok := c.env.LookupEnv(setting.Name)
	if setting.Required && !ok {
		return fmt.Errorf("'%s' variable is missing", setting.Name)
	}
	p := setting.pointer.(*string)
	// set a default value if env var was not defined
	if !ok && setting.Default != nil {
		d, ok := setting.Default.(string)
		if !ok {
			return fmt.Errorf("variable '%s' has a wrong default value type", setting.Name)
		}
		*p = d
		return nil
	}
	*p = v
	// validate value
	if setting.ValidationFunc != nil {
		return setting.ValidationFunc(v)
	}
	return nil
}

// parseInt lookups for the environment variable and assign it
// to the pointer. It return error if:
// - variable was not defined but it's required
// - default value has a wrong type
// - variable values has a wrong type
func (c *Config) parseInt(setting *Variable) error {
	v, ok := c.env.LookupEnv(setting.Name)
	if setting.Required && !ok {
		return fmt.Errorf("'%s' variable is missing", setting.Name)
	}
	p := setting.pointer.(*int)
	// set a default value if env var was not defined
	if !ok && setting.Default != nil {
		d, ok := setting.Default.(int)
		if !ok {
			return fmt.Errorf("variable '%s' has a wrong default value type", setting.Name)
		}
		*p = d
		return nil
	}
	iv, err := strconv.Atoi(v)
	if err != nil {
		return fmt.Errorf("variable '%s' has a wrong value type", setting.Name)
	}
	*p = iv
	// validate value
	if setting.ValidationFunc != nil {
		return setting.ValidationFunc(iv)
	}
	return nil
}

// parseInt64 lookups for the environment variable and assign it
// to the pointer. It return error if:
// - variable was not defined but it's required
// - default value has a wrong type
// - variable values has a wrong type
func (c *Config) parseInt64(setting *Variable) error {
	v, ok := c.env.LookupEnv(setting.Name)
	if setting.Required && !ok {
		return fmt.Errorf("'%s' variable is missing", setting.Name)
	}
	p := setting.pointer.(*int64)
	// set a default value if env var was not defined
	if !ok && setting.Default != nil {
		d, ok := setting.Default.(int64)
		if !ok {
			return fmt.Errorf("variable '%s' has a wrong default value type", setting.Name)
		}
		*p = d
		return nil
	}
	iv, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return fmt.Errorf("variable '%s' has a wrong value type", setting.Name)
	}
	*p = iv
	// validate value
	if setting.ValidationFunc != nil {
		return setting.ValidationFunc(iv)
	}
	return nil
}

// parseFloat32 lookups for the environment variable and assign it
// to the pointer. It return error if:
// - variable was not defined but it's required
// - default value has a wrong type
// - variable values has a wrong type
func (c *Config) parseFloat32(setting *Variable) error {
	v, ok := c.env.LookupEnv(setting.Name)
	if setting.Required && !ok {
		return fmt.Errorf("'%s' variable is missing", setting.Name)
	}
	p := setting.pointer.(*float32)
	// set a default value if env var was not defined
	if !ok && setting.Default != nil {
		d, ok := setting.Default.(float32)
		if !ok {
			return fmt.Errorf("variable '%s' has a wrong default value type", setting.Name)
		}
		*p = d
		return nil
	}
	fv, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return fmt.Errorf("variable '%s' has a wrong value type", setting.Name)
	}
	*p = float32(fv)
	// validate value
	if setting.ValidationFunc != nil {
		return setting.ValidationFunc(fv)
	}
	return nil
}

// parseFloat64 lookups for the environment variable and assign it
// to the pointer. It return error if:
// - variable was not defined but it's required
// - default value has a wrong type
// - variable values has a wrong type
func (c *Config) parseFloat64(setting *Variable) error {
	v, ok := c.env.LookupEnv(setting.Name)
	if setting.Required && !ok {
		return fmt.Errorf("'%s' variable is missing", setting.Name)
	}
	p := setting.pointer.(*float64)
	// set a default value if env var was not defined
	if !ok && setting.Default != nil {
		d, ok := setting.Default.(float64)
		if !ok {
			return fmt.Errorf("variable '%s' has a wrong default value type", setting.Name)
		}
		*p = d
		return nil
	}
	fv, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return fmt.Errorf("variable '%s' has a wrong value type", setting.Name)
	}
	*p = fv
	// validate value
	if setting.ValidationFunc != nil {
		return setting.ValidationFunc(fv)
	}
	return nil
}

// parseBool lookups for the environment variable and assign it
// to the pointer. It return error if:
// - variable was not defined but it's required
// - default value has a wrong type
// - variable values has a wrong type
func (c *Config) parseBool(setting *Variable) error {
	v, ok := c.env.LookupEnv(setting.Name)
	if setting.Required && !ok {
		return fmt.Errorf("'%s' variable is missing", setting.Name)
	}
	p := setting.pointer.(*bool)
	// set a default value if env var was not defined
	if !ok && setting.Default != nil {
		d, ok := setting.Default.(bool)
		if !ok {
			return fmt.Errorf("variable '%s' has a wrong default value type", setting.Name)
		}
		*p = d
		return nil
	}
	bv, err := strconv.ParseBool(v)
	if err != nil {
		return fmt.Errorf("variable '%s' has a wrong value type", setting.Name)
	}
	*p = bv
	// validate value
	if setting.ValidationFunc != nil {
		return setting.ValidationFunc(bv)
	}
	return nil
}

// Parse lookups for defined environment variables and asserts their types. It
// collects all parsing errors in single slice and return it.
//
// It return error if:
// - variable was not defined but it's required
// - default value has a wrong type
// - variable values has a wrong type
func (c *Config) Parse() error {
	errs := NewParseErrors()
	for _, v := range c.variables {
		switch v.valueType {
		case STRING:
			if err := c.parseString(v); err != nil {
				errs.Add(err)
				continue
			}
		case INT:
			if err := c.parseInt(v); err != nil {
				errs.Add(err)
				continue
			}
		case INT64:
			if err := c.parseInt64(v); err != nil {
				errs.Add(err)
				continue
			}
		case FLOAT32:
			if err := c.parseFloat32(v); err != nil {
				errs.Add(err)
				continue
			}
		case FLOAT64:
			if err := c.parseFloat64(v); err != nil {
				errs.Add(err)
				continue
			}
		case BOOL:
			if err := c.parseBool(v); err != nil {
				errs.Add(err)
				continue
			}
		}
	}
	if errs.IsNotNil() {
		return errs
	}
	return nil
}

// ParseErrors implements error interface but holds multiple errors
// instead of one.
type ParseErrors struct {
	errs []error
}

// NewParseErrors returns a new ParseErrors struct.
func NewParseErrors(errs ...error) *ParseErrors {
	return &ParseErrors{errs: append([]error{}, errs...)}
}

// Error implements Error method of error interface.
func (pe *ParseErrors) Error() string {
	m := "config parsing failed: "
	for i, e := range pe.errs {
		m += e.Error()
		if i != len(pe.errs)-1 {
			m += ", "
		}
	}
	return m
}

// Add adds error to the list.
func (pe *ParseErrors) Add(err error) {
	pe.errs = append(pe.errs, err)
}

// IsNotNil returns true if struct has errors.
func (pe *ParseErrors) IsNotNil() bool {
	return len(pe.errs) > 0
}
