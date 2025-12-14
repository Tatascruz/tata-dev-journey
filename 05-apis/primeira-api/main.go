package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler para a rota /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Olá, Tata! Esta é a sua primeira API em Go ")
}

// Handler para a rota /status
func statusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Método não permitido")
		return
	}
	fmt.Fprintln(w, "API rodando :) Tudo certo por aqui!!!")
}

func main() {
	// Registra a rota /hello
	http.HandleFunc("/hello", helloHandler)

	// Registra a rota /status
	http.HandleFunc("/status", statusHandler)

	//Registra a rota /produtos
	http.HandleFunc("/produtos", produtosHandler)
	http.HandleFunc("/produtos/", produtoPorIDHandler)

	fmt.Println("Servidor rodando em http://localhost:8081")
	fmt.Println("Rotas disponíveis:")
	fmt.Println("GET    /produtos")
	fmt.Println("POST   /produtos")
	fmt.Println("GET    /produtos/{id}")
	fmt.Println("DELETE /produtos/{id}")

	log.Fatal(http.ListenAndServe(":8081", nil))
}
