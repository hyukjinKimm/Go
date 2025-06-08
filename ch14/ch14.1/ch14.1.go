package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "멍멍!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "야옹~"
}

func saySomething(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{}
	cat := Cat{}

	saySomething(dog) // 멍멍!
	saySomething(cat) // 야옹~
}
