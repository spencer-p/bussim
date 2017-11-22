package main

type Input struct {
	RouteInfo []Edge
	AgentInfo []struct {
		Start        string
		Destination  string
		Count        int
		MaxLeaveTime int
	}
}
