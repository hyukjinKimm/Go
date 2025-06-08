package main

import "fmt"

func main() {
	s := "hello"
	runes := []rune(s)
	runes[0] = 'H'
	newStr := string(runes)
	fmt.Println("newStr: ", newStr)
}
