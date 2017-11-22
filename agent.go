package main

type Agent struct {
	name               string
	start, destination string
	leaveTime          int
}

func (a *Agent) Id() string {
	return a.name
}

func (a *Agent) Start() string {
	return a.start
}

func (a *Agent) Destination() string {
	return a.destination
}

func (a *Agent) LeaveTime() int {
	return a.leaveTime
}
