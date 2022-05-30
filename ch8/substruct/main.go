package main

import "fmt"

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

func main() {
	var subscriber *Subscriber = new(Subscriber)
	subscriber.Name = "AAA"
	subscriber.Street = "AAA"
	subscriber.Address.Street = "BBB"

	fmt.Println(*subscriber)
}
