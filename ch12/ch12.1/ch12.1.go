package main

import "fmt"

func main() {
	var x int = 10
	var p *int = &x

	fmt.Println("x의 값: ", x)
	fmt.Println("x의 주소: ", &x)
	fmt.Println("p가 가리키는 값:", *p)
}
