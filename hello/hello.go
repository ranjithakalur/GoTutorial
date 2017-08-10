package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Hello, World! Vamos ver quantos núcleos de processamento a tua máquina tem?\n")
	var x string
	fmt.Printf("Digite \"Sim\", para prosseguir.\n")
	fmt.Scan(&x)
	if x == "sim" {
		n := runtime.NumCPU()
		fmt.Printf("Este computador tem %v cores.\n", n)
	} else {
		fmt.Printf("Que peninha! Adeus! (Chato do caralho!)\n")
	}
}
