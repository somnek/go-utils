package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/somnek/go-utils/utils"
)

func TestSliceCount(t *testing.T) {
	t.Parallel()
	in := []string{"a", "b", "c", "a", "a", "c"}
	assert.Equal(t, 3, utils.SliceCount(in, "a"))
}
