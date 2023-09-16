package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	id                  int
	leftFork, rightFork *Fork
}

func (p *Philosopher) eat(id int, ch chan []bool) {

	v := <-ch

	if !v[p.leftFork.id] {
		go p.leftFork.pickUp(ch)
		fmt.Println("left fork is picked up")
		time.Sleep(1000)
	}

	if !v[p.rightFork.id] {
		go p.rightFork.pickUp(ch)
		fmt.Println("right fork is picked up")
		time.Sleep(1000)
	}

	if p.leftFork.isPickedUp && p.rightFork.isPickedUp {
		fmt.Println(say("eating", p.id))
		time.Sleep(1000)
		go p.leftFork.putDown(ch)
		go p.rightFork.putDown(ch)
	}
	/*
		while(true) {
				leftFork = pickUp()
				rightFork = pickUp()
				Sleep.time(random)
				leftFork = putDown()
				rightFork = putDown()
			}
	*/
	time.Sleep(1000)
	ch <- v
}

func say(phrase string, id int) string {
	return fmt.Sprintf("Philosopher %d is %s", id, phrase)
}
