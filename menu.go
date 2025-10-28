package main

import (
	"bufio" // Lê linhas inteiras (inclusive espaços)
	"fmt"
	"os"
	"strconv" // Converte texto -> número
	"strings" // Utilidades para texto
)

func main() {
	// reader lê a entrada do teclado lçinha por linha
	reader := bufio.NewReader(os.Stdin)

	// Mostrando opções na tela
	fmt.Println("------- MENU -------")
	fmt.Println("1 - Cadastrar uma caixa")
	fmt.Println("2 - Sair")

	// Lê a opção de TEXTO e converte pra número
	opcaoStr := readLine(reader, "Escolha uma opção: ")
	opcao, _ := strconv.Atoi(strings.TrimSpace(opcaoStr)) // se não for número, cai no else

	// Agora vamos verificar o que a pessoa escolheu
	// if = se | else = caso contrário
	if opcao == 1 {
		// Cadastro da caixa
		nome := readLine(reader, "Nome da caixa: ")
		material := readLine(reader, "Material: ")
		valor := readFloat(reader, "Valor: ")

		fmt.Println("\nCaixa cadastrada")
		fmt.Printf("%s | %s | R$ %.2f\n", nome, material, valor)

	} else if opcao == 2 {

		// Se escolher a opção 2, apenas finaliza
		fmt.Println("Saindo... Até amanhã")
	} else {
		// Se digitar qualquer outra coisa, não aceitamos
		fmt.Println("Opção inválida")
	}

}

// Funções auxiliares (com explicações)

// reaLine mostra um prompt e LÊ a LINHA INTEIRA  (aceita espaços)
func readLine(reader *bufio.Reader, promt string) string {
	fmt.Print(promt)
	texto, _ := reader.ReadString('\n') // Lê até apertar Enter
	return strings.TrimSpace(texto)     // Tira \n espaços extras das pontas
}

// readFloat lê um texto , aceita vírgula OU ponto e converte para float64
func readFloat(reader *bufio.Reader, prompt string) float64 {
	for {
		s := readLine(reader, prompt)
		s = strings.ReplaceAll(s, ",", ".") // Permite digitar 55,90
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			return f
		}
		fmt.Println("Valor inválido. Digite apenas número. Ex: 65.90")
	}
}
