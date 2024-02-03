package version

import (
	"fmt"
	"strconv"
	"strings"
)

// Provider describes version provider.
type Provider struct {
	getter Getter
}

// NewProvider returns new Provider.
func NewProvider(getter Getter) *Provider {
	return &Provider{
		getter: getter,
	}
}

// Version returns Version.
func (p Provider) Version() (Version, error) {
	src := p.getter.Get()
	ver := strings.Split(strings.ReplaceAll(src, "go", ""), "-")[0]

	parts := strings.Split(ver, ".")
	if ver == "" || len(parts) == 0 {
		return Version{}, fmt.Errorf("error: %w %q", ErrUnexpected, src)
	}

	var (
		major, minor, patch int
		err                 error
	)

	if major, err = strconv.Atoi(parts[0]); err != nil {
		return Version{}, fmt.Errorf(
			"error: %w, can't parse major version %q: %w", ErrUnexpected, parts[0], err,
		)
	}

	if len(parts) > 1 && parts[1] != "0" {
		if minor, err = strconv.Atoi(parts[1]); err != nil {
			return Version{}, fmt.Errorf(
				"error: %w, can't parse minor version %q: %w", ErrUnexpected, parts[1], err,
			)
		}
	}

	if len(parts) > 2 && parts[2] != "0" {
		if patch, err = strconv.Atoi(parts[2]); err != nil {
			return Version{}, fmt.Errorf(
				"error: %w, can't parse patch version %q: %w", ErrUnexpected, parts[2], err,
			)
		}
	}

	return Version{
		Major: p.intPtr(major),
		Minor: p.intPtr(minor),
		Patch: p.intPtr(patch),
	}, nil
}

func (p Provider) intPtr(v int) *int {
	if v > 0 {
		return &v
	}

	return nil
}
