package main

import "fmt"

func main() {
	grade := "F"
	switch grade {
	case "A", "B":
		fmt.Println("합격")
	case "C", "D":
		fmt.Println("재시험")
	default:
		fmt.Println("불합격")
	}
}
