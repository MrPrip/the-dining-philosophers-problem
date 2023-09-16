package main

func main() {
	forks := make([]*Fork, 5)

	for i := 0; i < 5; i++ {
		forks[i] = new(Fork)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id:        i,
			rightFork: forks[i],
			leftfFork: forks[(i+1)%5],
		}

		/// id = 1
		// rightfork = 1
		// leftfork = (2) % 5 = 2
		go philosophers[i].eat(i)
	}
}
