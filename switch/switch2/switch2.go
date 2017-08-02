package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	switch sub := numero - 10; {
	case sub < 0:
		fmt.Println("Negativo")
	case sub > 0:
		fmt.Println("Positivo")
	default:
		fmt.Println("Zero")
	}

	/*
		Pro caso de precisar usar uma variável uma única vez ou para uma única condição.
	*/
	if par := (numero%2 == 0); par {
		fmt.Printf("O número %v é par!\n", numero)
	}

}
