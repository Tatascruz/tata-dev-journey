package main

import (
	"bufio" // Para ler linhas inteiras (aceita espaços)
	"fmt"
	"os"
	"strings" // Para limpar espaços e quebras de linha

	"tata-dev-journey/04-mini-projetos/login/models"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Reader para ler textos com espaços (ex: "Maria da Silva")

	var usuarios []models.User // Slice para guardar todos os usúarios cadastrados

	// Loop principal do sistema
	for {
		fmt.Println("------------------")
		fmt.Println("SISTEMA DE LOGIN")
		fmt.Println("-------------------")
		fmt.Println("1 Cadastrar usuário")
		fmt.Println("2 Listar usuários")
		fmt.Println("3 Sair")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scan(&opcao)

		reader.ReadString('\n') // Aqui limpamos o ENTER que fica "sobrando" depois do Scan

		switch opcao {
		case 1:
			novo := cadastrarUsuario(reader, len(usuarios)+1)
			usuarios = append(usuarios, novo)

		case 2:
			listarUsuarios(usuarios)

		case 3:
			fmt.Println("Saindo... Até logo!")
			return

		default:
			fmt.Println("\n Opção inválida. Tente novamente.")
		}
	}

}

// Função para cadastrar um novo usúario
// Ela recebe lista de usúarios atual eo reader
// e devolve a lista ATUALIZADA com novo usúario
func cadastrarUsuario(reader *bufio.Reader, id int) models.User {
	var u models.User

	u.ID = id

	fmt.Print("Nome: ")
	nome, _ := reader.ReadString('\n')
	u.Nome = strings.TrimSpace(nome) // Remove \n e espaços das pontas

	// Lendo o email
	fmt.Print("email: ")
	email, _ := reader.ReadString('\n')
	u.Email = strings.TrimSpace(email)

	// Lendo a senha
	fmt.Print("Senha: ")
	senha, _ := reader.ReadString('\n')
	u.Senha = strings.TrimSpace(senha)

	fmt.Println("\nUsúario cadastrado com sucesso!")

	return u
}

// --------------LOGIN--------------

func listarUsuarios(usuarios []models.User) {
	if len(usuarios) == 0 {
		fmt.Println("Nnenhum usúario cadastrado ainda.")
		return
	}

	fmt.Println("\n ---- USUÁRIOS CADASTRADOS ----")
	for _, u := range usuarios {
		fmt.Printf("%d)  %s - %s\n", u.ID, u.Nome, u.Email)

	}
}
