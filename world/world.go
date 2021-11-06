// Package world is an abstraction layer that glues together aliens and cities.
// It should be used by business logic of accessing directly the underlying
// layers.
package world

import (
	"fmt"
	"io"

	"github.com/Pitasi/alien-invasion/aliens"
	"github.com/Pitasi/alien-invasion/cities"
	"github.com/Pitasi/alien-invasion/txtcities"
)

// World is composed by cities and aliens that moves between them.
type World struct {
	citiesRepo    *cities.Repository
	aliensRepo    *aliens.Repository
	aliensTracker *aliens.Tracker
}

// New return a new empty world with no cities or aliens.
func New() *World {
	return &World{
		citiesRepo:    cities.NewRepository(),
		aliensRepo:    aliens.NewRepository(),
		aliensTracker: aliens.NewTracker(),
	}
}

// New return a new world with no aliens, using the specified city
// repository.
func NewWithCities(citiesRepo *cities.Repository) *World {
	return &World{
		citiesRepo:    citiesRepo,
		aliensRepo:    aliens.NewRepository(),
		aliensTracker: aliens.NewTracker(),
	}
}

// Cities returns a slice of all the cities present in the world.
func (w *World) Cities() []*WorldCity {
	cities := w.citiesRepo.GetAll()
	return w.wrapCities(cities)
}

// Aliens returns a slice of the aliens present in the world.
func (w *World) Aliens() []*WorldAlien {
	aliens := w.aliensRepo.GetAll()
	return w.wrapAliens(aliens)
}

// AlienConfig contains the configuraition for adding new aliens to the world.
type AlienConfig struct {
	Name         string
	MovePolicy   aliens.MovePolicy
	MaximumMoves int
}

// AddAlien creates a new alien and adds it to the world.
func (w *World) AddAlien(config AlienConfig, c *WorldCity) error {
	a, err := aliens.NewAlien(config.Name, c.City, config.MovePolicy, config.MaximumMoves)
	if err != nil {
		return fmt.Errorf("adding alien: %w", err)
	}

	w.aliensRepo.Add(a)
	w.aliensTracker.Add(a)
	return nil
}

// DestroyCity removes the city from the world, returns the aliens that were in
// the city.
func (w *World) Destroy(c *WorldCity) []*WorldAlien {
	cities.Destroy(c.City)
	w.citiesRepo.Remove(c.City)
	destroyedAliens := w.aliensTracker.Destroy(c.City)
	w.aliensRepo.Remove(destroyedAliens...)
	return w.wrapAliens(destroyedAliens)
}

func (w *World) UnmarshalTxt(r io.Reader) error {
	u := &txtcities.RepositoryUnmarshaler{Repository: w.citiesRepo}
	return u.UnmarshalTxt(r)
}

func (w *World) MarshalTxt() ([]byte, error) {
	m := &txtcities.RepositoryMarshaler{Repository: w.citiesRepo}
	return m.MarshalTxt()
}

func (w *World) wrapCities(cs []*cities.City) []*WorldCity {
	worldCities := make([]*WorldCity, 0, len(cs))
	for _, c := range cs {
		worldCities = append(worldCities, &WorldCity{City: c, tracker: w.aliensTracker})
	}
	return worldCities
}

func (w *World) wrapAliens(as []*aliens.Alien) []*WorldAlien {
	worldAliens := make([]*WorldAlien, 0, len(as))
	for _, a := range as {
		worldAliens = append(worldAliens, &WorldAlien{Alien: a, tracker: w.aliensTracker})
	}
	return worldAliens
}
