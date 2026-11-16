package simulation

import (
	"dining-philosophers/fork"
	"dining-philosophers/philosopher"
	"fmt"
	"time"
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

func (s *Simulation) RunDeadlock() {
	fmt.Println("Starting dining philosophers deadllock simulation...")

	leftForksAcquired := make(chan struct{}, philosopherCount)
	startRightFork := make(chan struct{})
	finished := make(chan struct{}, philosopherCount)

	for _, p := range s.Philosophers {
		go func(p *philosopher.Philosopher) {
			p.Think()
			p.AcquireLeftFork()

			leftForksAcquired <- struct{}{}

			<-startRightFork

			p.AcquireRightFork()
			p.Eat()

			p.ReleaseForks()
			finished <- struct{}{}
		}(p)
	}

	for i := 0; i < philosopherCount; i++ {
		<-leftForksAcquired
	}

	fmt.Println("\nAll philosophers acquired their left forks")
	fmt.Println("Allowing philosophers to acquire their right forks...\n")

	close(startRightFork)

	select {
	case <-finished:
		fmt.Println("A philosopher managed to eat")
	case <-time.After(2 * time.Second):
		fmt.Println("\nDEADLOCK DETECTED!")
		fmt.Println("All philosophers are waiting for their right forks")
	}
}
