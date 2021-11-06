package txtcities

import (
	"fmt"

	"github.com/Pitasi/alien-invasion/cities"
)

// AssertCitiesAreConnected returns an error if any of the cities are not
// correctly connected with each other. This mean that a city has no roads, or
// there is a road that is not bidirectional (you can go from city A to city B,
// but not the other way around).
func AssertCitiesAreConnected(c []*cities.City) error {
	for _, city := range c {
		if !cityHasNeighbors(city) {
			return fmt.Errorf("city %s has no connections", city.Name)
		}

		if ok, neighbor := neighborsAreConnectedBack(city); !ok {
			return fmt.Errorf("city %s connected to %s, but %s is not connected to %s", city.Name, neighbor.Name, neighbor.Name, city.Name)
		}
	}
	return nil
}

func cityHasNeighbors(c *cities.City) bool {
	dirs := c.AvailableDirections()
	return len(dirs) > 0
}

func neighborsAreConnectedBack(c *cities.City) (bool, *cities.City) {
	for _, dir := range c.AvailableDirections() {
		neighbor := c.Visit(dir)
		expectedSelf := neighbor.Visit(cities.Opposite(dir))
		if expectedSelf != c {
			return false, neighbor
		}
	}

	return true, nil
}
