package main

import (
	"bufio" // Para ler linhas inteiras (aceita espaços)
	"fmt"
	"os"
	"strings" // Para limpar espaços e quebras de linha
)

// Struct que representa um usuário do sistema
type Usuario struct {
	Nome  string
	Email string
	Senha string
}

func main() {
	var usuarios []Usuario              // Slice para guardar todos os usúarios cadastrados
	reader := bufio.NewReader(os.Stdin) // Reader para ler textos com espaços (ex: "Maria da Silva")
	var opcao int

	// Loop principal do sistema
	for {
		fmt.Println("\n==============================")
		fmt.Println("       SISTEMA DE LOGIN         ")
		fmt.Println("\n==============================")
		fmt.Println("1 Cadastrar usuário")
		fmt.Println("2 Fazer login")
		fmt.Println("3 Listar usuários")
		fmt.Println("4 Sair")
		fmt.Println("=================================")
		fmt.Print("Escolha uma opção: ")

		fmt.Scan(&opcao)
		reader.ReadString('\n') // Aqui limpamos o ENTER que fica "sobrando" depois do Scan

		switch opcao {
		case 1:
			fmt.Println("\n Opção selecionada: Cadastrar usuário")
			usuarios = cadastrarUsuario(usuarios, reader)

		case 2:
			fmt.Println("\n Opção selecionada: Fazer login")
			fazerLogin(usuarios, reader)

		case 3:
			fmt.Println("\n Usuários cadastrados até agora:")
			listarUsuarios(usuarios)

		case 4:
			fmt.Println("\n Saindo do sistema... Até logo!!!")
			return

		default:
			fmt.Println("\n Opção inválida. Tente novamente.")
		}
	}

}

// Função para cadastrar um novo usúario
// Ela recebe lista de usúarios atual eo reader
// e devolve a lista ATUALIZADA com novo usúario
func cadastrarUsuario(usuarios []Usuario, reader *bufio.Reader) []Usuario {
	var u Usuario

	fmt.Println("\n Cadastro de novo usúario")

	// Lendo o nome
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

	// Adiona o novo usuário na lista
	usuarios = append(usuarios, u)

	fmt.Println("Usúario cadastrado com sucesso!")
	fmt.Print("Total de usúarios cadastrados: %d\n", len(usuarios))

	// Retorna a lista atualizada
	return usuarios
}

// --------------LOGIN--------------

func fazerLogin(usuarios []Usuario, reader *bufio.Reader) {
	// Se não tiver ninguém cadastrado ainda, nem adianta tentar logar
	if len(usuarios) == 0 {
		fmt.Println("Nenhum usúario cadastrado. Cadastre alguém antes de fazer login.")
		return
	}

	fmt.Println("\n Tela de Login")

	fmt.Print("Email: ")
	emailDigitado, _ := reader.ReadString('\n')
	emailDigitado = strings.TrimSpace(emailDigitado)

	fmt.Print("Senha: ")
	senhaDigitada, _ := reader.ReadString('\n')
	senhaDigitada = strings.TrimSpace(senhaDigitada)

	// Vamos procurar um usúario com esse email e essa senha
	for _, u := range usuarios {
		if u.Email == emailDigitado && u.Senha == senhaDigitada {
			fmt.Printf("\n Login realizado com sucesso! Bem-vindo(a), %s.\n", u.Nome)
			return // sai da função depois de logar
		}
	}

	// Se passou pelo for inteiro e não achou, é porque deu errado
	fmt.Println("\n Email ou senha incorretos.")

}

// ---------------LISTAGEM-------------

func listarUsuarios(usuarios []Usuario) {
	if len(usuarios) == 0 {
		fmt.Println("Nenhum usuário cadastrado ainda.")
		return
	}

	fmt.Println("Lista de usúarios:")
	for i, u := range usuarios {
		fmt.Printf("%d) %s - %s\n", i+1, u.Nome, u.Email)
	}
}
