package config

import (
	"fmt"
	"regexp"
	"strings"
)

func ValidateStringRegexpMatch(exp string) func(value interface{}) error {
	return func(value interface{}) error {
		v := value.(string)
		ok, err := regexp.MatchString(exp, v)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("value '%s' does not match regular expression '%s'", v, exp)
		}
		return nil
	}
}

func ValidateStringContains(s string) func(value interface{}) error {
	return func(value interface{}) error {
		v := value.(string)
		if ok := strings.Contains(v, s); !ok {
			return fmt.Errorf("value '%s' does not contain '%s'", v, s)
		}
		return nil
	}
}

func ValidateStringHasPrefix(s string) func(value interface{}) error {
	return func(value interface{}) error {
		v := value.(string)
		if ok := strings.HasPrefix(v, s); !ok {
			return fmt.Errorf("value '%s' does not start with '%s'", v, s)
		}
		return nil
	}
}

func ValidateStringHasSuffix(s string) func(value interface{}) error {
	return func(value interface{}) error {
		v := value.(string)
		if ok := strings.HasSuffix(v, s); !ok {
			return fmt.Errorf("value '%s' does not end with '%s'", v, s)
		}
		return nil
	}
}
