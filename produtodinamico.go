package main

import "fmt"

func main() {

	// Aqui estamos criando 3 variáveis para guardar os dados dos produtos
	// string = texto (palavras)
	// float64 = número com casas decimais (preço)
	var nomeProduto string
	var material string
	var valor float64

	// fmt.Print mostra algo na tela, mas não pula linha
	// Aqui estamos pedindo para o usuário difitar o nome do produto
	fmt.Print("Digite o nome do produto: ")
	fmt.Scan(&nomeProduto)

	// fmt.Scan lê o que o usuário digitou e guarda na variável
	// & significa "endereço a variável", algo do Go que você só precisa lembrar de usar
	fmt.Print("Digite o material: ")
	fmt.Scan(&material)

	fmt.Print("Digite o valor: ")
	fmt.Scan(&valor)

	// Agora vamos exibir tudo formatado bonitinho
	fmt.Printf("\n--- Produto Cadastrado ---\n")

	// fmt.Printf permite colocar valores dentro da frase
	// %s -> string (texto)
	// %.2f -> número com 2 casas decimais (preço bonito)
	fmt.Printf("Nome: %s\nMaterial: %s\nValor: R$ %.2f\n", nomeProduto, material, valor)
}
