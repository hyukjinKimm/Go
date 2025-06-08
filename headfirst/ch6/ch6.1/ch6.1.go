package main

import "fmt"

func main() {
	// 길이 2, 용량 5인 슬라이스 생성 (여유 용량 있음)
	s := make([]int, 2, 5)
	s[0], s[1] = 10, 20

	// 첫 요소 주소 출력
	fmt.Printf("Before append: %p\n", &s[0])

	// 용량이 충분해서 새 배열 할당 없이 append 됨
	s = append(s, 30)

	// 첫 요소 주소 다시 출력
	fmt.Printf("After append: %p\n", &s[0])

	fmt.Println("Slice:", s)
}
