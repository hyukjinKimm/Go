package main

import "fmt"

func main() {
	s := "Go언어"
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d]=%x\n", i, s[i])
	}
}
