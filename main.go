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
	}

	// Create a channel to signal when each philosopher finishes eating three times.
	doneCh := make(chan struct{}, 5)

	// Define a function for philosophers to eat.
	eatFn := func(philosopherID int) {
		for j := 0; j < 10; j++ { // Philosophers eat 3 times each
			philosophers[philosopherID].eat(philosopherID, activeForks)
		}
		doneCh <- struct{}{} // Signal that this philosopher has finished eating three times.
	}

	// Start eating for all philosophers.
	for i := 0; i < 5; i++ {
		go eatFn(i)
	}

	// Wait for all philosophers to finish eating three times.
	for i := 0; i < 5; i++ {
		<-doneCh
	}

	fmt.Println("All philosophers have eaten at least 3 times.")
	//go philosophers[1].eat(1, activeForks)
	<-activeForks
}
