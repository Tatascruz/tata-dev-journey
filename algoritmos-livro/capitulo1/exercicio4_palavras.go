package main

import "fmt"

// Pesquisa binária para palavras (com debug)
func pesquisaBinariaPalavras(lista []string, alvo string) (int, int) {
	esquerda := 0
	direita := len(lista) - 1
	tentativas := 0

	fmt.Println("Lista completa:")
	for i, nome := range lista {
		fmt.Printf("%d: [%s]\n", i, nome)
	}
	fmt.Println("---------")

	for esquerda <= direita {
		tentativas++
		meio := (esquerda + direita) / 2
		chute := lista[meio]

		fmt.Printf("Tentativa %d → meio=%d, chute=[%s]\n", tentativas, meio, chute)

		if chute == alvo {
			return meio, tentativas
		}

		if chute > alvo {
			fmt.Println("Resultado: chute > alvo → vou para metade da ESQUERDA")
			direita = meio - 1
		} else {
			fmt.Println("Resultado: chute < alvo → vou para metade da DIREITA")
			esquerda = meio + 1
		}
	}

	return -1, tentativas
}

func main() {
	lista := []string{
		"Ana", "Bia", "Carla", "Denise", "Elaine", "Fernanda", "Gi",
		"Helena", "Isabela", "Julia", "Karen", "Larissa", "Maria",
		"Nadia", "Olivia", "Patricia", "Rafaela", "Sofia", "Tata", "Yara",
	}

	alvo := "Tata"

	indice, tentativas := pesquisaBinariaPalavras(lista, alvo)

	if indice != -1 {
		fmt.Printf("\n'%s' encontrada na posição %d em %d tentativas.\n", alvo, indice, tentativas)
	} else {
		fmt.Printf("\n'%s' não foi encontrada após %d tentativas.\n", alvo, tentativas)
	}
}
