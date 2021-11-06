package txt

import (
	"fmt"
	"io"
)

// Decoder reads a byte slice from an input stream and decodes it into a
// TxtUnmarshaler.
type Decoder struct {
	r io.Reader
}

// NewDecoder returns a new Decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

// TxtUnmarshaler is the interface implemented by types that can be unmarshalled
// from byte slice.
type TxtUnmarshaler interface {
	UnmarshalTxt(r io.Reader) error
}

// Decode reads the byte slice from the stream.
func (e *Decoder) Decode(t TxtUnmarshaler) error {
	err := t.UnmarshalTxt(e.r)
	if err != nil {
		return fmt.Errorf("unmarshalling from txt: %s", err)
	}

	return nil
}
