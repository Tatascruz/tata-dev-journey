package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"tata-dev-journey/04-mini-projetos/models"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var produtos []models.Produto

	for {
		fmt.Println("\n---- MENU ----")
		fmt.Println("1 - Cadastrar produto")
		fmt.Println("2 - Listar Produtos")
		fmt.Println("3 - Finalizar compra e sair")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scan(&opcao)
		reader.ReadString('\n') // limpa ENTER

		switch opcao {
		case 1:
			fmt.Println() //Adiciona uma linha "um espaço"
			novo := cadastrarProduto(reader, len(produtos)+1)
			produtos = append(produtos, novo)

		case 2:
			listarProdutos(produtos)

		case 3:
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}

func cadastrarProduto(reader *bufio.Reader, id int) models.Produto {
	var p models.Produto
	p.ID = id

	fmt.Print("Nome do produto: ")
	nome, _ := reader.ReadString('\n')
	p.Nome = strings.TrimSpace(nome)

	fmt.Print("Preço: ")
	fmt.Scan(&p.Preco)
	reader.ReadString('\n')

	fmt.Print("Quantidade: ")
	fmt.Scan(&p.Quantidade)
	reader.ReadString('\n')

	fmt.Println("Produto cadastrado com sucesso!")
	return p
}

func listarProdutos(lista []models.Produto) {
	if len(lista) == 0 {
		fmt.Println("Nenhum produto cadastrado.")
		return
	}

	fmt.Println("\n--- PRODUTOS CADASTRADOS ---")
	for _, p := range lista {
		fmt.Printf("%d) %s | R$ %.2f | %d unidades\n", p.ID, p.Nome, p.Preco, p.Quantidade)
	}
}
