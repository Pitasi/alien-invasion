package aliens

import (
	"testing"

	"github.com/Pitasi/alien-invasion/cities"
	"github.com/stretchr/testify/assert"
)

func TestRepository_New(t *testing.T) {
	assert := assert.New(t)
	repo := NewRepository()
	assert.NotNil(repo)
}

func TestRepository_Add(t *testing.T) {
	assert := assert.New(t)
	repo := NewRepository()

	c, _ := cities.NewCity("city")
	a, _ := NewAlien("name", c, RandomMovePolicy, 0)
	repo.Add(a)

	allAliens := repo.GetAll()
	assert.Len(allAliens, 1)
	assert.Contains(allAliens, a)
}

func TestRepository_Remove(t *testing.T) {
	assert := assert.New(t)
	repo := NewRepository()

	c, _ := cities.NewCity("city")
	a, _ := NewAlien("name", c, RandomMovePolicy, 0)
	repo.Add(a)

	repo.Remove(a)

	allAliens := repo.GetAll()
	assert.Len(allAliens, 0)
}
