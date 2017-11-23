package main

type Input struct {
	BusRouteInfo []BusEdge
	CarRouteInfo []CarEdge
	AgentInfo    []struct {
		Start        string
		Destination  string
		Count        int
		MaxLeaveTime int
	}
}
