package main

func main() {
	forks := make([]*Fork, 5)
	ch := make(chan bool)
	isAvailable := make([]chan bool, 5)
	isNotAvailable := make([]chan bool, 5)
	for i := 0; i < 5; i++ {
		isAvailable[i] = make(chan bool)
		isNotAvailable[i] = make(chan bool)
	}

	for i := 0; i < 5; i++ {
		forks[i] = &Fork{
			id: i,
		}
		go fork.OnTable()
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
		go philosophers[i].eat(i, isAvailable)

	}
	<-ch
}
