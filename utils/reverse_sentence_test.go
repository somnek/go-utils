package utils_test

import (
	"testing"

	"github.com/somnek/go-utils/utils"
	"github.com/stretchr/testify/assert"
)

func TestReverseSentence(t *testing.T) {
	t.Parallel()

	input := []string{
		"Hello how are you?",
		"I don't know . . .",
		"🍤 🐟 🍆",
	}
	output := []string{
		"you? are how Hello",
		". . . know don't I",
		"🍆 🐟 🍤",
	}

	for i, s := range input {
		assert.Equal(t, output[i], utils.ReverseSentence(s))
	}
}
