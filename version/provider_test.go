package version_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/alexeyco/devtools/pointer"
	"github.com/alexeyco/devtools/version"
)

func TestProvider_Version(t *testing.T) {
	testData := [...]struct {
		name     string
		source   string
		expected version.Version
		err      error
	}{
		{
			name:   "Ok/Major",
			source: "go1",
			expected: version.Version{
				Major: pointer.Ptr(1),
			},
		},
		{
			name:   "Ok/MajorMinor",
			source: "go1.2",
			expected: version.Version{
				Major: pointer.Ptr(1),
				Minor: pointer.Ptr(2),
			},
		},
		{
			name:   "Ok/MajorMinorPatch",
			source: "go1.2.3",
			expected: version.Version{
				Major: pointer.Ptr(1),
				Minor: pointer.Ptr(2),
				Patch: pointer.Ptr(3),
			},
		},
		{
			name:   "Ok/MajorMinorPatchPreRelease",
			source: "go1.2.3-foobar",
			expected: version.Version{
				Major: pointer.Ptr(1),
				Minor: pointer.Ptr(2),
				Patch: pointer.Ptr(3),
			},
		},
		{
			name:   "Err/Empty",
			source: "go",
			err:    version.ErrUnexpected,
		},
		{
			name:   "Err/Major",
			source: "goA",
			err:    version.ErrUnexpected,
		},
		{
			name:   "Err/Minor",
			source: "go1.A",
			err:    version.ErrUnexpected,
		},
		{
			name:   "Err/Patch",
			source: "go1.2.A",
			err:    version.ErrUnexpected,
		},
	}

	ctrl := gomock.NewController(t)
	getter := NewMockGetter(ctrl)
	p := version.NewProvider(getter)

	for _, datum := range testData {
		datum := datum

		t.Run(datum.name, func(t *testing.T) {
			getter.EXPECT().Get().Return(datum.source)

			actual, err := p.Version()

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
