package txt

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEncoder(t *testing.T) {
	assert := assert.New(t)
	encoder := NewEncoder(new(strings.Builder))
	assert.NotNil(encoder)
}

// MockTxtMarshaller is a mock implementation of TxtMarshaller that returns the
// string used to initialize it.
type MockTxtMarshaler struct {
	s string
}

func NewMockTxtMarshaler(s string) *MockTxtMarshaler {
	return &MockTxtMarshaler{s}
}

func (m *MockTxtMarshaler) MarshalTxt() ([]byte, error) {
	return []byte(m.s), nil
}

func TestEncode(t *testing.T) {
	assert := assert.New(t)

	mock := NewMockTxtMarshaler("test string")

	b := new(strings.Builder)
	encoder := NewEncoder(b)

	err := encoder.Encode(mock)

	assert.Nil(err)
	assert.Equal("test string", b.String())
}

// MockTxtErrorMarshaler is a mock implementation of TxtMarshaler that always
// returns an error.
type MockTxtErrorMarshaler struct{}

func (m *MockTxtErrorMarshaler) MarshalTxt() ([]byte, error) {
	return nil, errors.New("mock error")
}

func TestEncode_ErrorMarshaling(t *testing.T) {
	assert := assert.New(t)

	mock := new(MockTxtErrorMarshaler)

	w := new(strings.Builder)
	encoder := NewEncoder(w)
	err := encoder.Encode(mock)

	assert.NotNil(err)
}

type MockErrorWriter struct{}

func (w *MockErrorWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("mock error")
}

func TestEncode_ErrorWriting(t *testing.T) {
	assert := assert.New(t)

	mock := NewMockTxtMarshaler("test string")

	w := new(MockErrorWriter)
	encoder := NewEncoder(w)
	err := encoder.Encode(mock)

	assert.NotNil(err)
}
