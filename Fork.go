package main

type Fork struct {
	id int
}

func (fork *Fork) table(id int, forks []chan bool, philos []chan bool) {
	for {
		// The current fork sends an 'true' value through its channel to signify its availability
		forks[id] <- true

		// This 'recieve' stops the fork go routine and waits for a corresponding send operation on the same channel
		<-philos[id]
	}
}
