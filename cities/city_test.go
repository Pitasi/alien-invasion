package cities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCity(t *testing.T) {
	assert := assert.New(t)

	c, err := NewCity("Test")

	assert.Nil(err)
	assert.Equal("Test", c.Name)
}

func TestNewCity_NameWithSpace(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCity("Test with spaces")

	assert.NotNil(err)
}

func TestNewCity_NameWithEquals(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCity("Testwith=equal")

	assert.NotNil(err)
}

func TestNewCity_EmptyName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCity("")

	assert.NotNil(err)
}

func TestVisit_Empty(t *testing.T) {
	assert := assert.New(t)

	c, _ := NewCity("Test")

	for _, d := range AllDirections {
		assert.Nil(c.Visit(d))
	}
}
