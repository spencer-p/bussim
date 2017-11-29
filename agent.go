package main

type Agent struct {
	name, group                       string
	start, destination                string
	leaveTime                         int
	timeWeight, envWeight, econWeight float64
}

func (a *Agent) Id() string {
	return a.name
}

func (a *Agent) Group() string {
	return a.group
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
