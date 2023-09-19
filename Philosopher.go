package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	id         int
	timesEaten int
}

func (p *Philosopher) eat(id int, forks []chan bool, philos []chan bool) {
	//rand.Seed(time.Now().UnixNano())
	//timeToWait := rand.Intn(5)

	fmt.Println(Say("thinking", id))

	for {
		var left bool
		var right bool

		if id%2 == 0 {
			left = <-forks[id]
			right = <-forks[((id + 1) % 5)]

			if !right {
				philos[id] <- true
			}
		} else {
			right = <-forks[((id + 1) % 5)]
			left = <-forks[id]
			if !left {
				philos[id] <- true
			}
		}

		if left && right {
			fmt.Println(Say("eating", id))
			//time.Sleep(time.Duration(timeToWait) * time.Second)
			time.Sleep(1 * time.Second)
			p.timesEaten++
			if p.timesEaten == 3 {
				fmt.Printf("Philo %d was done eating!!!!!!!!!!!!!!!!!!!!!!!!\n", id)
				philos[id] <- true
				philos[((id + 1) % 5)] <- true
				return
			} else {
				fmt.Println(Say("thinking", id))

			}
			philos[id] <- true
			philos[((id + 1) % 5)] <- true
		}

	}

}

func Say(phrase string, id int) string {
	return fmt.Sprintf("Philosopher %d is %s", id, phrase)
}
