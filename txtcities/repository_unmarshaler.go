package txtcities

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/Pitasi/alien-invasion/cities"
)

// RepositoryUnmarshaler is a wrapper over Repository for implementing a custom
// TxtUnmarshaler.
// The serialized format is a multi-line string composed as:
//   <city name> <neighbors>
//   <city name> <neighbors>
//   <city name> <neighbors>
//
// Neighbors are formatted as:
//   <direction>=<city name> <direction>=<city name> ...
type RepositoryUnmarshaler struct {
	*cities.Repository
}

// Unmarshal unmarshals the given data into the repository.
func (r *RepositoryUnmarshaler) UnmarshalTxt(reader io.Reader) error {
	s := bufio.NewScanner(reader)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			continue
		}

		err := r.unmarshalLine(line)
		if err != nil {
			return fmt.Errorf("unmarshaling line: %w (line was %s)", err, line)
		}
	}

	if err := s.Err(); err != nil {
		return fmt.Errorf("reading stream: %w", err)
	}

	resultCities := r.GetAll()
	if err := AssertCitiesAreConnected(resultCities); err != nil {
		return fmt.Errorf("cities are not correctly connected: %w", err)
	}

	return nil
}

func (r *RepositoryUnmarshaler) unmarshalLine(line string) error {
	toks := strings.SplitN(line, SpaceSeparator, 6)

	if len(toks) == 1 {
		return fmt.Errorf("invalid city, no connections: %s", line)
	}

	if len(toks) == 6 {
		return fmt.Errorf("malformed line: %s", line)
	}

	cityName := toks[0]
	neighbors := toks[1:]

	city, err := r.getOrCreateCity(cityName)
	if err != nil {
		return fmt.Errorf("creating city: %w", err)
	}

	for _, neighborStr := range neighbors {
		err := r.unmarshalNeighbor(city, neighborStr)
		if err != nil {
			return fmt.Errorf("unmarshaling neighbor: %w", err)
		}
	}

	return nil
}

func (r *RepositoryUnmarshaler) getOrCreateCity(name string) (*cities.City, error) {
	if r.Has(name) {
		return r.Get(name), nil
	}

	city, err := cities.NewCity(name)
	if err != nil {
		return nil, err
	}

	err = r.Add(city)
	if err != nil {
		return nil, fmt.Errorf("adding new city to the repository: %w", err)
	}

	return city, nil
}

func (r *RepositoryUnmarshaler) unmarshalNeighbor(city *cities.City, s string) error {
	directionNeighborStr := strings.SplitN(s, EqualSign, 2)
	if len(directionNeighborStr) != 2 || len(directionNeighborStr[0]) == 0 || len(directionNeighborStr[1]) == 0 {
		return fmt.Errorf("malformed neighbor string: %s", s)
	}

	directionToken := directionNeighborStr[0]
	neighborName := directionNeighborStr[1]

	neighbor, err := r.getOrCreateCity(neighborName)
	if err != nil {
		return fmt.Errorf("creating neighbor: %w", err)
	}

	direction, err := decodeDirection(directionToken)
	if err != nil {
		return fmt.Errorf("invalid direction: %s", directionToken)
	}

	err = cities.Connect(city, neighbor, direction)
	if err != nil {
		return fmt.Errorf("connecting cities: %w", err)
	}

	return nil
}

func decodeDirection(tok string) (cities.Direction, error) {
	switch tok {
	case "north":
		return cities.NORTH, nil
	case "south":
		return cities.SOUTH, nil
	case "east":
		return cities.EAST, nil
	case "west":
		return cities.WEST, nil
	}

	return 0, fmt.Errorf("%s is not a vaild direction value", tok)
}
