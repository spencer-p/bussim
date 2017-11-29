package main

import (
	"github.com/spencer-p/traffic"
)

// The null edge provides a costless transport between nodes.
type NullEdge struct {
	ToSt, FromSt string
}

func (e *NullEdge) To() string {
	return e.ToSt
}

func (e *NullEdge) From() string {
	return e.FromSt
}

func (e *NullEdge) Weight(_ traffic.Agent) float64 {
	return 0
}

func (e *NullEdge) Time() int {
	return 0
}

func (e *NullEdge) RemoveAgent() {
	return
}

func (e *NullEdge) AddAgent() {
	return
}
