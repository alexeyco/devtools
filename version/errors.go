package version

import "errors"

var (
	// ErrEmpty empty version.
	ErrEmpty = errors.New("empty version")

	// ErrUnexpected unexpected version.
	ErrUnexpected = errors.New("unexpected version")
)
