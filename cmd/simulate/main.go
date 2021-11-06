// Program simulation runs a alien invasion simulation.
//
// It reads the world shape from a file, prints the city being destroyed and the
// final world shape to the stdout.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Pitasi/alien-invasion/aliens"
	"github.com/Pitasi/alien-invasion/encoding/txt"
	"github.com/Pitasi/alien-invasion/simulation"
	"github.com/Pitasi/alien-invasion/spawner"
	"github.com/Pitasi/alien-invasion/world"
)

func main() {
	rand.Seed(time.Now().Unix())

	config, out, err := ParseFlags(os.Args[0], os.Args[1:])
	if err != nil {
		fmt.Println(out)
		os.Exit(2)
	}

	// create a new world
	w := world.New()

	// setup cities
	initCities(config.MapFilePath, w)

	// setup aliens
	alienConfig := spawner.AlienTemplate{
		MaximumMoves: config.MaximumAlienMoves,
		MovePolicy:   aliens.RandomMovePolicy,
	}
	spawner, err := spawner.New(alienConfig, w, spawner.ChooseRandomCity)
	check(err)

	err = spawner.Spawn(config.AliensCount)
	check(err)

	// prepare simulation
	events := make(chan simulation.Event)
	sim, err := simulation.New(
		w,
		config.HowManyAliensToDestroyCity,
		events,
	)
	check(err)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		// handle simulation events
		defer wg.Done()

		for event := range events {
			switch e := event.(type) {
			case *simulation.CityDestroyedEvent:
				handleCityDestroyed(e)
			}
		}
	}()

	go func() {
		// run simulation
		defer wg.Done()

		err := sim.Run()
		check(err)
	}()

	// wait for the simulation to finish and print results
	wg.Wait()
	printOutput(w)
}

func initCities(mapFilePath string, w *world.World) {
	f, err := os.Open(mapFilePath)
	check(err)

	dec := txt.NewDecoder(f)
	err = dec.Decode(w)
	check(err)
}

func printOutput(w *world.World) {
	e := txt.NewEncoder(os.Stdout)
	err := e.Encode(w)
	check(err)
}

func handleCityDestroyed(e *simulation.CityDestroyedEvent) {
	names := make([]string, 0, len(e.DestroyedAliens))
	for _, a := range e.DestroyedAliens {
		names = append(names, a.Name)
	}

	fmt.Printf("%s has been destroyed by aliens %s\n", e.City.Name, strings.Join(names, ", "))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
