package cities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	repo := NewRepository()
	assert.NotNil(t, repo)
}

func TestRepository_AddCity(t *testing.T) {
	assert := assert.New(t)

	repo := NewRepository()
	city, _ := NewCity("Test")

	err := repo.Add(city)

	assert.Nil(err)
}

func TestRepository_Has(t *testing.T) {
	assert := assert.New(t)

	repo := NewRepository()
	city, _ := NewCity("Test")
	repo.Add(city)

	assert.True(repo.Has("Test"))
	assert.False(repo.Has("AnotherName"))
}

func TestRepository_Get(t *testing.T) {
	assert := assert.New(t)

	repo := NewRepository()
	city, _ := NewCity("Test")
	repo.Add(city)

	c := repo.Get("Test")
	assert.Equal(city, c)
}

func TestRepository_Get_Missing(t *testing.T) {
	assert := assert.New(t)

	repo := NewRepository()

	c := repo.Get("Test")
	assert.Nil(c)
}

func TestAddCitySameName(t *testing.T) {
	assert := assert.New(t)

	repo := NewRepository()
	city, _ := NewCity("Test")
	citySameName, _ := NewCity("Test")

	err := repo.Add(city)
	assert.Nil(err)

	err = repo.Add(citySameName)
	assert.NotNil(err)
}

func TestGetAll(t *testing.T) {
	assert := assert.New(t)

	repo := NewRepository()
	london, _ := NewCity("London")
	newYork, _ := NewCity("NewYork")

	repo.Add(london)
	repo.Add(newYork)

	cities := repo.GetAll()
	assert.Contains(cities, london)
	assert.Contains(cities, newYork)
}

func TestRepository_Remove(t *testing.T) {
	assert := assert.New(t)

	repo := NewRepository()
	city, _ := NewCity("Test")

	repo.Add(city)
	repo.Remove(city)

	assert.Empty(repo.GetAll())
}

func TestRepository_RemoveNonAddedCity(t *testing.T) {
	assert := assert.New(t)

	city, _ := NewCity("Test")

	repo := NewRepository()
	repo.Remove(city)

	assert.Empty(repo.GetAll())
}
