package version

import (
	"fmt"
	"strconv"
)

// Version describes current go version.
type Version struct {
	Major *int
	Minor *int
	Patch *int
}

// Tags returns image tags list by go version.
func (v Version) Tags() ([]string, error) {
	tags, err := v.tags()
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	return append(tags, "latest"), nil
}

func (v Version) tags() ([]string, error) {
	var t []string

	if v.Major == nil {
		return nil, ErrEmpty
	}

	t = append(t, strconv.Itoa(*v.Major))

	if v.Minor != nil {
		t = append(t, fmt.Sprintf("%d.%d", *v.Major, *v.Minor))
	}

	if v.Patch != nil {
		t = append(t, fmt.Sprintf("%d.%d.%d", *v.Major, *v.Minor, *v.Patch))
	}

	return t, nil
}

// String returns go version as a string.
func (v Version) String() string {
	if v.Major == nil {
		return ""
	}

	if v.Minor == nil {
		return strconv.Itoa(*v.Major)
	}

	if v.Patch == nil {
		return fmt.Sprintf("%d.%d", *v.Major, *v.Minor)
	}

	return fmt.Sprintf("%d.%d.%d", *v.Major, *v.Minor, *v.Patch)
}
