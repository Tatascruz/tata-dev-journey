package main

import "fmt"

// Busca linear: procura item por item
func buscaLinear(lista []int, alvo int) int {
	for i, num := range lista {
		if num == alvo {
			return i
		}
	}
	return -1
}

func main() {
	lista := []int{1, 3, 5, 7, 9, 11, 13}
	alvo := 7

	resultado := buscaLinear(lista, alvo)

	if resultado != -1 {
		fmt.Printf("Encontrei o número %d na posição %d\n", alvo, resultado)
	} else {
		fmt.Println("Número não encontrado")
	}
}
