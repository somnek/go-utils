package utils_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/somnek/go-utils/utils"
)

func TestMapKeys(t *testing.T) {
	out := []int{1, 2, 3}
	in := map[int]struct{}{
		1: {},
		2: {},
		3: {},
	}

	v := utils.MapKeys(in)
	assert.Equal(t, 3, len(v))
	sort.SliceStable(v, func(i, j int) bool {
		return v[i] < v[j]
	})
	for i := range out {
		assert.Equal(t, out[i], v[i])
	}
}
