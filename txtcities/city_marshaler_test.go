package txtcities

import (
	"errors"
	"strings"
	"testing"

	"github.com/Pitasi/alien-invasion/cities"
	"github.com/Pitasi/alien-invasion/encoding/txt"
	"github.com/stretchr/testify/assert"
)

func marshalToString(item *CityMarshaler) string {
	b := new(strings.Builder)
	item.MarshalTxtToWriter(b)
	return b.String()
}

func TestCityTxt_NoNeighbors(t *testing.T) {
	assert := assert.New(t)

	city, _ := cities.NewCity("London")
	marshalled := marshalToString(&CityMarshaler{city})

	assert.Equal("London", marshalled)
}

func TestCityTxt_OneNeighbor(t *testing.T) {
	assert := assert.New(t)

	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")
	cities.Connect(london, newYork, cities.EAST)

	marshLondon := marshalToString(&CityMarshaler{london})

	assert.Equal("London east=NewYork", marshLondon)
}

func TestCityTxt_AllNeighors(t *testing.T) {
	assert := assert.New(t)

	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")
	chicago, _ := cities.NewCity("Chicago")
	pisa, _ := cities.NewCity("Pisa")
	rome, _ := cities.NewCity("Rome")

	cities.Connect(london, newYork, cities.EAST)
	cities.Connect(london, chicago, cities.SOUTH)
	cities.Connect(london, pisa, cities.WEST)
	cities.Connect(london, rome, cities.NORTH)

	marshLondon := marshalToString(&CityMarshaler{london})

	// we use Contains instead of Equal since the actual order of neighbors
	// is not important
	assert.Contains(marshLondon, "London ")
	assert.Contains(marshLondon, "east=NewYork")
	assert.Contains(marshLondon, "west=Pisa")
	assert.Contains(marshLondon, "south=Chicago")
	assert.Contains(marshLondon, "north=Rome")
}

type MockWriterFailsAfterNWrites struct {
	N int
}

func (w *MockWriterFailsAfterNWrites) Write(p []byte) (n int, err error) {
	w.N--
	if w.N == 0 {
		return 0, errors.New("mock error")
	}
	return len(p), nil
}

func TestCityMarshaler_Marshal_WriterError(t *testing.T) {
	assert := assert.New(t)

	london, _ := cities.NewCity("London")
	marsh := &CityMarshaler{london}

	newYork, _ := cities.NewCity("NewYork")
	chicago, _ := cities.NewCity("Chicago")

	cities.Connect(london, newYork, cities.EAST)
	cities.Connect(london, chicago, cities.WEST)

	// Simulate writer error happening after i writes.
	// This test is a bit brittle, but it's the only way to test the error
	// handling in (almost) every possible point. The number of writes may
	// actually change if the implementation changes.
	for i := 1; i <= 9; i++ {
		err := marsh.MarshalTxtToWriter(&MockWriterFailsAfterNWrites{i})
		assert.Error(err, "expected fail after %d writes", i)
	}
}

func TestCityMarshaler_Encoder(t *testing.T) {
	assert := assert.New(t)

	london, _ := cities.NewCity("London")
	marsh := &CityMarshaler{london}

	encoder := txt.NewEncoder(new(strings.Builder))
	err := encoder.Encode(marsh)

	assert.Nil(err)
}
