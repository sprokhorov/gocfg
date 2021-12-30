package config

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateStringFuncs(t *testing.T) {
	testcases := []struct {
		name  string
		value string
		vfunc func(value interface{}) error
		err   error
	}{
		{"check regexp func", "https://api.example.com", ValidateStringRegexpMatch(`^http(s)?\:\/\/.*$`), nil},
		{"check regexp func failed", "https://api.example.com", ValidateStringRegexpMatch(`^ftp`), errors.New("value 'https://api.example.com' does not match regular expression '^ftp'")},
		{"check contains func", "https://api.example.com", ValidateStringContains("example"), nil},
		{"check contains func failed", "https://api.example.com", ValidateStringContains("google"), errors.New("value 'https://api.example.com' does not contain 'google'")},
		{"check has prefix func", "https://api.example.com", ValidateStringHasPrefix("https"), nil},
		{"check has prefix func failed", "https://api.example.com", ValidateStringHasPrefix("ssh"), errors.New("value 'https://api.example.com' does not start with 'ssh'")},
		{"check has suffix func", "https://api.example.com", ValidateStringHasSuffix("com"), nil},
		{"check has suffix func failed", "https://api.example.com", ValidateStringHasSuffix("org"), errors.New("value 'https://api.example.com' does not end with 'org'")},
	}
	for _, tc := range testcases {
		err := tc.vfunc(tc.value)
		assert.Equal(t, err, tc.err, fmt.Sprintf("Test case: %s", tc.name))
	}
}
