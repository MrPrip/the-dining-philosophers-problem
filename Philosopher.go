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

	fmt.Println(Say("thinking", id))

	for {
		// Currently the philosopher have no forks in hand
		forkInLeftHand := false
		forkInRightHand := false

		/*select {
		case left = <-forks[id]:
			break
		default:
			left = false
			break
		}

		if !left {
			counter++
			//fmt.Println("damn ---------------------- left")
			continue
		}

		//try and get right one
		select {
		case right = <-forks[((id + 1) % 5)]:
			break
		default:
			right = false
			break
		}

		if !right {
			counter++
			// then free left as well
			philos[id] <- true
			//fmt.Println("damn ---------------------- right")
			continue
		}
		*/
		//if id != 4 {

		// if the philosophers id is not equal to zero he starts by taking his left fork and then his right fork
		// if he did not get his right fork he puts the left fork back onto the table
		// if the id is zero the philosopher start by taking his right fork and then his left.
		// This is to prevent at deadlock-scenario where every philosopher would pick-up their left fork all at once.
		if id != 0 {
			forkInLeftHand = <-forks[id]
			forkInRightHand = <-forks[((id + 1) % 5)]
			// if the philosopher did not get his right fork he puts the left fork back onto the table
			if !forkInRightHand {
				philos[id] <- true
			}
		} else {
			forkInRightHand = <-forks[((id + 1) % 5)]
			forkInLeftHand = <-forks[id]
			if !forkInLeftHand {
				// if the philosopher did not get his left fork he puts the right fork back onto the table
				philos[((id + 1) % 5)] <- true
			}
		}

		// If the current philosopher have both his left and right fork he can eat
		if forkInLeftHand && forkInRightHand {
			fmt.Println(Say("eating", id))
			// He eats for one second
			time.Sleep(1 * time.Second)
			p.timesEaten++

			// If the philosopher has eaten three times he has had enough and dose not need to eat anymore
			if p.timesEaten == 3 {
				// Sends a signal to the fork go routies to indicate they can be used again.
				philos[id] <- true
				philos[((id + 1) % 5)] <- true
				fmt.Println(Say("done eating. He can't eat anymore or he will BLOW!!", id))
				return
			} else {
				// Sends a signal to the fork go routies to indicate they can be used again.
				philos[id] <- true
				philos[((id + 1) % 5)] <- true
				fmt.Println(Say("thinking", id))
			}

		}
	}
}

// A function for whenever the philosopher need to say anything about his current state
func Say(phrase string, id int) string {
	return fmt.Sprintf("Philosopher %d is %s", id, phrase)
}
