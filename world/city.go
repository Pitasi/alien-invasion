package world

import (
	"github.com/Pitasi/alien-invasion/cities"
)

// AlienCounter returns the number of aliens in the city.
type AlienCounter interface {
	AliensCount(c *cities.City) int
}

// WorldCity represents a city in the world.
type WorldCity struct {
	*cities.City
	tracker AlienCounter
}

// AliensCount returns the number of aliens in the city.
func (c *WorldCity) AliensCount() int {
	return c.tracker.AliensCount(c.City)
}
