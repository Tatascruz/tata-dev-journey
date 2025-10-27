package main

import "fmt"

func main() {
	// 1- Usando Print (sem quebra de linha)
	fmt.Print("Bom dia")
	fmt.Print("Tata")

	// 2- Usando Println (com quebra de linha )
	fmt.Println("\nHoje vamos aprender Go!!")

	// 3- Usando Printf (formando texto)
	nome := "Tata"
	idade := 40
	fmt.Printf("Meu nome é %s e eu tenho %d anos.\n", nome, idade)
}
