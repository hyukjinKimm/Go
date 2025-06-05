package main

import "fmt"

func main() {
	switch x := 7 % 2; x {
	case 0:
		fmt.Println("짝수")
	case 1:
		fmt.Println("홀수")
	}
}
