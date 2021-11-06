package txtcities

import (
	"errors"
	"strings"
	"testing"

	"github.com/Pitasi/alien-invasion/cities"
	"github.com/Pitasi/alien-invasion/encoding/txt"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshaler(t *testing.T) {
	assert := assert.New(t)

	content := `London north=NewYork
NewYork south=London`
	repo := cities.NewRepository()

	d := txt.NewDecoder(strings.NewReader(content))
	err := d.Decode(&RepositoryUnmarshaler{repo})

	assert.Nil(err)
	assert.Len(repo.GetAll(), 2)
}

func TestUnmarshaler_AllNeighbors(t *testing.T) {
	assert := assert.New(t)

	content := `London north=NewYork south=Chicago east=Paris west=Berlin
NewYork south=London
Chicago north=London
Paris west=London
Berlin east=London
`
	repo := cities.NewRepository()

	d := txt.NewDecoder(strings.NewReader(content))
	err := d.Decode(&RepositoryUnmarshaler{repo})

	assert.Nil(err)
	assert.Len(repo.GetAll(), 5)
}

func TestUnmarshaler_IgnoreEmptyLines(t *testing.T) {
	assert := assert.New(t)

	content := `London north=NewYork

NewYork south=London`
	repo := cities.NewRepository()

	d := txt.NewDecoder(strings.NewReader(content))
	err := d.Decode(&RepositoryUnmarshaler{repo})

	assert.Nil(err)
	assert.Len(repo.GetAll(), 2)
}

func TestUnmarshaler_IgnoreTrailingNewLine(t *testing.T) {
	assert := assert.New(t)

	content := `London north=NewYork
NewYork south=London
`
	repo := cities.NewRepository()

	d := txt.NewDecoder(strings.NewReader(content))
	err := d.Decode(&RepositoryUnmarshaler{repo})

	assert.Nil(err)
	assert.Len(repo.GetAll(), 2)
}

func TestUnmarshaler_InvalidInputs(t *testing.T) {

	table := []struct {
		content string
		message string
	}{
		{
			content: `London north=NewYork`,
			message: "missing connection",
		},
		{
			content: `London invalid=NewYork`,
			message: "invalid direction",
		},
		{
			content: "London",
			message: "city without roads",
		},
		{
			content: "London north=",
			message: "malformed neighbor",
		},
		{
			content: "London north= south=NewYork\nNewYork north=London",
			message: "malformed neighbor among valid neighbors",
		},
		{
			content: "line with too many words this should not be parsed at all",
			message: "too many words",
		},
		{
			content: `London north=NewYork north=Chicago
NewYork south=London
Chicago south=London`,
			message: "multiple roads in same direction",
		},
		{
			content: `London=NewYork north=NewYork north=Chicago`,
			message: "equal in city name",
		},
		{
			content: `London=NewYork north=New=York north=Chicago`,
			message: "equal in neighbor name",
		},
	}

	for _, test := range table {
		t.Run(test.message, func(t *testing.T) {
			assert := assert.New(t)

			repo := cities.NewRepository()
			d := txt.NewDecoder(strings.NewReader(test.content))
			err := d.Decode(&RepositoryUnmarshaler{repo})

			assert.NotNil(err, "expected fail because: "+test.message)
		})
	}
}

type MockReaderError struct{}

func (r *MockReaderError) Read([]byte) (int, error) {
	return 0, errors.New("mock error")
}

func TestUnmarshaler_Reader_Error(t *testing.T) {
	assert := assert.New(t)

	repo := cities.NewRepository()

	unmarsh := &RepositoryUnmarshaler{repo}
	err := unmarsh.UnmarshalTxt(&MockReaderError{})

	assert.NotNil(err)
}
