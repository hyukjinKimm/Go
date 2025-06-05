package main

import "fmt"

func main() {

	grade := "B"
	switch grade {
	case "A":
		fmt.Println("우수")
	case "B":
		fmt.Println("양호")
	case "C":
		fmt.Println("보통")
	default:
		fmt.Println("재시험")

	}

}
