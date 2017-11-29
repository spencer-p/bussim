package main

import (
	"github.com/spencer-p/traffic"
	"math"
)

const (
	// Car length in miles (15ft/1mi)
	CARLENGTH = 15.0 / 5280.0

	// Capacity k value - calculated from 300ft stop at 60mph
	KCAPACITY = 3
)

type CarEdge struct {
	ToIntersection, FromIntersection string
	SpeedLimit, Distance             float64
	LaneCount, CarCount              float64
	EnvironmentalCost, Cost          float64
}

func (e *CarEdge) To() string {
	return e.ToIntersection
}

func (e *CarEdge) From() string {
	return e.FromIntersection
}

func (e *CarEdge) Weight(ta traffic.Agent) float64 {
	a, ok := ta.(*Agent)
	if !ok {
		// This should never happen
		return math.Inf(0)
	}

	// Weigh everything together
	return a.timeWeight*float64(e.Time()) + a.envWeight*e.EnvironmentalCost + a.econWeight*e.Cost
}

func (e *CarEdge) Time() int {
	// Road capacity percentage is inversely related to the speed limit.
	roadCapacityPerc := KCAPACITY / e.SpeedLimit

	// Actual capactiy of car real estate over road real estate
	actualCapacityPerc := (CARLENGTH * e.CarCount) / (e.Distance * e.LaneCount)

	// Max capacity will be the optimal road capacity, doubled
	maxCapacityPerc := 2 * roadCapacityPerc

	minTime := int(e.Distance / (e.SpeedLimit / 60))
	maxTime := int(e.Distance / (1.0 / 60.0))

	if actualCapacityPerc <= roadCapacityPerc {
		// If the cars have not exceeded capacity, travel at speed limit
		return minTime
	} else if actualCapacityPerc > maxCapacityPerc {
		// If they have exceeded 100% road space:
		// Road is at more than capacity! Speed limit is 1
		return maxTime
	} else {
		// Road is congested. Decrease speed by exceeded capacity.
		// Create quadratic fn to map extra capacity to extra time
		speedFn := quadGenerator(roadCapacityPerc, float64(maxTime-minTime))
		return minTime + int(speedFn(actualCapacityPerc-roadCapacityPerc))
	}
}

func (e *CarEdge) RemoveAgent() {
	e.CarCount = math.Max(0, e.CarCount-1)
}

func (e *CarEdge) AddAgent() {
	e.CarCount++
}
