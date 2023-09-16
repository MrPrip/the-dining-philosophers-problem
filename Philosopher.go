package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	id                  int
	leftFork, rightFork *Fork
}

func (p *Philosopher) eat(id int, isAvailable []chan bool) {
	var left, right bool

	select {
	case left = <-isAvailable[id]:
		break
	default:
		left = false
		break
	}

	select {
	case right = <-isAvailable[(id+1)%5]:
		break
	default:
		right = false
		break
	}

	if !right {
		left = false
	}

	if left && right {
		time.Sleep(2 * time.Second)
		fmt.Println(say("I'm eating", id))

	}
}

func say(phrase string, id int) string {
	return fmt.Sprintf("Philosopher %d is %s", id, phrase)
}
