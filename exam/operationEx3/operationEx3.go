package main

import (
	"fmt"
	"reflect"
)

func main() {
	var score int32 = 80

	if score >= 80 && score <= 100 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	fmt.Println(reflect.TypeOf(score))
}
