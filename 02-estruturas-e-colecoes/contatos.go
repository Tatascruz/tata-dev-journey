package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Criamos um tipo chamado (struct)
// Uma struct serve para agrupar várias informações relacionadas
type Contato struct {
	Nome     string
	Telefone string
	Email    string
}

func main() {
	var contatos []Contato              // slice de contatos
	reader := bufio.NewReader(os.Stdin) // leitor para ler textos com espaços

	fmt.Println("*** Sistema de Contatos ***")

	for {
		fmt.Println("\n----- MENU -----")
		fmt.Println("1 - Adicionar contato")
		fmt.Println("2 - Ver contatos")
		fmt.Println("3 - Sair")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scan(&opcao)
		reader.ReadString('\n') // Limpa o ENTER do buffer

		if opcao == 1 {
			var c Contato // Cria um novo contato

			fmt.Print("Nome: ")
			nome, _ := reader.ReadString('\n')
			c.Nome = strings.TrimSpace(nome)

			fmt.Print("Telefone: ")
			tel, _ := reader.ReadString('\n')
			c.Telefone = strings.TrimSpace(tel)

			fmt.Print("E-mail: ")
			email, _ := reader.ReadString('\n')
			c.Email = strings.TrimSpace(email)

			contatos = append(contatos, c)
			fmt.Println("Contato adicionado com sucesso!")

		} else if opcao == 2 {
			if len(contatos) == 0 {
				fmt.Println("Nenhum contato cadastrado.")
			} else {
				fmt.Println("\nLista de Contatos:")
				for i, contato := range contatos {
					fmt.Printf("%d. %s | %s | %s\n", i+1, contato.Nome, contato.Telefone, contato.Email)
				}
			}

		} else if opcao == 3 {
			fmt.Println("Saindo... Até mais")
			break
		} else {
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
