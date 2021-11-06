package world

import (
	"errors"
	"testing"

	"github.com/Pitasi/alien-invasion/aliens"
	"github.com/Pitasi/alien-invasion/cities"
	"github.com/stretchr/testify/assert"
)

type MockInvalidater struct {
	ShouldError bool
}

func (m MockInvalidater) Invalidate(a *aliens.Alien, c *cities.City) error {
	if m.ShouldError {
		return errors.New("error")
	}
	return nil
}

func TestWorldAlien_Move(t *testing.T) {
	assert := assert.New(t)

	c, _ := cities.NewCity("city")
	a, _ := aliens.NewAlien("a", c, aliens.RandomMovePolicy, 10)

	wa := &WorldAlien{
		Alien:   a,
		tracker: MockInvalidater{},
	}

	err := wa.Move()
	assert.Nil(err)
}

func TestWorldAlien_Move_Error(t *testing.T) {
	assert := assert.New(t)

	c, _ := cities.NewCity("city")
	a, _ := aliens.NewAlien("a", c, aliens.RandomMovePolicy, 10)

	wa := &WorldAlien{
		Alien:   a,
		tracker: MockInvalidater{ShouldError: true},
	}

	err := wa.Move()
	assert.NotNil(err)
}
