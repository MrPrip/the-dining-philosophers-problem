package main

import (
	"time"
)

func main() {
	numberOfForksAndPhilos := 5

	// create array of forks and philosopher pointers
	forks := make([]*Fork, numberOfForksAndPhilos)
	philosophers := make([]*Philosopher, numberOfForksAndPhilos)

	// create array for storing and accessing channels for communication between forks and philosophers
	forkChannels := make([]chan bool, numberOfForksAndPhilos)
	philoChannelse := make([]chan bool, numberOfForksAndPhilos)

	// Populate the arrays with boolean channels
	for i := 0; i < numberOfForksAndPhilos; i++ {
		forkChannels[i] = make(chan bool)
		philoChannelse[i] = make(chan bool)
	}

	// Create forks and philosopers
	for i := 0; i < numberOfForksAndPhilos; i++ {
		forks[i] = &Fork{
			id: i,
		}
		go forks[i].table(i, forkChannels, philoChannelse)
		philosophers[i] = &Philosopher{
			id: i,
		}
		go philosophers[i].eat(i, forkChannels, philoChannelse)
	}

	time.Sleep(25 * time.Second)

}
