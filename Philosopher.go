package main

import (
	"fmt"
)

type Philosopher struct {
	id                  int
	leftFork, rightFork *Fork
	timesEaten          int
}

func (p *Philosopher) eat(id int, ch chan []bool) {

	v := <-ch

	fmt.Println(say("is thinking", id))

	if !v[p.leftFork.id] {
		go p.leftFork.pickUp(ch)
	}

	if !v[p.rightFork.id] {
		go p.rightFork.pickUp(ch)
	} else {
		go p.leftFork.putDown(ch)
	}

	if p.leftFork.isPickedUp && p.rightFork.isPickedUp {
		fmt.Println(say("starts eating", id))
		go p.leftFork.putDown(ch)
		go p.rightFork.putDown(ch)
		fmt.Println(say("is done eating", id))
		p.timesEaten += 1
		fmt.Print("Philosopher ")
		fmt.Print(id)
		fmt.Print(" has eaten ")
		fmt.Print(p.timesEaten)
		fmt.Println(" times")

	}
	ch <- v
}

func say(phrase string, id int) string {
	return fmt.Sprintf("Philosopher %d  %s", id, phrase)
}
