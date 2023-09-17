package main

type Fork struct {
	isPickedUp bool
	id         int
}

func (fork *Fork) OnTable() {

}

func (fork *Fork) pickUp(c chan []bool) bool {

	fork.isPickedUp = true

	v := <-c
	v[fork.id] = true
	c <- v

	return false
}

func (fork *Fork) putDown(c chan []bool) bool {

	fork.isPickedUp = false

	v := <-c
	v[fork.id] = false
	c <- v

	return false
}
