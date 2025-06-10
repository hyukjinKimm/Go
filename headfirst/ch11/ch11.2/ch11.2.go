package main

import "fmt"

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

func main() {
	var err error
	err = ComedyError("hello")
	fmt.Println(err)
}
