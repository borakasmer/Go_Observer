package main

import "fmt"

type customer struct {
	id string
}

func (c *customer) update(productName string) {
	fmt.Printf("%s müşterisine %s ürünü için email gönderiliyor.\n", c.id, productName)
}

func (c *customer) getID() string {
	return c.id
}
