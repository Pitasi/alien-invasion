package txt

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDecoder(t *testing.T) {
	assert := assert.New(t)
	encoder := NewEncoder(new(strings.Builder))
	assert.NotNil(encoder)
}

// MockTxtUnmarshaler is a mock implementation of TxtUnmarshaler that returns
// the string read from the input stream.
type MockTxtUnmarshaler struct {
	Result string
}

func (m *MockTxtUnmarshaler) UnmarshalTxt(r io.Reader) error {
	bytes, _ := io.ReadAll(r)
	m.Result = string(bytes)
	return nil
}

func TestDecode(t *testing.T) {
	assert := assert.New(t)

	mock := new(MockTxtUnmarshaler)

	decoder := NewDecoder(strings.NewReader("some string"))
	err := decoder.Decode(mock)

	assert.Nil(err)
	assert.Equal("some string", mock.Result)
}

// MockTxtErrorUnmarshaler is a mock implementation of TxtUnmarshaler that always
// returns an error.
type MockTxtErrorUnmarshaler struct{}

func (m *MockTxtErrorUnmarshaler) UnmarshalTxt(r io.Reader) error {
	return errors.New("mock error")
}

func TestDecode_Errors(t *testing.T) {
	assert := assert.New(t)

	mock := new(MockTxtErrorUnmarshaler)

	decoder := NewDecoder(strings.NewReader("some string"))
	err := decoder.Decode(mock)

	assert.NotNil(err)
}
