package utils_test

import (
	"testing"

	"github.com/somnek/go-utils/utils"
	"github.com/stretchr/testify/assert"
)

func TestSliceConcat(t *testing.T) {
	t.Parallel()

	listA := []string{"a", "b", "c"}
	listB := [][]string{
		{"d"},
		{},
		{"g", "h"},
	}

	out := []string{"a", "b", "c", "d", "g", "h"}

	result := utils.SliceConcat(listA, listB...)
	assert.Equal(t, out, result)
}
