package main

import (
	"math"
)

type BusEdge struct {
	ToStop, FromStop string
	Distance         float64
	Capacity         float64
	Speed            float64
	Waiting          float64
	VehicleCount     float64
}

func (e *BusEdge) To() string {
	return e.ToStop
}

func (e *BusEdge) From() string {
	return e.FromStop
}

func (e *BusEdge) Weight() float64 {
	// The weight is directly related to # of waiting and distance and inversely
	// related to the capacity, number of buses, and bus speed
	return (e.Distance * (e.Waiting + 1)) / (e.Capacity * e.Speed * e.VehicleCount)
}

func (e *Edge) Time() int {
	// Base speed is distance/speed - number of vehicles cannot lower this
	// Speed has to be reduced to minutes (div by 60)
	// Extra speed is # of extra people times the base time
	base := (e.Distance / (e.Speed / 60))
	var extra float64
	if e.Waiting > e.Capacity*e.VehicleCount {
		extra = (e.Waiting - e.Capacity*e.VehicleCount) * base
	} else {
		extra = 0
	}
	return int(base + extra)
}

func (e *BusEdge) RemoveAgent() {
	e.Waiting = math.Max(0, e.Waiting-1)
}

func (e *BusEdge) AddAgent() {
	e.Waiting++
}
