package utils_test

import (
	"testing"

	"github.com/somnek/go-utils/utils"
	"github.com/stretchr/testify/assert"
)

func TestSliceRemove(t *testing.T) {
	t.Parallel()

	// test 1
	in1 := []int{1, 2, 3, 4, 5}
	x1 := 3
	out1 := []int{1, 2, 4, 5}

	// test 2
	in2 := []string{}
	x2 := "foo"
	out2 := []string{}

	// test 3
	in3 := []string{"foo", "bar", "baz"}
	x3 := "foo"
	out3 := []string{"bar", "baz"}

	assert.Equal(t, out1, utils.SliceRemove(in1, x1))
	assert.Equal(t, out2, utils.SliceRemove(in2, x2))
	assert.Equal(t, out3, utils.SliceRemove(in3, x3))
}
