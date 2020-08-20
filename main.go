package main

func main() {
	ps5 := newProduct("Sony Playstation 5")
	observer1 := &customer{id: "bora@borakasmer.com"}
	observer2 := &customer{id: "martin@martinfowler.com"}
	ps5.registerList([]observer{observer1, observer2})
	ps5.updateAvailability()
}
