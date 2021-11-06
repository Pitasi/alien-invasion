package aliens

import (
	"math/rand"
	"testing"

	"github.com/Pitasi/alien-invasion/cities"
	"github.com/stretchr/testify/assert"
)

func TestNewAlien(t *testing.T) {
	c, _ := cities.NewCity("city")
	a, _ := NewAlien("name", c, RandomMovePolicy, 0)
	assert.NotNil(t, a)
}

func TestNewAlien_NilCity(t *testing.T) {
	_, err := NewAlien("name", nil, RandomMovePolicy, 0)
	assert.NotNil(t, err)
}

func TestNewAlien_NilMovePolicy(t *testing.T) {
	c, _ := cities.NewCity("city")
	_, err := NewAlien("name", c, nil, 0)
	assert.NotNil(t, err)
}

func TestNewAlien_NegativeMoveNumber(t *testing.T) {
	c, _ := cities.NewCity("city")
	_, err := NewAlien("name", c, RandomMovePolicy, -1)
	assert.NotNil(t, err)
}

func TestAlien_CanMove_NoNeighbors(t *testing.T) {
	assert := assert.New(t)

	cityNoNeighbors, _ := cities.NewCity("city")
	alien, _ := NewAlien("name", cityNoNeighbors, RandomMovePolicy, 100)

	assert.False(alien.CanMove())
}

func TestAlien_CanMove_Neighbors(t *testing.T) {
	assert := assert.New(t)

	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")
	cities.Connect(london, newYork, cities.EAST)

	alien, _ := NewAlien("name", london, RandomMovePolicy, 100)

	assert.True(alien.CanMove())
}

func TestAlien_CanMove_Neighbors_NoMovesRemaining(t *testing.T) {
	assert := assert.New(t)

	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")
	cities.Connect(london, newYork, cities.EAST)

	alien, _ := NewAlien("name", london, RandomMovePolicy, 0)

	assert.False(alien.CanMove())
}

func StayInCity(c *cities.City) *cities.City {
	return c
}

func TestAlien_Move(t *testing.T) {
	assert := assert.New(t)

	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")
	cities.Connect(london, newYork, cities.EAST)

	alien, _ := NewAlien("name", london, StayInCity, 100)
	from, to := alien.Move()

	assert.Equal(from, to)
	assert.Equal(from, london)
}

func TestAlien_MovesCountIsNotUpdatedWhenStill(t *testing.T) {
	assert := assert.New(t)

	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")
	cities.Connect(london, newYork, cities.EAST)

	alien, _ := NewAlien("name", london, StayInCity, 1)
	alien.Move()
	alien.Move()
	alien.Move()

	assert.True(alien.CanMove())
}

func AlwaysMove(c *cities.City) *cities.City {
	return c.Visit(c.AvailableDirections()[0])
}

func TestAlien_MovesCountIsUpdatedWhenNotStill(t *testing.T) {
	assert := assert.New(t)

	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")
	cities.Connect(london, newYork, cities.EAST)

	alien, _ := NewAlien("name", london, AlwaysMove, 1)
	alien.Move()

	assert.False(alien.CanMove())
}

func TestAlien_DontMoveWhenNoMovesLeft(t *testing.T) {
	assert := assert.New(t)

	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")
	cities.Connect(london, newYork, cities.EAST)

	alien, _ := NewAlien("name", london, AlwaysMove, 0)
	from, to := alien.Move()

	assert.Equal(from, to)
}

func TestRandomMovePolicy(t *testing.T) {
	assert := assert.New(t)

	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")
	chicago, _ := cities.NewCity("Chicago")
	rome, _ := cities.NewCity("Rome")
	moscow, _ := cities.NewCity("Moscow")
	cities.Connect(london, newYork, cities.NORTH)
	cities.Connect(london, chicago, cities.WEST)
	cities.Connect(london, rome, cities.EAST)
	cities.Connect(london, moscow, cities.SOUTH)

	rand.Seed(0)
	next := RandomMovePolicy(london)
	assert.Equal(london, next)

	next = RandomMovePolicy(london)
	assert.Equal(london, next)

	next = RandomMovePolicy(london)
	assert.Equal(chicago, next)
}
