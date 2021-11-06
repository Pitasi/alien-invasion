package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFlagsCorrect(t *testing.T) {
	var tests = []struct {
		args []string
		conf Configuration
	}{
		{
			[]string{},
			Configuration{MapFilePath: "sample.txt", MaximumAlienMoves: 10000, AliensCount: 10, HowManyAliensToDestroyCity: 2},
		},
		{
			[]string{"-path", "somepath.txt"},
			Configuration{MapFilePath: "somepath.txt", MaximumAlienMoves: 10000, AliensCount: 10, HowManyAliensToDestroyCity: 2},
		},
		{
			[]string{"-alienmoves", "42"},
			Configuration{MapFilePath: "sample.txt", MaximumAlienMoves: 42, AliensCount: 10, HowManyAliensToDestroyCity: 2},
		},
		{
			[]string{"-aliencount", "42"},
			Configuration{MapFilePath: "sample.txt", MaximumAlienMoves: 10000, AliensCount: 42, HowManyAliensToDestroyCity: 2},
		},
		{
			[]string{"-aliensdestroycity", "42"},
			Configuration{MapFilePath: "sample.txt", MaximumAlienMoves: 10000, AliensCount: 10, HowManyAliensToDestroyCity: 42},
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			assert := assert.New(t)

			conf, _, err := ParseFlags("prog", tt.args)
			assert.Nil(err)
			assert.Equal(tt.conf, conf)
		})
	}
}

func TestParseFlagsError(t *testing.T) {
	var tests = []struct {
		args []string
	}{
		{
			[]string{"-alienmoves", "notanumber"},
		},
		{
			[]string{"-alienscount", "notanumber"},
		},
		{
			[]string{"-aliensdestroycity", "notanumber"},
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			assert := assert.New(t)

			_, _, err := ParseFlags("prog", tt.args)
			assert.NotNil(err)
		})
	}
}
