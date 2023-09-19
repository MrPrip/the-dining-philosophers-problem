package main

import (
	"fmt"
	"time"
)

func main() {
	numberOfForksAndPhilos := 5
	forks := make([]*Fork, numberOfForksAndPhilos)
	philosophers := make([]*Philosopher, numberOfForksAndPhilos)
	forkChannels := make([]chan bool, numberOfForksAndPhilos)
	philoChannelse := make([]chan bool, numberOfForksAndPhilos)

	for i := 0; i < numberOfForksAndPhilos; i++ {
		forkChannels[i] = make(chan bool)
		philoChannelse[i] = make(chan bool)
	}

	for i := 0; i < numberOfForksAndPhilos; i++ {
		forks[i] = &Fork{
			id: i,
		}
		go forks[i].table(i, forkChannels, philoChannelse)
		fmt.Printf("Fork %d created \n", i)
		philosophers[i] = &Philosopher{
			id:        i,
			rightFork: forks[i],
			leftFork:  forks[(i+1)%2],
		}
		go philosophers[i].eat(i, forkChannels, philoChannelse)
		fmt.Printf("Philosoper %d created \n", i)
	}

	time.Sleep(20 * time.Second)
}
