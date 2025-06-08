package main

import (
	"fmt"
	"unsafe"
)

type Sample struct {
	a int32 // 4 bytes
	b int64 // 8 bytes
	c byte  // 1 byte
}

func main() {
	var s Sample
	fmt.Println("구조체 크기 (바이트):", unsafe.Sizeof(s))
}
