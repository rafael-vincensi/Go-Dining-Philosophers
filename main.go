package main

import (
	"dining-philosophers/simulation"
	"fmt"
)

func main() {
	fmt.Println("\nSCENARIO 1 - DEADLOCK DEMONSTRATION")

	deadlockSimulation := simulation.NewSimulation()
	deadlockSimulation.RunDeadlock()

	fmt.Println("\nSCENARIO 2 - DEADLOCK PREVENTION")

	solutionSimulation := simulation.NewSimulation()
	solutionSimulation.RunSolution()
}
