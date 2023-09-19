package main

func main() {
	numberOfForksAndPhilos := 5
	finishedPhilosophers := 0
	getFinishedPhilosopherRoutines := make(chan int)

	// create array of forks and philosopher pointers
	forks := make([]*Fork, numberOfForksAndPhilos)
	philosophers := make([]*Philosopher, numberOfForksAndPhilos)

	// create array for storing and accessing channels for communication between forks and philosophers
	forkChannels := make([]chan bool, numberOfForksAndPhilos)
	philoChannels := make([]chan bool, numberOfForksAndPhilos)

	// Populate the arrays with boolean channels
	for i := 0; i < numberOfForksAndPhilos; i++ {
		forkChannels[i] = make(chan bool)
		philoChannels[i] = make(chan bool)
	}

	// Create forks and philosopers
	for i := 0; i < numberOfForksAndPhilos; i++ {
		forks[i] = &Fork{
			id: i,
		}
		go forks[i].table(i, forkChannels, philoChannels)
		philosophers[i] = &Philosopher{
			id: i,
		}
		go philosophers[i].eat(i, forkChannels, philoChannels, getFinishedPhilosopherRoutines)
	}

	for {
		if finishedPhilosophers == numberOfForksAndPhilos {
			for i := 0; i < numberOfForksAndPhilos; i++ {
				//close(forkChannels[i])
				close(philoChannels[i])
			}
			return
		}
		returnValue := <-getFinishedPhilosopherRoutines
		finishedPhilosophers += returnValue
	}
}
