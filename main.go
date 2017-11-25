package main

import (
	"encoding/json"
	"flag"
	"github.com/spencer-p/traffic"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// Get input/output files
	inputFileName := flag.String("input", "", "Input JSON file")
	outputFileName := flag.String("output", "", "File to output JSON agent history")

	flag.Parse()

	// TODO seed rand, or have it as a flag

	// Check flags are valid
	if *inputFileName == "" || *outputFileName == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Open input file
	input, err := os.Open(*inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	// Decode the input into a struct
	dec := json.NewDecoder(input)
	var data Input
	err = dec.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	sim := traffic.NewSimulation()

	// Add bus routes
	for _, e := range data.BusRouteInfo {
		e := e // New reference
		sim.AddEdge(&e)
	}

	// Add car routes
	for _, e := range data.CarRouteInfo {
		e := e // New reference
		sim.AddEdge(&e)
	}

	for groupnum, info := range data.AgentInfo {
		// The agent info is not actually agents but info to create agents
		for i := 0; i < info.Count; i++ {
			sim.AddAgent(&Agent{
				name:        strconv.Itoa(1000*(groupnum+1) + i),
				start:       info.Start,
				destination: info.Destination,
				leaveTime:   rand.Intn(info.MaxLeaveTime)})
		}
	}

	// Run the simulation, timing it and checking for errors
	start := time.Now()
	err = sim.Simulate()
	elapsed := time.Since(start)
	if err != nil {
		log.Println("Simulation encountered error after", elapsed)
		log.Fatal(err)
	}
	log.Println("Finished simulation in", elapsed)

	// Open output to write
	output, err := os.Create(*outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	sim.PrintHistory(output)
}
