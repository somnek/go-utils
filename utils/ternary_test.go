package utils_test

import (
	"testing"

	"github.com/somnek/go-utils/utils"
	"github.com/stretchr/testify/assert"
)

func TestTernary(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "a", utils.Ternary(true, "a", "b"))
	assert.Equal(t, "b", utils.Ternary(1 == 2, "b", "b"))
}
