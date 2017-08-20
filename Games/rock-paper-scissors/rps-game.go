package main

import (
	"fmt"
)

func main() {
	var vermelho, azul string
	fmt.Printf("Jogador Vermelho, escolha entre Pedra (r), Papel (p) ou Tesoura (s)\n\n")
	fmt.Scan(&vermelho)
	fmt.Printf("Jogador Azul, escolha entre Pedra (r), Papel (p) ou Tesoura (s)\n\n")
	fmt.Scan(&azul)
	switch checkWinner(vermelho, azul) {
	case 0:
		fmt.Println("Empate!")
	case 1:
		fmt.Printf("Meu coração é \033[31mvermelho\033[0m! Jogador Vermelho venceu!\n")
	case 2:
		fmt.Printf("Meu sangue é \033[34mazul\033[0m! Jogador Azul venceu!\n")
	case 3:
		fmt.Println("Opa! Você tem de escolher entre Pedra, Papel e Tesoura. Não adiantar vir com facão, amigão!")
	}

}

func checkWinner(playerRed string, playerBlue string) int8 {

	if playerRed == playerBlue {
		return 0
	}
	switch {
	case playerRed == "r" && playerBlue == "s":
		return 1
	case playerRed == "s" && playerBlue == "p":
		return 1
	case playerRed == "p" && playerBlue == "r":
		return 1
	case playerBlue == "r" && playerRed == "s":
		return 2
	case playerBlue == "s" && playerRed == "p":
		return 2
	case playerBlue == "p" && playerRed == "r":
		return 2
	default:
		var result int8 = 3
		return result
	}
}
