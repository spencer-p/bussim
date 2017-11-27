package main

import (
	"github.com/spencer-p/traffic"
	"math"
)

// Car length in miles (15ft/1mi)
const CARLENGTH = 15 / 5280

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

	// Directly related to car count and distance
	// Inversely related to lane count and speed
	timeCost := (e.Distance * (e.CarCount + 1)) / (e.SpeedLimit * e.LaneCount)

	// Weigh everything together
	return a.timeWeight*timeCost + a.envWeight*e.EnvironmentalCost + a.econWeight*e.Cost
}

func (e *CarEdge) Time() int {
	// Road capacity percentage is inversely related to the speed limit.
	// 20 mph - half the road can be filled. 40 mph - 25% can be filled. etc
	roadCapacityPerc := 10 / e.SpeedLimit

	// Actual capactiy of car real estate over road real estate
	actualCapacityPerc := (CARLENGTH * e.CarCount) / (e.Distance * e.LaneCount)

	if actualCapacityPerc <= roadCapacityPerc {
		// If the cars have not exceeded capacity, travel at speed limit
		return int(e.Distance / (e.SpeedLimit / 60))
	} else if actualCapacityPerc > 1 {
		// If they have exceeded 100% road space:
		// Road is at more than capacity! Speed limit is 1
		return int(e.Distance)
	} else {
		// Road is congested. Decrease speed by exceeded capacity.

		// Create custum fn to map speed loss
		// Width is max road capacity to 100%,
		// max is real speedlimit (with 1 mph to spare)
		lostSpeedFn := quadGenerator(1-roadCapacityPerc, e.SpeedLimit-1)

		// The speed is calculated as speed limit - lost speed
		return int(e.Distance / ((e.SpeedLimit - lostSpeedFn(actualCapacityPerc-roadCapacityPerc)) / 60))
	}
}

func (e *CarEdge) RemoveAgent() {
	e.CarCount = math.Max(0, e.CarCount-1)
}

func (e *CarEdge) AddAgent() {
	e.CarCount++
}
