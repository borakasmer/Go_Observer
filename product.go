package main

import (
	"fmt"
	"strings"
)

type product struct {
	observerList []observer
	name         string
	inStock      bool
}

func newProduct(name string) *product {
	return &product{name: name}
}

func (p *product) updateAvailability() {
	fmt.Printf("Product %s artık stoklardadır\n", p.name)
	fmt.Printf(strings.Repeat("-", 100))
	fmt.Printf("\n")
	p.inStock = true
	p.notifyAll()
}

func (p *product) notifyAll() {
	for _, observer := range p.observerList {
		observer.update(p.name)
	}
}

func (p *product) register(o observer) {
	p.observerList = append(p.observerList, o)
}

func (p *product) registerList(olist []observer) {

	for _, o := range olist {
		p.observerList = append(p.observerList, o)
	}
}

func (p *product) unregister(o observer) {
	p.observerList = removeFromlist(p.observerList, o)
}

func removeFromlist(observerList []observer, removeObserver observer) []observer {
	for index, observer := range observerList {
		if removeObserver.getID() == observer.getID() {
			return append(observerList[:index], observerList[index+1:]...)
		}
	}
	return observerList
}
