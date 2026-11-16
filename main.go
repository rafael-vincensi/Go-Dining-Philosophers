package main

import "dining-philosophers/simulation"

func main() {
	sim := simulation.NewSimulation()
	sim.RunDeadlock()
}
