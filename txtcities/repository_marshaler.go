package txtcities

import (
	"io"
	"strings"

	"github.com/Pitasi/alien-invasion/cities"
	"github.com/Pitasi/alien-invasion/encoding/txt"
)

// RepositoryMarshaler is a wrapper over cities.Repository that implements the
// Marshaler interface.
type RepositoryMarshaler struct {
	*cities.Repository
}

// MarshalTxt returns the marshaled representation of the repository.
func (r *RepositoryMarshaler) MarshalTxt() ([]byte, error) {
	builder := new(strings.Builder)

	err := r.MarshalTxtToWriter(builder)
	if err != nil {
		return nil, err
	}

	return []byte(builder.String()), nil
}

// MarshalTxt writes the marshaled representation of the repository to the
// specified stream.
func (r *RepositoryMarshaler) MarshalTxtToWriter(w io.Writer) error {
	enc := txt.NewEncoder(w)

	for _, c := range r.GetAll() {
		if len(c.AvailableDirections()) == 0 {
			continue
		}

		err := enc.Encode(&CityMarshaler{c})
		if err != nil {
			return err
		}

		_, err = w.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}

	return nil
}
