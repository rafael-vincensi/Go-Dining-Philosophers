package philosopher

import (
	"dining-philosophers/fork"
	"fmt"
)

type State string

const (
	Thinking State = "thinking"
	Waiting  State = "waiting"
	Eating   State = "eating"
)

type Philosopher struct {
	Id        int
	LeftFork  *fork.Fork
	RightFork *fork.Fork
	State     State
}

func (p *Philosopher) Think() {
	p.State = Thinking
	fmt.Printf("Philosoper %d is thinking...\n", p.Id)
}

func (p *Philosopher) AcquireForks() {
	p.State = Waiting

	fmt.Printf("Philosopher %d is waiting for the left fork...\n", p.Id)
	p.LeftFork.Mutex.Lock()

	fmt.Printf("Philosopher %d piccked up the left fork.\n", p.Id)
	fmt.Printf("Philosopher %d is waiting for the right fork...\n", p.Id)
	p.RightFork.Mutex.Lock()

	fmt.Printf("Philosopher %d picked up the rihht fork\n", p.Id)
}

func (p *Philosopher) Eat() {
	p.State = Eating
	fmt.Printf("Philosopher %d is eating\n", p.Id)
}

func (p *Philosopher) ReleaseForks() {
	p.RightFork.Mutex.Unlock()
	p.LeftFork.Mutex.Unlock()

	p.State = Thinking
	fmt.Printf("Philosopher %d releaesed both forks\n", p.Id)
}
