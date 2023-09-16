package main

import (
	"fmt"
)

type Philosopher struct {
	id                   int
	leftfFork, rightFork *Fork
}

func (p *Philosopher) eat(id int) {
	/*
		while(true) {
				leftFork = pickUP()
				rightFork = pickUP()
				Sleep.time(random)
				leftFork = putDown()
				rightFork = putDown()
			}
	*/
}

func say(phrase string, id int) string {
	return fmt.Sprintf("Philosopher %d is %s", id, phrase)
}
