package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Criando uma lista (slice) vazia de tarefas
	var tarefas []string

	// Criando um leitor para digitar tesxtos completos (aceita espaços)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- Controle de Tarefas ---")

	for {
		fmt.Println("\n----- MENU -----")
		fmt.Println("1 - Adicionar tarefa")
		fmt.Println("2 - Ver tarefas")
		fmt.Println("3 - Sair")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scan(&opcao)

		// Limpa o ENTER que sobra no buffer do teclado
		reader.ReadString('\n')

		if opcao == 1 {
			// Adiciona uma nova tarefa
			fmt.Print("Digite a nova tarefa: ")
			texto, _ := reader.ReadString('\n')
			texto = strings.TrimSpace(texto)

			tarefas = append(tarefas, texto) // Adiciona no slice
			fmt.Println("Tarefas adicionada com sucesso!")

		} else if opcao == 2 {
			// Mostra todas as tarefas
			if len(tarefas) == 0 {
				fmt.Println("Nenhuma tarefa ainda.")
			} else {
				fmt.Println("\nSuas tarefas:")
				for i, tarefa := range tarefas {
					fmt.Printf("%d. %s\n", i+1, tarefa)
				}
			}
		} else if opcao == 3 {
			fmt.Println("Saindo... Até logo")
			break

		} else {
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
