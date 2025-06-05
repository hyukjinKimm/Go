package main

import "fmt"

const Pi float64 = 3.14

// const Pi = 3.14 이것이 리터럴 상수, 자동 캐스팅
func main() {
	x := Pi
	y := Pi
	fmt.Println(x, y)
}
