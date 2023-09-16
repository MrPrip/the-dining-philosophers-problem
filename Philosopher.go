package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	id                  int
	leftFork, rightFork *Fork
}

func (p *Philosopher) eat(id int, c chan []bool) {
	left := false
	right := false

	v := <-c

	fmt.Println("hej")
	fmt.Println(!v[p.leftFork.id])
	fmt.Println(!v[p.rightFork.id])

	if !v[p.leftFork.id] {
		p.leftFork.pickUp(c)
	}

	if !v[p.rightFork.id] {
		p.rightFork.pickUp(c)
	}

	if p.leftFork.isPickedUp && p.rightFork.isPickedUp {
		say("eating", p.id)
		time.Sleep(1000)
		p.leftFork.putDown(c)
		p.rightFork.putDown(c)
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
}

func say(phrase string, id int) string {
	return fmt.Sprintf("Philosopher %d is %s", id, phrase)
}
