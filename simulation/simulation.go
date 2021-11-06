// Package simulation implements the business logic for running an alien
// invasion simulation.
package simulation

import (
	"fmt"

	"github.com/Pitasi/alien-invasion/world"
)

// Simulation is a simulation of an alien invasion in a certain world.
type Simulation struct {
	AliensCountDestroyCity int

	world  *world.World
	events chan Event
}

// New creates a new simulation in the specified world.
// The events chan can be nil, in this case no events will be dispatched.
func New(
	world *world.World,
	aliensCountDestroyCity int,
	events chan Event,
) (*Simulation, error) {
	if world == nil {
		return nil, fmt.Errorf("world cannot be nil")
	}

	if aliensCountDestroyCity < 0 {
		return nil, fmt.Errorf("the number of aliens required to destroy a city cannot be negative")
	}

	return &Simulation{
		world:                  world,
		AliensCountDestroyCity: aliensCountDestroyCity,
		events:                 events,
	}, nil
}

// Run starts the simulation.
func (s *Simulation) Run() error {
	for {
		if s.endConditions() {
			break
		}
		s.destroyCities()
		err := s.moveAliens()
		if err != nil {
			return err
		}
	}

	if s.events != nil {
		close(s.events)
	}
	return nil
}

func (s *Simulation) endConditions() bool {
	aliens := s.world.Aliens()
	return len(aliens) == 0 || !aliensCanMove(aliens)
}

func (s *Simulation) destroyCities() {
	for _, c := range s.world.Cities() {
		count := c.AliensCount()
		if count >= s.AliensCountDestroyCity {
			destroyedAliens := s.world.Destroy(c)
			s.emitDestroyEvent(c, destroyedAliens)
		}
	}
}

func (s *Simulation) emitDestroyEvent(c *world.WorldCity, destroyedAliens []*world.WorldAlien) {
	s.emit(&CityDestroyedEvent{
		City:            c,
		DestroyedAliens: destroyedAliens,
	})
}

func (s *Simulation) emit(e interface{}) {
	if s.events != nil {
		s.events <- e
	}
}

func (s *Simulation) moveAliens() error {
	for _, a := range s.world.Aliens() {
		err := a.Move()
		if err != nil {
			return fmt.Errorf("moving alien %s: %w", a.Name, err)
		}
	}

	return nil
}

func aliensCanMove(aliens []*world.WorldAlien) bool {
	for _, a := range aliens {
		if a.CanMove() {
			return true
		}
	}
	return false
}
