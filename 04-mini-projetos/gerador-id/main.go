package main

import (
	"bufio"
	"fmt"
	"os"

	"gerador-id/models" // mesmo nome do module no go.mod
	"gerador-id/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var ids []models.ID // Slice para guardar todos os IDs gerados

	for {
		fmt.Println("\n---- GERADOR DE ID ----")
		fmt.Println("1 - Gerar novo ID")
		fmt.Println("2 - Listar IDs gerados")
		fmt.Println("3 - Sair")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scan(&opcao)
		reader.ReadString('\n') // Limpa o ENTER depois do Scan

		switch opcao {
		case 1:
			// gerar novo ID usando o pacote utils
			novo := utils.GeraRID()
			ids = append(ids, novo)

			fmt.Println("\nID gerado com sucesso:")
			fmt.Println("novo.Codigo")

		case 2:
			if len(ids) == 0 {
				fmt.Println("\nNenhum ID gerado ainda.")
			} else {
				fmt.Println("\n---- IDs GERADOS ----")
				for i, id := range ids {
					fmt.Printf("%d) %s\n", i+1, id.Codigo)
				}
			}

		case 3:
			fmt.Println("\nSaindo... Até logo!")
			return

		default:
			fmt.Println("\nOpção inválida. Tente novamente.")

		}
	}
}
