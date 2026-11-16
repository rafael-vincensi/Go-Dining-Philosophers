package simulation

import (
	"dining-philosophers/fork"
	"dining-philosophers/philosopher"
)

const philosopherCount = 5

type Simulation struct {
	Philosophers []*philosopher.Philosopher
	Forks        []*fork.Fork
}

func NewSimulation() *Simulation {
	simulation := &Simulation{
		Philosophers: make([]*philosopher.Philosopher, philosopherCount),
		Forks:        make([]*fork.Fork, philosopherCount),
	}

	for i := 0; i < philosopherCount; i++ {
		simulation.Forks[i] = &fork.Fork{
			ID: i + 1,
		}
	}

	for i := 0; i < philosopherCount; i++ {
		simulation.Philosophers[i] = &philosopher.Philosopher{
			Id:        i + 1,
			LeftFork:  simulation.Forks[i],
			RightFork: simulation.Forks[(i+1)%philosopherCount],
			State:     philosopher.Thinking,
		}
	}

	return simulation
}
