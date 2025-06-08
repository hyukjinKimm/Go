package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	fmt.Println("슬라이스:", slice)
	fmt.Println("길이:", len(slice))
	fmt.Println("용량:", cap(slice))
}
