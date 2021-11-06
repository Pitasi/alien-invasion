package cities

// Direction represents one of the four cardinal directions.
type Direction int

const (
	NORTH Direction = iota
	SOUTH
	EAST
	WEST
)

var AllDirections = []Direction{NORTH, SOUTH, EAST, WEST}

// Opposite returns the opposite direction. It panics if the given direction is
// not one of the ones provided by this package.
func Opposite(d Direction) Direction {
	switch d {
	case NORTH:
		return SOUTH
	case SOUTH:
		return NORTH
	case EAST:
		return WEST
	case WEST:
		return EAST
	}

	// instead of returning an error, panic because something "funny" is going
	// on. The given argument was not one of the constant provided by this
	// package.
	panic("Invalid direction")
}
