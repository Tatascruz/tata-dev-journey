package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler para a rota /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// só permitimos o método GET aqui
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		fmt.Fprintln(w, "Método não permitido")
		return
	}

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

	fmt.Println("Servidor rodando em http://localhost:8080")
	fmt.Println("Acesse http://localhost:8080/hello ou /status no navegador")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
