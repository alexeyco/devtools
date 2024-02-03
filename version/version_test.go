package version_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/devtools/pointer"
	"github.com/alexeyco/devtools/version"
)

func TestVersion_Tags(t *testing.T) {
	testData := [...]struct {
		name     string
		v        version.Version
		expected []string
		err      error
	}{
		{
			name: "Ok/Major",
			v: version.Version{
				Major: pointer.Ptr(1),
			},
			expected: []string{
				"1",
				"latest",
			},
			err: nil,
		},
		{
			name: "Ok/MajorMinor",
			v: version.Version{
				Major: pointer.Ptr(1),
				Minor: pointer.Ptr(2),
			},
			expected: []string{
				"1",
				"1.2",
				"latest",
			},
			err: nil,
		},
		{
			name: "Ok/MajorMinorPatch",
			v: version.Version{
				Major: pointer.Ptr(1),
				Minor: pointer.Ptr(2),
				Patch: pointer.Ptr(3),
			},
			expected: []string{
				"1",
				"1.2",
				"1.2.3",
				"latest",
			},
			err: nil,
		},
		{
			name: "Err/Empty",
			v:    version.Version{},
			err:  version.ErrEmpty,
		},
	}

	for _, datum := range testData {
		datum := datum

		t.Run(datum.name, func(t *testing.T) {
			actual, err := datum.v.Tags()

			if datum.err != nil {
				assert.ErrorIs(t, err, datum.err)
				assert.Empty(t, actual)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, datum.expected, actual)
		})
	}
}

func TestVersion_String(t *testing.T) {
	testData := [...]struct {
		name     string
		v        version.Version
		expected string
	}{
		{
			name: "Empty",
		},
		{
			name: "Major",
			v: version.Version{
				Major: pointer.Ptr(1),
			},
			expected: "1",
		},
		{
			name: "MajorMinor",
			v: version.Version{
				Major: pointer.Ptr(1),
				Minor: pointer.Ptr(2),
			},
			expected: "1.2",
		},
		{
			name: "MajorMinorPatch",
			v: version.Version{
				Major: pointer.Ptr(1),
				Minor: pointer.Ptr(2),
				Patch: pointer.Ptr(3),
			},
			expected: "1.2.3",
		},
	}

	for _, datum := range testData {
		datum := datum

		t.Run(datum.name, func(t *testing.T) {
			actual := datum.v.String()

			assert.Equal(t, datum.expected, actual)
		})
	}
}
