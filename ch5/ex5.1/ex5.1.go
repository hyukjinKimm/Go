package main

import "fmt"

func main() {
 var a int = 10
 var b int = 20
 var f float64 = 32799438743.8297

 fmt.Print("a:", a, "b:", b) // ①
 fmt.Println("a:", a, "b:", b, "f:", f) // ②
 fmt.Printf("a: %d b: %d f:%f\n", a, b, f) // ③
}
