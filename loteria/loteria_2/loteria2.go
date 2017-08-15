package main

import (
	"fmt"
	"time"
)

func main() {
	const n = 1000000
	m := make(map[int32]struct{}, n)
	t := time.Now()
	for i := int32(0); i < n; i++ {
		m[i] = zero
	}
	fmt.Println(time.Since(t))
	fmt.Println("len: ", len(m))
	// for k := range m {
	//		fmt.Print(k, " ")
	// }
}

var zero = struct{}{}
