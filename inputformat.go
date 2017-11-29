package main

type Input struct {
	BusRouteInfo  []BusEdge
	CarRouteInfo  []CarEdge
	NullRouteInfo []NullEdge
	AgentInfo     []struct {
		Start                    string
		Destination              string
		Count                    int
		MaxLeaveTime             int
		EnvironmentalWeightRange [2]int
		CostWeightRange          [2]int
		TimeWeightRange          [2]int
	}
}
