package cities

import "errors"

// Repository is an in-memory repository of cities, uniquely identified by
// their name.
type Repository struct {
	cities map[string]*City
}

// NewRepository returns an initialized CitiesRepository.
func NewRepository() *Repository {
	return &Repository{
		cities: make(map[string]*City),
	}
}

// Has returns true if the city is present in the repository.
func (r *Repository) Has(name string) bool {
	return r.cities[name] != nil
}

// Get returns the city with the given name from the repository, or `nil` if the
// city is not present.
func (r *Repository) Get(name string) *City {
	return r.cities[name]
}

// GetAll returns a slice of all cities in the repository.
func (r *Repository) GetAll() []*City {
	cities := make([]*City, 0, len(r.cities))
	for _, c := range r.cities {
		cities = append(cities, c)
	}

	return cities
}

var errAlreadyExisting = errors.New("a city with this name already exists")

// Add adds a city to the repository.
func (r *Repository) Add(city *City) error {
	if r.cities[city.Name] != nil {
		return errAlreadyExisting
	}

	r.cities[city.Name] = city
	return nil
}

// Remove removes a city from the repository. If the city was not present, this
// is a no-op.
func (r *Repository) Remove(city *City) {
	delete(r.cities, city.Name)
}
