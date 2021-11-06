package txtcities

import (
	"io"
	"strings"

	"github.com/Pitasi/alien-invasion/cities"
)

const (
	SpaceSeparator = " "
	EqualSign      = "="
)

// CityMarshaler is a wrapper over City for implementing a custom TxtMarshaler.
// The serialized format is a line composed as:
//   <city name> <neighbors>
// Neighbors are formatted as:
//   <direction>=<city name> <direction>=<city name> ...
type CityMarshaler struct {
	*cities.City
}

// TxtMarshal returns a byte slice representation of the city and its neighbors.
func (c *CityMarshaler) MarshalTxt() ([]byte, error) {
	b := new(strings.Builder)
	err := c.MarshalTxtToWriter(b)
	if err != nil {
		return nil, err
	}

	return []byte(b.String()), nil

}

// TxtMarshal writes a byte slice representation of the city and its neighbors
// to the specified writer.
func (c *CityMarshaler) MarshalTxtToWriter(w io.Writer) error {
	if err := marshalCityName(w, c.City); err != nil {
		return err
	}

	for _, dir := range c.AvailableDirections() {
		if err := marshalConnection(w, c.City, dir); err != nil {
			return err
		}
	}

	return nil
}

func marshalConnection(w io.Writer, c *cities.City, dir cities.Direction) error {
	if _, err := w.Write([]byte(SpaceSeparator)); err != nil {
		return err
	}

	if err := marshalDirection(w, dir); err != nil {
		return err
	}

	if _, err := w.Write([]byte(EqualSign)); err != nil {
		return err
	}

	if err := marshalCityName(w, c.Visit(dir)); err != nil {
		return err
	}

	return nil
}

func marshalCityName(w io.Writer, c *cities.City) error {
	_, err := w.Write([]byte(c.Name))
	return err
}

func marshalDirection(w io.Writer, d cities.Direction) error {
	_, err := w.Write([]byte(encodeDirection(d)))
	return err
}

func encodeDirection(dir cities.Direction) string {
	switch dir {
	case cities.NORTH:
		return "north"
	case cities.EAST:
		return "east"
	case cities.SOUTH:
		return "south"
	case cities.WEST:
		return "west"
	}

	panic("invalid direction")
}
