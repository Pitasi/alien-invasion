package world

import (
	"github.com/Pitasi/alien-invasion/aliens"
	"github.com/Pitasi/alien-invasion/cities"
)

// Invalidater tracks an alien position and gets called after each move.
type Invalidater interface {
	Invalidate(*aliens.Alien, *cities.City) error
}

// WorldAlien represents an Alien in the world.
type WorldAlien struct {
	*aliens.Alien
	tracker Invalidater
}

func (w *WorldAlien) Move() error {
	from, _ := w.Alien.Move()
	err := w.tracker.Invalidate(w.Alien, from)
	return err
}
