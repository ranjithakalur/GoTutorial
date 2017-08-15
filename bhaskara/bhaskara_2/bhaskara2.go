package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Printf("Insira o valor de A:\n")
	fmt.Scan(&a)
	fmt.Printf("Insira o valor de B:\n")
	fmt.Scan(&b)
	fmt.Printf("Insira o valor de C:\n")
	fmt.Scan(&c)
	x1, x2, valRoots := calculateBhaskara(a, b, c)
	switch valRoots {
	case 0:
		fmt.Printf("A equação não tem raízes.\n")
	case 1:
		fmt.Printf("O valor de X é: %v\n", x1)
	case 2:
		fmt.Printf("O valor de X1 é: %v\n", x1)
		fmt.Printf("O valor de X2 é: %v\n", x2)
	}

}

func calculateBhaskara(a float64, b float64, c float64) (float64, float64, int) {
	var d, raizd, x1, x2 float64
	var valRoots int
	if a == 0 {
		return 0, 0, 0
	}
	d = math.Pow(b, 2) - 4*a*c
	raizd = math.Sqrt(d)
	fmt.Printf("Delta = %v\n", d)
	switch {
	case d < 0:
		valRoots = 0
	case d == 0:
		valRoots = 1
	case d > 0:
		valRoots = 2
	}
	x1 = (-b + raizd) / (2 * a)
	x2 = (-b - raizd) / (2 * a)

	return x1, x2, valRoots
}
