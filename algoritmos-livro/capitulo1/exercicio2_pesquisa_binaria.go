package main

import "fmt"

// Pesquisa binária
func pesquisaBinaria(lista []int, alvo int) int {
	esquerda := 0
	direita := len(lista) - 1

	for esquerda <= direita {
		meio := (esquerda + direita) / 2
		chute := lista[meio]

		if chute == alvo {
			return meio
		}

		if chute > alvo {
			direita = meio - 1
		} else {
			esquerda = meio + 1
		}
	}

	return -1
}

func main() {
	lista := []int{1, 3, 7, 9, 11, 13}
	alvo := 13

	resultado := pesquisaBinaria(lista, alvo)

	if resultado != -1 {
		fmt.Printf("Número %d encontrado na posição %d\n", alvo, resultado)
	} else {
		fmt.Println("Número não encontrado")
	}
}
