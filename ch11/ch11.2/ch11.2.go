package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	p1 := Person{name: "Bob", age: 25}
	p2 := Person{"Charlie", 40}

	fmt.Println(p1)
	fmt.Println(p2)
}
