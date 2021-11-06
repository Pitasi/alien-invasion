// Package txt implements encoding and decoding of byte slices into types that
// satisfy the TxtMarshaler interface.
package txt

import (
	"fmt"
	"io"
)

// Encoder writes encoded values into an output stream.
type Encoder struct {
	w io.Writer
}

// NewEncoder returns a new Encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// TxtMarshaler is the interface implemented by types that can be marshalled
// into byte slice.
type TxtMarshaler interface {
	MarshalTxt() ([]byte, error)
}

// Encode writes the encoding of v into the stream.
func (e *Encoder) Encode(t TxtMarshaler) error {
	bytes, err := t.MarshalTxt()
	if err != nil {
		return fmt.Errorf("marshalling into txt: %s", err)
	}

	_, err = e.w.Write(bytes)
	if err != nil {
		return fmt.Errorf("writing to buffer: %s", err)
	}

	return nil
}
