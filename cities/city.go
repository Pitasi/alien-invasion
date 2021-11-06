// Package cities contains the City type and a repository for collecting them.
package cities

import (
	"errors"
	"fmt"
	"strings"
)

// City represents a city that can be connected to other cities (neighbors).
// Each city can have a maximum of 4 neighbors, one for each Direction.
type City struct {
	Name string

	neighbors map[Direction]*City
}

// NewCity creates a new city with the given name.
// Returns an error if the specified name is invalid: it cannot contains
// spaces or equals.
func NewCity(name string) (*City, error) {
	if len(name) == 0 {
		return nil, errors.New("city name must not be empty")
	}

	if strings.ContainsAny(name, " =") {
		return nil, fmt.Errorf("city name cannot contain spaces or equals")
	}

	return &City{
		Name: name,
		neighbors: map[Direction]*City{
			NORTH: nil,
			SOUTH: nil,
			EAST:  nil,
			WEST:  nil,
		},
	}, nil
}

// Visit returns the neighbor in the given direction.
func (c *City) Visit(d Direction) *City {
	return c.neighbors[d]
}

// AvailableDirections returns the list of directions in which there are neighbors.
func (c *City) AvailableDirections() []Direction {
	var directions = make([]Direction, 0, 4)
	for _, d := range AllDirections {
		if c.neighbors[d] != nil {
			directions = append(directions, d)
		}
	}
	return directions
}
