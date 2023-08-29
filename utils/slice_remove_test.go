package utils_test

import (
	"testing"

	"github.com/somnek/go-utils/utils"
	"github.com/stretchr/testify/assert"
)

func TestSliceRemove(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []int
		index    int
		expected []int
	}{
		{
			name:     "remove last",
			input:    []int{1, 2, 3, 4, 5},
			index:    4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "remove from empty slice",
			input:    []int{},
			index:    0,
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := utils.SliceRemove(tc.input, tc.index)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
