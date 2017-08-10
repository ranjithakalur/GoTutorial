package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n1 := rand.Intn(60) + 1
	n2 := rand.Intn(60) + 1
	n3 := rand.Intn(60) + 1
	n4 := rand.Intn(60) + 1
	n5 := rand.Intn(60) + 1
	n6 := rand.Intn(60) + 1
	fmt.Printf("Os teus números da sorte são:\n\n[ %v, %v, %v, %v, %v, %v ]\n\n", n1, n2, n3, n4, n5, n6)
}
