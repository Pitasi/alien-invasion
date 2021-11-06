package txtcities

import (
	"strings"
	"testing"

	"github.com/Pitasi/alien-invasion/cities"
	"github.com/Pitasi/alien-invasion/encoding/txt"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryMarshaler_MarshalTxtToWriter(t *testing.T) {
	assert := assert.New(t)

	repo := cities.NewRepository()
	marsh := &RepositoryMarshaler{repo}

	b := new(strings.Builder)
	err := marsh.MarshalTxtToWriter(b)

	assert.Nil(err)
}

func TestRepositoryMarshaler_MarshalTxtToWriter_LinesCount(t *testing.T) {
	assert := assert.New(t)

	repo := cities.NewRepository()
	moscow, _ := cities.NewCity("Moscow")
	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")

	cities.Connect(moscow, london, cities.EAST)
	cities.Connect(london, moscow, cities.WEST)
	cities.Connect(london, newYork, cities.NORTH)
	cities.Connect(newYork, london, cities.SOUTH)

	repo.Add(moscow)
	repo.Add(london)
	repo.Add(newYork)

	marsh := &RepositoryMarshaler{repo}

	b := new(strings.Builder)
	marsh.MarshalTxtToWriter(b)
	out := strings.TrimSpace(b.String())

	lines := strings.Split(out, "\n")
	assert.Len(lines, 3)
}

func TestRepositoryMarshaler_MarshalTxtToWriter_IgnoreDisconnectedCities(t *testing.T) {
	assert := assert.New(t)

	repo := cities.NewRepository()
	moscow, _ := cities.NewCity("Moscow")
	london, _ := cities.NewCity("London")

	cities.Connect(moscow, london, cities.EAST)

	repo.Add(moscow)
	repo.Add(london)

	marsh := &RepositoryMarshaler{repo}

	b := new(strings.Builder)
	marsh.MarshalTxtToWriter(b)
	out := b.String()

	assert.Contains(out, "Moscow ")
	assert.NotContains(out, "London ")
}

func TestRepositoryMarshaler_MarshalTxt_Encoder(t *testing.T) {
	assert := assert.New(t)

	repo := cities.NewRepository()
	moscow, _ := cities.NewCity("Moscow")
	london, _ := cities.NewCity("London")
	newYork, _ := cities.NewCity("NewYork")

	cities.Connect(moscow, london, cities.EAST)
	cities.Connect(london, moscow, cities.WEST)
	cities.Connect(london, newYork, cities.NORTH)
	cities.Connect(newYork, london, cities.SOUTH)

	repo.Add(moscow)
	repo.Add(london)
	repo.Add(newYork)

	marsh := &RepositoryMarshaler{repo}

	b := new(strings.Builder)
	enc := txt.NewEncoder(b)
	err := enc.Encode(marsh)
	assert.Nil(err)

	out := b.String()

	assert.Contains(out, "Moscow ")
	assert.Contains(out, "London ")
	assert.Contains(out, "NewYork ")
}
