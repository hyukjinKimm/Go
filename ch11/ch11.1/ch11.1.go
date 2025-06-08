package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var p Person
	p.name = "Alice"
	p.age = 30
	fmt.Println(p)
}
