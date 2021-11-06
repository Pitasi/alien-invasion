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

func TestVisit_North(t *testing.T) {
	assert := assert.New(t)

	london, _ := NewCity("London")
	c, _ := NewCity("Test")
	Connect(c, london, NORTH)

	assert.Equal(london, c.Visit(NORTH))
}

func TestConnect(t *testing.T) {
	assert := assert.New(t)

	london, _ := NewCity("London")
	c, _ := NewCity("Test")

	err := Connect(c, london, NORTH)

	assert.Nil(err)
	assert.Equal(london, c.Visit(NORTH))
	assert.Nil(london.Visit(SOUTH))
}

func TestConnect_AlreadyOccupied(t *testing.T) {
	assert := assert.New(t)

	london, _ := NewCity("London")
	c, _ := NewCity("Test")

	err := Connect(c, london, NORTH)
	assert.Nil(err)

	newYork, _ := NewCity("New York")
	err2 := Connect(c, newYork, NORTH)
	assert.NotNil(err2)
}

func TestConnect_MultipleSameCitiesDifferentDirection(t *testing.T) {
	assert := assert.New(t)

	london, _ := NewCity("London")
	c, _ := NewCity("Test")

	err := Connect(c, london, NORTH)
	assert.Nil(err)

	err = Connect(c, london, EAST)
	assert.Nil(err)

	err = Connect(c, london, WEST)
	assert.Nil(err)

	err = Connect(c, london, SOUTH)
	assert.Nil(err)
}

func TestConnect_WithItself(t *testing.T) {
	assert := assert.New(t)

	c, _ := NewCity("Test")

	err := Connect(c, c, NORTH)
	assert.NotNil(err)
}

func TestAvailableDirection_Empty(t *testing.T) {
	assert := assert.New(t)

	c, _ := NewCity("Test")

	directions := c.AvailableDirections()

	assert.Equal(0, len(directions))
}

func TestAvailableDirection_One(t *testing.T) {
	assert := assert.New(t)

	london, _ := NewCity("London")
	c, _ := NewCity("Test")
	Connect(c, london, NORTH)

	directions := c.AvailableDirections()

	assert.Equal(1, len(directions))
	assert.Contains(directions, NORTH)
}

func TestAvailableDirection_All(t *testing.T) {
	assert := assert.New(t)

	london, _ := NewCity("London")
	newYork, _ := NewCity("NewYork")
	rome, _ := NewCity("Rome")
	paris, _ := NewCity("Paris")
	c, _ := NewCity("Test")

	Connect(c, london, NORTH)
	Connect(c, newYork, SOUTH)
	Connect(c, rome, EAST)
	Connect(c, paris, WEST)

	directions := c.AvailableDirections()

	assert.Equal(4, len(directions))
	assert.Contains(directions, NORTH)
	assert.Contains(directions, SOUTH)
	assert.Contains(directions, EAST)
	assert.Contains(directions, WEST)
}

func Test_Destroy(t *testing.T) {
	assert := assert.New(t)

	london, _ := NewCity("London")
	c, _ := NewCity("Test")
	Connect(c, london, NORTH)

	Destroy(c)

	assert.Empty(c.AvailableDirections())
	assert.Nil(london.Visit(SOUTH))
}

func Test_Destroy_NoNeighbors(t *testing.T) {
	assert := assert.New(t)

	c, _ := NewCity("Test")

	Destroy(c)

	assert.Empty(c.AvailableDirections())
}

func Test_Destroy_Nil(t *testing.T) {
	assert := assert.New(t)
	assert.NotPanics(func() { Destroy(nil) })
}
