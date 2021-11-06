package txtcities

import (
	"testing"

	"github.com/Pitasi/alien-invasion/cities"
	"github.com/stretchr/testify/assert"
)

func TestAssertCitiesAreConnected(t *testing.T) {
	assert := assert.New(t)

	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")

	cities.Connect(london, newYork, cities.NORTH)
	cities.Connect(newYork, london, cities.SOUTH)

	err := AssertCitiesAreConnected([]*cities.City{london, newYork})
	assert.Nil(err)
}

func TestAssertCitiesAreConnected_NoConnections(t *testing.T) {
	assert := assert.New(t)

	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")

	err := AssertCitiesAreConnected([]*cities.City{london, newYork})
	assert.NotNil(err)
}

func TestAssertCitiesAreConnected_MissingConnection(t *testing.T) {
	assert := assert.New(t)

	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")

	cities.Connect(london, newYork, cities.NORTH)

	err := AssertCitiesAreConnected([]*cities.City{london, newYork})
	assert.NotNil(err)
}

func TestAssertCitiesAreConnected_MissingConnectionWrongDirection(t *testing.T) {
	assert := assert.New(t)

	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")

	cities.Connect(london, newYork, cities.NORTH)
	cities.Connect(newYork, london, cities.EAST)

	err := AssertCitiesAreConnected([]*cities.City{london, newYork})
	assert.NotNil(err)
}
