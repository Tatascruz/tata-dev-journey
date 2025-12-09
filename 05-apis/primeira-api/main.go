package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler para a rota  /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Só permitimos o método GET aqui
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Método não permetido")
		return
	}

	fmt.Fprintln(w, "Olá, Tata! Esta é a sua primeira API em Go")
}

func main() {
	// registra s rota /hello e diz quem cuida dela é helloHandler
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	fmt.Println("Acesse http://localhost:8080/hello no navegador")

	// sobe servidor na porta 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
