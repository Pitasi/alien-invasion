package main

import (
	"bytes"
	"flag"
)

type Configuration struct {
	MapFilePath                string
	MaximumAlienMoves          int
	AliensCount                int
	HowManyAliensToDestroyCity int
}

func ParseFlags(progname string, arguments []string) (Configuration, string, error) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	var buf bytes.Buffer
	flags.SetOutput(&buf)

	s := flags.String("path", "sample.txt", "filepath containing the map")
	moves := flags.Int("alienmoves", 10000, "maximum number of times an alien can move")
	aliensCount := flags.Int("aliencount", 10, "number of aliens to spawn")
	aliensDestroyCity := flags.Int("aliensdestroycity", 2, "number of aliens required to destroy a city")

	err := flags.Parse(arguments)
	if err != nil {
		return Configuration{}, buf.String(), err
	}

	return Configuration{
		MapFilePath:                *s,
		MaximumAlienMoves:          *moves,
		AliensCount:                *aliensCount,
		HowManyAliensToDestroyCity: *aliensDestroyCity,
	}, buf.String(), nil
}
