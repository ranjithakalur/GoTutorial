package main

import (
	"fmt"
)

func main() {
	var x, y, result float32
	var operation string
	fmt.Printf("Go Calc\n")
	fmt.Printf("Por favor, selecione a operação desejada.\n")
	fmt.Printf("Adição (+) | Subtração (-) | Multiplicação (*) | Divisão (/):\n")
	fmt.Scan(&operation)
	fmt.Println()
	fmt.Printf("Insira o primeiro número a ser calculado:\n")
	fmt.Scan(&x)
	fmt.Printf("Insira o segundo número a ser calculado:\n")
	fmt.Scan(&y)
	fmt.Println()
	switch operation {
	case "+":
		result = x + y
		fmt.Printf("%v + %v = %v\n", x, y, result)
	case "-":
		result = x - y
		fmt.Printf("%v - %v = %v\n", x, y, result)
	case "*":
		result = x * y
		fmt.Printf("%v * %v = %v\n", x, y, result)
	case "/":
		if y == 0 {
			fmt.Printf("Errou feio, errou rude! Não se divide um número por zero.\n")
		} else {
			result = x / y
			fmt.Printf("%v / %v = %v\n", x, y, result)
		}
	default:
		fmt.Printf("Operador inválido. Por favor, reinicie o programa.\n")
	}

}
