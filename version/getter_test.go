package version_test

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/devtools/version"
)

func TestDefaultGetter_Get(t *testing.T) {
	expected := runtime.Version()
	actual := version.NewDefaultGetter().Get()

	assert.Equal(t, expected, actual)
}
