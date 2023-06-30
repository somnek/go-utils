package utils_test

import (
	"testing"

	"github.com/somnek/go-utils/utils"
	"github.com/stretchr/testify/assert"
)

func TestMapKeys(t *testing.T) {
	t.Parallel()

	input := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	output := []string{
		"key1",
		"key2",
		"key3",
	}
	assert.Equal(t, output, utils.MapKeys(input))
	
}

