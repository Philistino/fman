//go:build windows
// +build windows

package entry

import (
	"errors"
	"strings"
)

var (
	errInvalidFilenameEmptyWindows        = errors.New("name is invalid, must not be empty")
	errInvalidFilenameWindowsSpacePeriod  = errors.New("name is invalid, must not end in space or period on Windows")
	errInvalidFilenameWindowsReservedName = errors.New("name is invalid, contains Windows reserved name (NUL, COM1, etc.)")
	errInvalidFilenameWindowsReservedChar = errors.New(`name is invalid, contains Windows reserved character (<>:"|?*\/)`)
)

const windowsDisallowedCharacters = (`<>:"|?*\/` +
	"\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f")

func InvalidFilename(name string) error {
	if name == "" {
		return errInvalidFilenameEmptyWindows
	}
	// The path must not contain any disallowed characters.
	if strings.ContainsAny(name, windowsDisallowedCharacters) {
		return errInvalidFilenameWindowsReservedChar
	}

	// The name should not end in space or period
	switch name[len(name)-1] {
	case ' ', '.':
		return errInvalidFilenameWindowsSpacePeriod
	}

	// or be a reserved name.
	if windowsIsReserved(name) {
		return errInvalidFilenameWindowsReservedName
	}
	return nil
}

func windowsIsReserved(part string) bool {
	// nul.txt.jpg is also disallowed.
	dot := strings.IndexByte(part, '.')
	if dot != -1 {
		part = part[:dot]
	}

	// Check length to skip allocating ToUpper.
	if len(part) != 3 && len(part) != 4 {
		return false
	}

	// COM0 and LPT0 are missing from the Microsoft docs,
	// but Windows Explorer treats them as invalid too.
	// (https://docs.microsoft.com/windows/win32/fileio/naming-a-file)
	switch strings.ToUpper(part) {
	case "CON", "PRN", "AUX", "NUL",
		"COM0", "COM1", "COM2", "COM3", "COM4",
		"COM5", "COM6", "COM7", "COM8", "COM9",
		"LPT0", "LPT1", "LPT2", "LPT3", "LPT4",
		"LPT5", "LPT6", "LPT7", "LPT8", "LPT9":
		return true
	}
	return false
}
