package aliens

// Repository is a collection of Aliens.
type Repository struct {
	aliens map[*Alien]struct{}
}

// NewRepository returns an empty initalized Repository.
func NewRepository() *Repository {
	return &Repository{
		aliens: make(map[*Alien]struct{}),
	}
}

// Add adds a new alien to the repository. If the alien already exists, it's a
// no-op.
func (r *Repository) Add(a *Alien) {
	r.aliens[a] = struct{}{}
}

// Remove removes the specified aliens from the repository. If the alien wasn't
// in the repository, it's a no-op.
func (r *Repository) Remove(aliens ...*Alien) {
	for _, a := range aliens {
		delete(r.aliens, a)
	}
}

// GetAll returns a slice with all the aliens in the repository.
func (r *Repository) GetAll() []*Alien {
	aliens := make([]*Alien, 0, len(r.aliens))
	for a := range r.aliens {
		aliens = append(aliens, a)
	}
	return aliens
}
