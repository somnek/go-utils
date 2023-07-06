package utils_test

import (
	"testing"

	"github.com/somnek/go-utils/utils"
	"github.com/stretchr/testify/assert"
)

func TestSliceEuqal(t *testing.T) {
	t.Parallel()
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c"}
	c := []string{"e", "f", "g"}
	d := []string{"a", "b", "c", "d"}
	e := []string{"b", "c", "a"}

	assert.True(t, utils.SliceEqual(a, b))
	assert.False(t, utils.SliceEqual(a, c))
	assert.False(t, utils.SliceEqual(a, d))
	assert.False(t, utils.SliceEqual(a, e))

}
