// Package flag contains flag unmarshalers for the github.com/jessevdk/go-flags
// parsing library.
package flag

import (
	"errors"
	"strings"
)

// Ext represents a file extension being passed in by a user on the
// command-line.
type Ext string

// ErrEmptyExt is returned when the user provides an empty extension.
var ErrEmptyExt = errors.New("extension may not be empty")

// UnmarshalFlag parses the raw string from the command line into a usable
// object. It normalizes the extension to contain a leading dot. An error will
// be returned if an empty extension is passed.
func (e *Ext) UnmarshalFlag(value string) error {
	if value == "" {
		return ErrEmptyExt
	}

	*e = Ext(value)
	if !strings.HasPrefix(value, ".") {
		*e = "." + *e
	}

	return nil
}
