package pointer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/devtools/pointer"
)

func TestPointer(t *testing.T) {
	testData := [...]struct {
		name     string
		expected any
	}{
		{
			name:     "String",
			expected: "var",
		},
		{
			name:     "Int",
			expected: 123,
		},
		{
			name:     "Bool",
			expected: true,
		},
	}

	for _, datum := range testData {
		datum := datum

		t.Run(datum.name, func(t *testing.T) {
			actual := pointer.Val(pointer.Ptr(datum.expected))

			assert.Equal(t, datum.expected, actual)
		})
	}
}
