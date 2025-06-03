package main

import "fmt"

func main() {
 var a int = 10 // ① a 변수 선언
 var msg string = "Hello Variable" // ② msg 변수 선언

 a = 20 // ③ a값 변경
 msg = "Good Morning" // ④ msg값 변경
 fmt.Println(msg, a) // ⑤ msg와 a값 출력
}