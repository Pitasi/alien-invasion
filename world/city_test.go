package world

import (
	"testing"

	"github.com/Pitasi/alien-invasion/cities"
	"github.com/stretchr/testify/assert"
)

type MockAlienCounter struct {
	c *cities.City
	n int
}

func (m MockAlienCounter) AliensCount(c *cities.City) int {
	if c == m.c {
		return m.n
	}
	return 0
}

func TestWorldCity_AliensCount(t *testing.T) {
	assert := assert.New(t)

	c, _ := cities.NewCity("city")
	wc := WorldCity{
		City:    c,
		tracker: MockAlienCounter{c, 42},
	}

	count := wc.AliensCount()
	assert.Equal(42, count)
}
