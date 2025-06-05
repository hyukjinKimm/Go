package main

import "fmt"

func main() {
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer
			}
			fmt.Println(i, j)
		}
	}
}
