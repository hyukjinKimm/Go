package main

import "fmt"

const (
	Red = iota
	Green
	Blue
)

func main() {
	r := Red
	g := Green
	b := Blue

	fmt.Println(r, g, b)
}
