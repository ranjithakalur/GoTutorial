package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	switch {
	case numero < 0:
		fmt.Println("Negativo")
	case numero > 0:
		fmt.Println("Positivo")
	default:
		fmt.Println("Zero")
	}
}
