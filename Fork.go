package main

type Fork struct {
	isPickedUp bool
	id         int
}

func (fork *Fork) pickUp(ch chan []bool) {

	fork.isPickedUp = true
	v := <-ch

	tmp := make([]bool, 5)
	copy(v, tmp)
	tmp[fork.id] = true
	ch <- tmp
}

func (fork *Fork) putDown(ch chan []bool) {

	fork.isPickedUp = false

	v := <-ch

	tmp := make([]bool, 5)
	copy(v, tmp)
	tmp[fork.id] = false
	ch <- tmp
}
