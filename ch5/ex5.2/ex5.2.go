package main

import "fmt"

func main() {
  var a = 123
  var b = 456
  var c = 123456789

  fmt.Printf("%5d, %5d\n", a, b) // ① 최소 너비보다 짧은 값 너비 지정
  fmt.Printf("%05d, %05d\n", a, b) // ② 최소 너비보다 짧은 값 0 채우기
  fmt.Printf("%-5d, %-05d\n", a, b) // ③ 최소 너비보다 짧은 값 왼쪽 정렬

  fmt.Printf("%5d, %5d\n", c, c) // ④ 최소 너비보다 긴 값 너비 지정
  fmt.Printf("%05d, %05d\n", c, c) // ⑤ 최소 너비보다 긴 값 0 채우기
  fmt.Printf("%-5d, %-05d\n", c, c) // ⑥ 최소 너비보다 긴 값 왼쪽 정렬
}