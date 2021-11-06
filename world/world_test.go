package world

import (
	"strings"
	"testing"

	"github.com/Pitasi/alien-invasion/aliens"
	"github.com/Pitasi/alien-invasion/cities"
	"github.com/Pitasi/alien-invasion/encoding/txt"
	"github.com/stretchr/testify/assert"
)

func TestNewEmptyWorld(t *testing.T) {
	w := New()
	assert.NotNil(t, w)
}

func TestWorld_Empty_Cities(t *testing.T) {
	w := New()
	assert.Len(t, w.Cities(), 0)
}

func TestWorld_Empty_Aliens(t *testing.T) {
	w := New()
	assert.Len(t, w.Aliens(), 0)
}

func TestNewWorldFromRepo_CitiesCount(t *testing.T) {
	repo := cities.NewRepository()
	c, _ := cities.NewCity("city")
	repo.Add(c)

	w := NewWithCities(repo)

	assert.Len(t, w.Cities(), len(repo.GetAll()))
}

func TestWorld_AddAlien(t *testing.T) {
	assert := assert.New(t)

	repo := cities.NewRepository()
	c, _ := cities.NewCity("city")
	repo.Add(c)

	w := NewWithCities(repo)

	worldCity := w.Cities()[0]

	err := w.AddAlien(AlienConfig{
		Name:         "alien",
		MovePolicy:   aliens.RandomMovePolicy,
		MaximumMoves: 10,
	}, worldCity)

	assert.Nil(err)
	assert.Len(w.Aliens(), 1)
	assert.Equal(worldCity.AliensCount(), 1)
}

func TestWorld_AddAlien_WrongConfig(t *testing.T) {
	repo := cities.NewRepository()
	c, _ := cities.NewCity("city")
	repo.Add(c)
	w := NewWithCities(repo)
	worldCity := w.Cities()[0]

	table := []struct {
		message string
		config  AlienConfig
	}{
		{
			message: "empty move policy",
			config: AlienConfig{
				Name:         "alien",
				MaximumMoves: 10,
			},
		},
		{
			message: "negative maximum moves",
			config: AlienConfig{
				Name:         "alien",
				MovePolicy:   aliens.RandomMovePolicy,
				MaximumMoves: -10,
			},
		},
	}

	for _, test := range table {
		t.Run(test.message, func(t *testing.T) {
			assert := assert.New(t)

			err := w.AddAlien(test.config, worldCity)

			assert.NotNil(err)
		})
	}
}

func TestWorld_Destroy(t *testing.T) {
	assert := assert.New(t)

	repo := cities.NewRepository()
	c, _ := cities.NewCity("city")
	repo.Add(c)

	w := NewWithCities(repo)

	worldCity := w.Cities()[0]

	w.AddAlien(AlienConfig{
		Name:         "alien",
		MovePolicy:   aliens.RandomMovePolicy,
		MaximumMoves: 10,
	}, worldCity)

	w.Destroy(worldCity)

	assert.Len(w.Cities(), 0)
	assert.Len(w.Aliens(), 0)
}

func TestWorld_Empty_Encoder(t *testing.T) {
	assert := assert.New(t)

	w := New()

	b := new(strings.Builder)
	enc := txt.NewEncoder(b)
	err := enc.Encode(w)

	assert.Nil(err)
}

func TestWorld_Encoder(t *testing.T) {
	assert := assert.New(t)

	repo := cities.NewRepository()
	c, _ := cities.NewCity("city")
	repo.Add(c)
	w := NewWithCities(repo)

	b := new(strings.Builder)
	enc := txt.NewEncoder(b)
	err := enc.Encode(w)

	assert.Nil(err)
}
func TestWorld_Decoder(t *testing.T) {
	assert := assert.New(t)

	content := `
London south=Chicago
Chicago north=London
`

	w := New()
	dec := txt.NewDecoder(strings.NewReader(content))
	err := dec.Decode(w)

	assert.Nil(err)
	assert.Len(w.Cities(), 2)
}
