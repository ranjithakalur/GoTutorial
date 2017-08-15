package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := genNumber(1, 6)
	fmt.Println(x)
}

func genNumber(minNumber int, maxNumber int) int {
	return rand.Intn(maxNumber-minNumber+1) + minNumber
}
