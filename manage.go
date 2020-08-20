package main

type manage interface {
	register(Observer observer)
	registerList(Observer []observer)
	unregister(Observer observer)
	notifyAll()
}
