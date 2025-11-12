package main

import "fmt"

func main() {

	// Variáveis da primeira caixa
	// Cada variável guarda uma informação (nome, material, valor)
	var nome1 string
	var material1 string
	var valor1 float64

	//Variáveis da segunda caixa
	var nome2 string
	var material2 string
	var valor2 float64

	// --- Cadastro da primeira caixa ---
	fmt.Println("Cadastro da primeira caixa:")

	fmt.Print("Nome: ")
	fmt.Scan(&nome1)

	fmt.Print("Material: ")
	fmt.Scan(&material1)

	fmt.Print("Valor: ")
	fmt.Scan(&valor1)

	fmt.Println("\n-----------------------\n")

	// --- Cadastro da segunda caixa ---
	fmt.Println("Cadastro da segunda caixa:")

	fmt.Print("Nome: ")
	fmt.Scan(&nome2)

	fmt.Print("Material: ")
	fmt.Scan(&material2)

	fmt.Print("Valor: ")
	fmt.Scan(&valor2)

	// Exibindo as duas caixas
	fmt.Println("\nCaixas cadastradas com sucesso!")

	// %s = texto | %.2f = número com duas casas decimais
	fmt.Printf("\n1) %s | %s | R$ %.2f\n", nome1, material1, valor1)
	fmt.Printf("2) %s | %s | R$ %.2f\n", nome2, material2, valor2)
}
