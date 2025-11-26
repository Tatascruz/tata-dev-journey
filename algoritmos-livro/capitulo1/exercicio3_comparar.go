package main

import "fmt"

// Busca linear
func buscaLinear(lista []int, alvo int) (int, int) {
	tentativas := 0
	for i, num := range lista {
		tentativas++
		if num == alvo {
			return i, tentativas
		}
	}
	return -1, tentativas
}

// Pesquisa binária
func pesquisaBinaria(lista []int, alvo int) (int, int) {
	esquerda := 0
	direita := len(lista) - 1
	tentativas := 0

	for esquerda <= direita {
		tentativas++
		meio := (esquerda + direita) / 2

		if lista[meio] == alvo {
			return meio, tentativas
		}

		if lista[meio] > alvo {
			direita = meio - 1
		} else {
			esquerda = meio + 1
		}
	}

	return -1, tentativas
}

func main() {
	// Criando lista com 5000 números
	lista := make([]int, 5000)
	for i := 0; i < 5000; i++ {
		lista[i] = i + 1
	}

	alvo := 3510

	indiceLinear, passosLinear := buscaLinear(lista, alvo)
	indiceBinario, passosBinarios := pesquisaBinaria(lista, alvo)

	fmt.Println("Resultado Busca Linear:", indiceLinear, "passos:", passosLinear)
	fmt.Println("Resultado Pesquisa Binária:", indiceBinario, "passos:", passosBinarios)
}
