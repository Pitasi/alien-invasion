package simulation

import (
	"github.com/Pitasi/alien-invasion/world"
)

type Event interface{}

// CityDestroyedEvent is sent when a city is destroyed.
type CityDestroyedEvent struct {
	City            *world.WorldCity
	DestroyedAliens []*world.WorldAlien
}
