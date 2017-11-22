package main

import (
	"math"
)

type Edge struct {
	ToStop, FromStop string
	Distance         float64
	Capacity         float64
	Speed            float64
	Waiting          float64
}

func (e *Edge) To() string {
	return e.ToStop
}

func (e *Edge) From() string {
	return e.FromStop
}

func (e *Edge) Weight() float64 {
	// The weight is directly related to # of waiting and distance and inversely
	// related to the capacity and bus speed
	return (e.Distance * (e.Waiting + 1)) / (e.Capacity * e.Speed)
}

func (e *Edge) Time() int {
	// Base speed is speed * distance
	// Extra speed is bus capacity divided by waiting times the base speed
	base := e.Speed * e.Distance
	var extra float64
	if e.Waiting > e.Capacity {
		extra = (e.Capacity / e.Waiting) * base
	} else {
		extra = 0
	}
	return int(base + extra)
}

func (e *Edge) RemoveAgent() {
	e.Waiting = math.Max(0, e.Waiting-1)
}

func (e *Edge) AddAgent() {
	e.Waiting++
}
