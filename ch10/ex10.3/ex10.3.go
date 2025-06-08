package main

import "fmt"

func main() {
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2차원 배열:")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d", matrix[i][j])
		}
		fmt.Println()
	}
}
