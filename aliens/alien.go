// Package aliens contains the Alien type and the other types required to track
// aliens in cities.
package aliens

import (
	"errors"
	"math/rand"

	"github.com/Pitasi/alien-invasion/cities"
)

// MovePolicy decides how to move from one city to the next.
type MovePolicy func(from *cities.City) *cities.City

// RandomMovePolicy randomly decides if moving from the current city, and if so,
// to which neighbor.
// The decision of moving or not and each neighbor all have the same
// probability.
//
// Note: the package math/rand is used to generate random numbers and should be
// seeded before using this function.
func RandomMovePolicy(from *cities.City) *cities.City {
	directions := from.AvailableDirections()
	choice := rand.Intn(len(directions) + 1)

	if choice == len(directions) {
		return from
	}

	return from.Visit(directions[choice])
}

// Alien is represent an alien in city.
type Alien struct {
	Name        string
	currentCity *cities.City
	movePolicy  MovePolicy
	movesCount  int
	maxMoves    int
}

func NewAlien(name string, c *cities.City, movePolicy MovePolicy, maxMoves int) (*Alien, error) {
	if c == nil {
		return nil, errors.New("city must be specified")
	}

	if movePolicy == nil {
		return nil, errors.New("move policy must be specified")
	}

	if maxMoves < 0 {
		return nil, errors.New("max moves must be greater than 0")
	}

	return &Alien{
		Name:        name,
		currentCity: c,
		movePolicy:  movePolicy,
		movesCount:  0,
		maxMoves:    maxMoves,
	}, nil
}

// CanMove returns true if the alien has not reached the maximum number of moves
// allowed and the city its currently in have at least one neighbor.
func (a *Alien) CanMove() bool {
	return a.movesCount < a.maxMoves && len(a.currentCity.AvailableDirections()) > 0
}

// Move moves the alien to the next city. If the alien has reached the maximum
// number of moves allowed, it will stay in the current city.
func (a *Alien) Move() (from, to *cities.City) {
	if !a.CanMove() {
		return a.currentCity, a.currentCity
	}

	current := a.currentCity
	next := a.movePolicy(a.currentCity)
	if next != current {
		a.currentCity = next
		a.movesCount++
	}

	return current, next
}
