package utils_test

import (
	"testing"

	"github.com/somnek/go-utils/utils"
)

func TestSliceFind(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []int
		value    int
		expected int
	}{
		{
			name:     "value exists in slice",
			input:    []int{1, 2, 3, 4, 5},
			value:    3,
			expected: 2,
		},
		{
			name:     "value does not exist in slice",
			input:    []int{1, 2, 3, 4, 5},
			value:    6,
			expected: -1,
		},
		{
			name:     "empty slice",
			input:    []int{},
			value:    1,
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := utils.SliceFind(tc.input, tc.value)
			if actual != tc.expected {
				t.Errorf("expected %d but got %d", tc.expected, actual)
			}
		})
	}
}
