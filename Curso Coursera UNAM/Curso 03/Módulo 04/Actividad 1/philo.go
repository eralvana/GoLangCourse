package main

import (
	"fmt"
	"time"

	"sync"
)

var eatWG sync.WaitGroup

type chopstick struct{ sync.Mutex }

type philosopher struct {
	name            int
	leftCS, rightCS *chopstick
}

func (p philosopher) eat() {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		say("starting to eat", p.name)
		time.Sleep(time.Second)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		say("finishing eating", p.name)
		time.Sleep(time.Second)
	}
	eatWG.Done()
}

func say(order string, name int) {
	fmt.Printf("Philosopher #%d is %s\n", name+1, order)
}

func main() {
	number := 5

	chopsticks := make([]*chopstick, number)
	for i := 0; i < number; i++ {
		chopsticks[i] = new(chopstick)
	}

	philosophers := make([]*philosopher, number)
	for i := 0; i < number; i++ {
		philosophers[i] = &philosopher{
			name: i, leftCS: chopsticks[i], rightCS: chopsticks[(i+1)%number]}
		eatWG.Add(1)
		go philosophers[i].eat()

	}
	eatWG.Wait()
}
