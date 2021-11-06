package aliens

import (
	"testing"

	"github.com/Pitasi/alien-invasion/cities"
	"github.com/stretchr/testify/assert"
)

func TestTracker_New(t *testing.T) {
	tracker := NewTracker()
	assert.NotNil(t, tracker)
}

func TestTracker_Add_NewCity(t *testing.T) {
	assert := assert.New(t)
	tracker := NewTracker()

	c, _ := cities.NewCity("city")
	a, _ := NewAlien("a", c, RandomMovePolicy, 0)

	tracker.Add(a)

	assert.Equal(1, tracker.AliensCount(c))
}

func TestTracker_Add_ExistingCity(t *testing.T) {
	assert := assert.New(t)
	tracker := NewTracker()

	c, _ := cities.NewCity("city")
	a, _ := NewAlien("a", c, RandomMovePolicy, 0)
	b, _ := NewAlien("b", c, RandomMovePolicy, 0)

	tracker.Add(a)
	tracker.Add(b)

	assert.Equal(2, tracker.AliensCount(c))
}

func TestTracker_Invalidate(t *testing.T) {
	assert := assert.New(t)
	tracker := NewTracker()

	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")
	cities.Connect(london, newYork, cities.EAST)

	a, _ := NewAlien("a", london, AlwaysMove, 1)
	tracker.Add(a)

	a.Move()
	tracker.Invalidate(a, london)

	assert.Equal(0, tracker.AliensCount(london))
	assert.Equal(1, tracker.AliensCount(newYork))
}

func TestTracker_Invalidate_WrongCity(t *testing.T) {
	assert := assert.New(t)
	tracker := NewTracker()

	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")
	cities.Connect(london, newYork, cities.EAST)

	a, _ := NewAlien("a", london, AlwaysMove, 1)
	tracker.Add(a)

	a.Move()
	err := tracker.Invalidate(a, newYork)

	assert.NotNil(err)
}

func TestTracker_Destroy(t *testing.T) {
	assert := assert.New(t)
	tracker := NewTracker()

	london, _ := cities.NewCity("London")

	a, _ := NewAlien("a", london, AlwaysMove, 1)
	tracker.Add(a)

	aliens := tracker.Destroy(london)

	assert.Len(aliens, 1)
	assert.Contains(aliens, a)
}

func TestTracker_Destroy_NonTrackedCity(t *testing.T) {
	assert := assert.New(t)
	tracker := NewTracker()

	london, _ := cities.NewCity("London")
	aliens := tracker.Destroy(london)

	assert.Len(aliens, 0)
}
