package main

import "fmt"

func start(c chan []bool) {
	v := make([]bool, 5)

	c <- v
}

func main() {
	forks := make([]*Fork, 5)

	activeForks := make(chan []bool)
	go start(activeForks)

	for i := 0; i < 5; i++ {
		forks[i] = &Fork{
			id: i,
		}
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id:        i,
			rightFork: forks[i],
			leftFork:  forks[(i+1)%5],
		}

		/// id = 1
		// rightfork = 1
		// leftfork = (2) % 5 = 2
		go philosophers[i].eat(i, activeForks)

	}
	//go philosophers[1].eat(1, activeForks)
	fmt.Println("")
	<-activeForks
}
