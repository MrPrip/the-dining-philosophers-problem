package main

func main() {
	forks := make([]*Fork, 5)

	ch := make(chan []bool, 5)

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
		//go philosophers[i].eat(i, ch)

	}
	philosophers[1].eat(1, ch)
	<-ch
}
