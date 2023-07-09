//go:build !windows
// +build !windows

package entry

import (
	"errors"
	"strings"
)

var (
	errInvalidFilenameEmptyUnix        = errors.New("name is invalid, must not be empty")
	errInvalidFilenameReservedCharUnix = errors.New(`name is invalid, contains Windows reserved character (<>:"|?*\/)`)
)

const unixDisallowedCharacters = `/`

func InvalidFilename(name string) error {
	if name == "" {
		return errInvalidFilenameEmptyUnix
	}
	// The path must not contain any disallowed characters.
	if strings.ContainsAny(name, unixDisallowedCharacters) {
		return errInvalidFilenameReservedCharUnix
	}
	return nil
}
