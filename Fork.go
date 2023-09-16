package main

type Fork struct {
	isPickedUp bool
}

func pickUp(fork *Fork) bool {
	fork.isPickedUp = true
	return false
}

func putDown(fork *Fork) bool {
	fork.isPickedUp = false
	return false
}
