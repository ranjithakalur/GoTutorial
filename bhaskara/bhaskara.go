package main

// Importar os pacotes necessários
import (
	"fmt"
	"math"
)

var a, b, c, d, raizd, x1, x2 float64

func main() {

	fmt.Printf("Olá! Seja bem vindo a calculadora de equações de segundo grau.\n")
	// Inserir os valores numéricos dos coeficientes a, b e c.
	fmt.Printf("Insira o valor de A:\n")
	fmt.Scan(&a)
	fmt.Printf("Insira o valor de B:\n")
	fmt.Scan(&b)
	fmt.Printf("Insira o valor de C:\n")
	fmt.Scan(&c)
	calculation()
}

func calculation() {

	// Avaliar se o coeficiente A = 0
	if a == 0 {
		fmt.Printf("Se o coeficiente A for igual a zero, a parábola não poderá ser calculada.\n")
		fmt.Printf("Por favor, reinsira os valores dos coeficientes.\n")
		return
	}

	// Calcular o valor de delta
	d := math.Pow(b, 2) - 4*a*c
	raizd := math.Sqrt(d)
	fmt.Printf("O valor de Delta é: %v\n", d)

	// Verificar se o valor de Delta é menor que zero (ou seja, negativo)
	if d < 0 {
		fmt.Printf("Outch! :-( ~ Infelizmente o valor de Delta é negativo. Por favor, reinsira os dados.\n")
		return
	}

	/*
		Verificar se o valor de Delta é igual a zero (assim se tem apenas uma raiz)
		Aplicar a fórmula de Bhaskara
		x = -b ± √Δ / 2*a
	*/
	if d == 0 {
		x1 = (-b + d) / (2 * a)
		fmt.Printf("O valor de X é: %v\n", x1)
	} else {
		/*
			Verificar se o valor de Delta é maior que zero (assim se tem duas raízes)
			Aplicar a fórmula de Bhaskara
			x = -b ± √Δ / 2*a
		*/
		x1 = (-b + raizd) / (2 * a)
		x2 = (-b - raizd) / (2 * a)
		fmt.Printf("O valor de X' é: %v\n", x1)
		fmt.Printf("O valor de X\" é: %v \n", x2)
	}
}
