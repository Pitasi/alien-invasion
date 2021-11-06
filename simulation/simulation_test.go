package simulation

import (
	"sync"
	"testing"

	"github.com/Pitasi/alien-invasion/aliens"
	"github.com/Pitasi/alien-invasion/cities"
	"github.com/Pitasi/alien-invasion/world"
	"github.com/stretchr/testify/assert"
)

// sampleWorld returns a world with three cities composed like this:
//
//             Chicago
//               |
//   NewYork - London
func sampleWorld() *world.World {
	repo := cities.NewRepository()
	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")
	chicago, _ := cities.NewCity("Chicago")

	cities.Connect(london, newYork, cities.EAST)
	cities.Connect(newYork, london, cities.WEST)
	cities.Connect(london, chicago, cities.NORTH)
	cities.Connect(chicago, london, cities.SOUTH)

	repo.Add(london)
	repo.Add(newYork)
	repo.Add(chicago)

	w := world.NewWithCities(repo)
	return w
}

func TestSimulationNew(t *testing.T) {
	world := sampleWorld()
	sim, _ := New(world, 2, nil)
	assert.NotNil(t, sim)
}

func TestSimulationNew_NilWorld(t *testing.T) {
	_, err := New(nil, 2, nil)
	assert.NotNil(t, err)
}

func TestSimulationNew_AlienCountNegative(t *testing.T) {
	_, err := New(world.New(), -42, nil)
	assert.NotNil(t, err)
}

func TestSimulationRun_NoAliens(t *testing.T) {
	world := sampleWorld()
	sim, _ := New(world, 2, nil)
	err := sim.Run()
	assert.Nil(t, err)
}

func TestSimulationRun_OneAlien(t *testing.T) {
	w := sampleWorld()
	cities := w.Cities()

	w.AddAlien(world.AlienConfig{
		Name:         "a",
		MovePolicy:   aliens.RandomMovePolicy,
		MaximumMoves: 10000,
	}, cities[0])

	sim, _ := New(w, 2, nil)
	err := sim.Run()

	assert.Nil(t, err)
}

func TestSimulationRun_ThreeAliens(t *testing.T) {
	assert := assert.New(t)

	w := sampleWorld()
	cs := w.Cities()

	// put an alien in one city
	w.AddAlien(world.AlienConfig{
		Name:         "a",
		MovePolicy:   aliens.RandomMovePolicy,
		MaximumMoves: 10000,
	}, cs[0])

	// put two aliens in the same city
	w.AddAlien(world.AlienConfig{
		Name:         "a",
		MovePolicy:   aliens.RandomMovePolicy,
		MaximumMoves: 10000,
	}, cs[2])
	w.AddAlien(world.AlienConfig{
		Name:         "a",
		MovePolicy:   aliens.RandomMovePolicy,
		MaximumMoves: 10000,
	}, cs[2])

	// set aliens threshold to two, cs[2] should be immediately destroyed
	sim, _ := New(w, 2, nil)
	err := sim.Run()

	assert.Nil(err)
	assert.Len(w.Cities(), 2)
}

func TestSimulationRun_Events(t *testing.T) {
	assert := assert.New(t)

	w := sampleWorld()
	cs := w.Cities()

	// put an alien in one city
	w.AddAlien(world.AlienConfig{
		Name:         "a",
		MovePolicy:   aliens.RandomMovePolicy,
		MaximumMoves: 10000,
	}, cs[0])

	// put two aliens in the same city
	w.AddAlien(world.AlienConfig{
		Name:         "a",
		MovePolicy:   aliens.RandomMovePolicy,
		MaximumMoves: 10000,
	}, cs[2])
	w.AddAlien(world.AlienConfig{
		Name:         "a",
		MovePolicy:   aliens.RandomMovePolicy,
		MaximumMoves: 10000,
	}, cs[2])

	// set aliens threshold to two, cs[2] should be immediately destroyed
	events := make(chan Event)
	sim, _ := New(w, 2, events)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Done()
		err := sim.Run()
		assert.Nil(err)
	}()

	event := <-events
	switch e := event.(type) {
	case *CityDestroyedEvent:
		assert.Equal(cs[2], e.City)
		assert.Len(e.DestroyedAliens, 2)
	default:
		assert.Fail("unexpected event type")
	}

	wg.Wait()

	_, open := <-events
	assert.False(open, "events channel not closed after simulation end")
}
