package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/somnek/go-utils/utils"
)

func TestSliceContain(t *testing.T) {
	t.Parallel()

	l := []string{"boo", "2", "z", ""}
	in := []string{"boo", "foo", ""}
	out := []bool{true, false, true}

	for i, v := range in {
		assert.Equal(t, out[i], utils.SliceContain(l, v))
	}
}
