package main

type manage interface {
	register(Observer observer)
	unregister(Observer observer)
	notifyAll()
}
