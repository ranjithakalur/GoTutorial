package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	switch numero {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("Um")
	case 2:
		fmt.Println("Dois")
	case 3:
		fmt.Println("Três")
		fallthrough
	case 4, 5:
		fmt.Println("Número tá grande...")
	default:
		fmt.Println("Não sei o número")
	}
}
