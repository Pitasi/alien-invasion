package aliens

import (
	"fmt"

	"github.com/Pitasi/alien-invasion/cities"
)

var presentFlag = struct{}{}

// Presence is a convenient collection of aliens that have O(1) access cost.
type Presence map[*Alien]struct{}

// Tracker tracks the presence of aliens in cities.
type Tracker struct {
	presences map[*cities.City]Presence
}

// NewTracker creates a new empty tracker.
func NewTracker() *Tracker {
	return &Tracker{
		presences: make(map[*cities.City]Presence),
	}
}

// Add adds an alien to the tracker in the current alien position.
func (at *Tracker) Add(a *Alien) {
	at.initCity(a.currentCity)
	at.presences[a.currentCity][a] = presentFlag
}

// Invalidate marks an alien as not present in `from` city and make the tracker
// updates its position.
func (at *Tracker) Invalidate(a *Alien, from *cities.City) error {
	if !at.isInCity(a, from) {
		return fmt.Errorf("alien %v was not in city %v", a, from)
	}

	at.remove(a, from)
	at.Add(a)
	return nil
}

// AliensCount returns the number of aliens in the specified city.
func (at *Tracker) AliensCount(c *cities.City) int {
	return len(at.presences[c])
}

// Destroy destroys a city and all the aliens inside it. The destroyed aliens
// are returned.
func (at *Tracker) Destroy(c *cities.City) []*Alien {
	presences := at.presences[c]
	aliens := make([]*Alien, 0, len(presences))
	for a := range presences {
		aliens = append(aliens, a)
	}

	return aliens
}

func (at *Tracker) isInCity(a *Alien, c *cities.City) bool {
	_, cityRegistered := at.presences[c]
	if !cityRegistered {
		return false
	}

	_, alienInCity := at.presences[c][a]
	return alienInCity
}

func (at *Tracker) remove(a *Alien, c *cities.City) {
	delete(at.presences[c], a)
}

func (at *Tracker) initCity(c *cities.City) {
	if at.presences[c] == nil {
		at.presences[c] = make(Presence)
	}
}
