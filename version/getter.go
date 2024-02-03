package version

import "runtime"

// Getter describes go version getter interface.
type Getter interface {
	Get() string
}

// DefaultGetter describes default go version getter.
type DefaultGetter struct{}

// NewDefaultGetter returns new DefaultGetter.
func NewDefaultGetter() *DefaultGetter {
	return &DefaultGetter{}
}

// Get returns go version.
func (DefaultGetter) Get() string {
	return runtime.Version()
}
