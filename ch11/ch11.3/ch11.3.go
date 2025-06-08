package main

import "fmt"

type Address struct {
	city string
	zip  string
}

type Employee struct {
	name    string
	age     int
	address Address
}

func main() {
	e := Employee{
		name: "Dave",
		age:  35,
		address: Address{
			city: "Seoul",
			zip:  "12345",
		},
	}

	fmt.Println(e)
	fmt.Println("도시: ", e.address.city)
}
