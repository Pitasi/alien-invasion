package spawner

import (
	"testing"

	"github.com/Pitasi/alien-invasion/aliens"
	"github.com/Pitasi/alien-invasion/cities"
	"github.com/Pitasi/alien-invasion/world"
	"github.com/stretchr/testify/assert"
)

func TestNewSpawner(t *testing.T) {
	w := world.New()
	spawner, _ := New(AlienTemplate{
		MaximumMoves: 10,
		MovePolicy:   aliens.RandomMovePolicy,
	}, w, ChooseRandomCity)
	assert.NotNil(t, spawner)
}

func TestNewSpawner_WorldNil(t *testing.T) {
	_, err := New(AlienTemplate{
		MaximumMoves: 10,
		MovePolicy:   aliens.RandomMovePolicy,
	}, nil, ChooseRandomCity)
	assert.NotNil(t, err)
}

func TestNewSpawner_PolicyNil(t *testing.T) {
	_, err := New(AlienTemplate{
		MaximumMoves: 10,
		MovePolicy:   aliens.RandomMovePolicy,
	}, world.New(), nil)
	assert.NotNil(t, err)
}

func TestSpawnerSpawn(t *testing.T) {
	assert := assert.New(t)

	repo := cities.NewRepository()
	c, _ := cities.NewCity("city")
	repo.Add(c)
	w := world.NewWithCities(repo)

	spawner, _ := New(AlienTemplate{
		MaximumMoves: 10,
		MovePolicy:   aliens.RandomMovePolicy,
	}, w, ChooseRandomCity)

	spawner.Spawn(5)
	assert.Len(w.Aliens(), 5)

	spawner.Spawn(25)
	assert.Len(w.Aliens(), 30)
}

func TestSpawnerSpawn_EmptyWorld(t *testing.T) {
	assert := assert.New(t)

	w := world.New()

	spawner, _ := New(AlienTemplate{
		MaximumMoves: 10,
		MovePolicy:   aliens.RandomMovePolicy,
	}, w, ChooseRandomCity)

	err := spawner.Spawn(5)
	assert.NotNil(err)
}

func TestSpawnerSpawn_NegativeAliens(t *testing.T) {
	assert := assert.New(t)

	w := world.New()

	spawner, _ := New(AlienTemplate{
		MaximumMoves: 10,
		MovePolicy:   aliens.RandomMovePolicy,
	}, w, ChooseRandomCity)

	err := spawner.Spawn(-5)
	assert.NotNil(err)
}

func TestChooseRandomCity_Nil(t *testing.T) {
	_, err := ChooseRandomCity(nil)
	assert.NotNil(t, err)
}

func TestSpawnerSpawn_WrongTemplate(t *testing.T) {
	repo := cities.NewRepository()
	c, _ := cities.NewCity("city")
	repo.Add(c)
	w := world.NewWithCities(repo)

	table := []struct {
		message  string
		template AlienTemplate
	}{
		{
			message:  "empty template",
			template: AlienTemplate{},
		},
		{
			message: "missing MovePolicy",
			template: AlienTemplate{
				MaximumMoves: 10,
			},
		},
		{
			message: "negative maximum moves",
			template: AlienTemplate{
				MaximumMoves: -4,
			},
		},
	}

	for _, test := range table {
		t.Run(test.message, func(t *testing.T) {
			assert := assert.New(t)

			spawner, _ := New(test.template, w, ChooseRandomCity)
			err := spawner.Spawn(5)

			assert.NotNil(err)
		})
	}
}
