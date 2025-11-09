package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Slice para guardar as movimentações (positivas e negativas)
	var movimentos []float64
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Controle Financeiro")

	for {
		fmt.Println("\n----MENU----")
		fmt.Println("1 - Registrar entrada")
		fmt.Println("2 - Registrar saída")
		fmt.Println("3 - Ver saldo atual")
		fmt.Println("4 - Sair")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scan(&opcao)
		reader.ReadString('\n') // Limpa o enter

		if opcao == 1 {
			// Registrar entrada
			valor := lerValor(reader, "Digite o valor da entrada: ")
			movimentos = append(movimentos, valor)
			fmt.Println("Entrada registrada com sucesso!")

		} else if opcao == 2 {
			// Registrar Saída
			valor := lerValor(reader, "Digite o valor da saída: ")
			movimentos = append(movimentos, -valor) // Saída é valor negativo
			fmt.Println("Saída registrada com sucesso!")

		} else if opcao == 3 {
			// Cacular saldo total
			saldo := 0.0
			for _, m := range movimentos {
				saldo += m
			}

			fmt.Println("\nResumo Financeiro")
			fmt.Printf("Total de movimentações: %d\n", len(movimentos))
			fmt.Printf("Saldo atual: R$ %.2f\n", saldo)

			if saldo > 0 {
				fmt.Println("Saldo positivo! Continue assim!")
			} else if saldo == 0 {
				fmt.Println("Nenhum saldo disponível.")
			} else {
				fmt.Println("Saldo negativo! Cuidado com as finanças!")
			}

		} else if opcao == 4 {
			fmt.Println("Saindo... até a próxima")
			break
		} else {
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}

// Função auxiliar para ler valores com vírgula ou ponto
func lerValor(reader *bufio.Reader, prompt string) float64 {
	for {
		fmt.Print(prompt)
		texto, _ := reader.ReadString('\n')
		texto = strings.TrimSpace(texto)
		texto = strings.ReplaceAll(texto, ",", ".")
		valor, err := strconv.ParseFloat(texto, 64)
		if err == nil {
			return valor
		}
		fmt.Println("Valor inválido. Digite novamente (ex: 125.90)")
	}
}
