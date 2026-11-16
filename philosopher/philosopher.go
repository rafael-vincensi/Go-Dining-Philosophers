package philosopher

import "dining-philosophers/fork"

type Philosopher struct {
	Id        int
	LeftFork  *fork.Fork
	RightFork *fork.Fork
}
