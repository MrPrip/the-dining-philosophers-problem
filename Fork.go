package main

type Fork struct {
	id int
}

func (fork *Fork) table(id int, forks []chan bool, philos []chan bool) {
	for {
		forks[id] <- true

		<-philos[id]
	}
}
