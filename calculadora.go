package main

import "fmt"

func main() {
	// Declarações das variáveis
	var num1, num2 float64
	var operacao string

	fmt.Println("------ Calculadora Simples ------")

	// Entrada dos números
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	// Escolha da operação
	fmt.Println("\nEscolha a operação:")
	fmt.Println("+ para soma")
	fmt.Println("- para subtração")
	fmt.Println("* para multiplecação")
	fmt.Println("/ para divisão")
	fmt.Println("Operação: ")
	fmt.Scan(&operacao)

	// Aqui usamos switch, que é como um "menu inteligente"
	switch operacao {
	case "+":
		fmt.Printf("Resultado: %.2f\n", somar(num1, num2))
	case "-":
		fmt.Printf("Resultado: %.2f\n", subtrair(num1, num2))
	case "*":
		fmt.Printf("Resultado: %.2f\n", multiplicar(num1, num2))
	case "/":
		// Tratando divisão por zero
		if num2 == 0 {
			fmt.Println("Erro: divisão por zero não é permitida.")
		} else {
			fmt.Printf("Resultado: %.2f\n", dividir(num1, num2))
		}
	default:
		fmt.Println("Operação inválida")
	}
}

// Funções da calculadora
// Criamos funções separadas pra cada operação
// Isso ajuda a deixar o código mais organizado e fácil de entender

func somar(a, b float64) float64 {
	return a + b
}

func subtrair(a, b float64) float64 {
	return a - b
}

func multiplicar(a, b float64) float64 {
	return a * b
}

func dividir(a, b float64) float64 {
	return a / b
}
