package cities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpposite(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		given    Direction
		expected Direction
	}{
		{NORTH, SOUTH},
		{SOUTH, NORTH},
		{EAST, WEST},
		{WEST, EAST},
	}

	for _, tc := range tt {
		assert.Equal(tc.expected, Opposite(tc.given))
	}
}

func TestOpposite_Panics(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		Opposite(Direction(-1))
	})
}
