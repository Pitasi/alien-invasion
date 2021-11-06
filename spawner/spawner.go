// Package spawner implements logic to create new aliens and add them to a
// world.
package spawner

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Pitasi/alien-invasion/aliens"
	"github.com/Pitasi/alien-invasion/world"
)

// ChooseCityPolicy is a policy for choosing a city to spawn an alien in from a
// list of cities.
type ChooseCityPolicy func([]*world.WorldCity) (*world.WorldCity, error)

// ChooseRandomCity choose a random city among the available ones.
func ChooseRandomCity(cs []*world.WorldCity) (*world.WorldCity, error) {
	if len(cs) == 0 {
		return nil, errors.New("no cities to choose from")
	}

	i := rand.Intn(len(cs))
	return cs[i], nil
}

// Spawner adds new aliens to a world.
type Spawner struct {
	alienTemplate AlienTemplate
	world         *world.World
	cityChooser   ChooseCityPolicy
}

// AlienTemplate specifies the configuration of the new aliens.
type AlienTemplate struct {
	MaximumMoves int
	MovePolicy   aliens.MovePolicy
}

// New returns a new spawner.
func New(config AlienTemplate, world *world.World, cityChooser ChooseCityPolicy) (*Spawner, error) {
	if world == nil {
		return nil, errors.New("world cannot be nil")
	}

	if cityChooser == nil {
		return nil, errors.New("chose city policy cannot be nil")
	}

	return &Spawner{
		alienTemplate: config,
		world:         world,
		cityChooser:   cityChooser,
	}, nil
}

// Spawn spawns new aliens in the world using the configured policy.
func (s *Spawner) Spawn(count int) error {
	if count < 0 {
		return errors.New("count cannot be negative")
	}

	if count == 0 {
		return nil
	}

	cities := s.world.Cities()

	for i := 0; i < count; i++ {
		name := "A" + strconv.Itoa(i)
		c, err := s.cityChooser(cities)
		if err != nil {
			return fmt.Errorf("choosing spawn city: %w", err)
		}

		err = s.spawnAlien(name, c)
		if err != nil {
			return fmt.Errorf("spawning alien: %w", err)
		}
	}

	return nil
}

func (s *Spawner) spawnAlien(name string, c *world.WorldCity) error {
	return s.world.AddAlien(world.AlienConfig{
		Name:         name,
		MovePolicy:   s.alienTemplate.MovePolicy,
		MaximumMoves: s.alienTemplate.MaximumMoves,
	}, c)
}
