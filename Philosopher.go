package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	id                  int
	leftFork, rightFork *Fork
	timesEaten          int
}

func (p *Philosopher) eat(id int, forks []chan bool, philos []chan bool) {
	//rand.Seed(time.Now().UnixNano())
	//timeToWait := rand.Intn(5)
	for {
		var left bool
		var right bool

		left = <-forks[id]
		right = <-forks[((id + 1) % 5)]

		if !right {
			philos[id] <- true
		}

		if left && right {
			//time.Sleep(time.Duration(timeToWait) * time.Second)
			time.Sleep(1 * time.Second)
			p.timesEaten++
			philos[id] <- true
			philos[((id + 1) % 5)] <- true
		}
		if p.timesEaten == 3 {
			fmt.Printf("\nPhilo %d was done eating!!!!!!!!!!!!!!!!!!!!!!!!", id)
			return
		}
	}

}

func say(phrase string, id int) string {
	return fmt.Sprintf("Philosopher %d is %s", id, phrase)
}
