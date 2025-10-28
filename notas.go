package main

import "fmt"

func main() {
	// Declarando as variáveis
	var nome string
	var nota1, nota2, media float64

	fmt.Println("--- Sistemas de Notas ---")

	// Pedindo o nome do aluno
	fmt.Print("Digite o nome do aluno: ")
	fmt.Scan(&nome)

	// Pedindo as notas
	fmt.Print("Digite a primeira nota: ")
	fmt.Scan(&nota1)

	fmt.Print("Digite a segunda nota: ")
	fmt.Scan(&nota2)

	// Calculando a média
	media = (nota1 + nota2) / 2

	fmt.Printf("\nAluno: %s\n", nome)
	fmt.Printf("Média: %.2f\n", media)

	// Estrutura condicional para avaliar o resultado
	if media >= 7 {
		fmt.Println("Aprovado! Parabéns!")
	} else if media >= 5 {
		fmt.Println("Recuperação. Ainda da tempo de melhorar!!!")
	} else {
		fmt.Println("Reprovado. Tente novamente no próximo semestre.")
	}
}
