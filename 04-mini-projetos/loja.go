package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Estrutura para representar um produto
type Produto struct {
	Nome  string
	Preco float64
}

func main() {
	// Lista de produtos disponíveis na loja
	produtos := []Produto{
		{"Caixa Decorada", 59.90},
		{"Porta Joias", 89.90},
		{"Necessaire", 39.90},
		{"Porta-Retrato", 45.00},
	}

	// Carrinho de compras (slice vazio no início)
	var carrinho []Produto
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Bem-vindo á Loja Lumi")

	for {
		fmt.Println("\n---- MENU ----")
		fmt.Println("1 - Ver produtos disponíveis")
		fmt.Println("2 - Adicionar produto ao carrinho")
		fmt.Println("3 - Ver carrinho e total")
		fmt.Println("4 - Finalizar compra e sair")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scan(&opcao)
		reader.ReadString('\n') // limpa ENTER

		switch opcao {
		case 1:
			fmt.Println("\nProdutos disponíveis")
			for i, p := range produtos {
				fmt.Printf("%d. %s - R$ %.2f\n", i+1, p.Nome, p.Preco)
			}

		case 2:
			fmt.Print("Digite o número do produto que deseja adicionar: ")
			texto, _ := reader.ReadString('\n')
			texto = strings.TrimSpace(texto)
			indice, err := strconv.Atoi(texto)

			if err != nil || indice < 1 || indice > len(produtos) {
				fmt.Println("Produto inválido.")
			} else {
				carrinho = append(carrinho, produtos[indice-1])
				fmt.Printf("%s adicionado ao carrinho!\n", produtos[indice-1].Nome)
			}

		case 3:
			if len(carrinho) == 0 {
				fmt.Println("Seu carrinho esta vazio.")
			} else {
				fmt.Println("\n Itens no seu carrinho:")
				total := 0.0
				for i, p := range carrinho {
					fmt.Printf("%d. %s - R$ %.2f\n", i+1, p.Nome, p.Preco)
					total += p.Preco
				}
				fmt.Printf("\n Total de compra: R$ %.2f\n", total)
			}

		case 4:
			fmt.Println("Obrigado por comprar na Loja Lumi! Até a próxima!!")
			return

		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
