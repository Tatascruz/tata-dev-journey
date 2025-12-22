package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

/*
ROTAS:
GET /produtos
POST /produtos
GET /produtos/{id}
PUT /produtos/{id}
DELETE /produtos/{id}
*/

func produtosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(produtos)

	case http.MethodPost:
		var novo Produto

		if err := json.NewDecoder(r.Body).Decode(&novo); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "JSON inválido")
			return
		}

		// Gera ID e salva
		novo.ID = gerarNovoID()
		produtos = append(produtos, novo)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(novo)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Método não permitido")
	}

}

// Produtos/{id} -> GET busca / PUT atualiza / DELETE remove
func produtoPorIDHandler(w http.ResponseWriter, r *http.Request) {
	//Ex: /produtos/1 -> queremos pegar "1"
	idStr := strings.TrimPrefix(r.URL.Path, "/produtos/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "ID inválido")
		return
	}

	switch r.Method {

	case http.MethodGet:
		p := buscaProdutoPorID(id)
		if p == nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "Produto não encontrado")
			return

		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)

	case http.MethodPut:
		p := buscaProdutoPorID(id)
		if p == nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "Produto não encontrado")
			return
		}

		var dados Produto
		if err := json.NewDecoder(r.Body).Decode(&dados); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "JSON inválido")
			return
		}

		//Atualiza mantendo o ID
		p.Nome = dados.Nome
		p.Preco = dados.Preco
		p.Ativo = dados.Ativo
		p.Estoque = dados.Estoque

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)

	case http.MethodDelete:
		ok := removeProdutoPorID(id)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "Produto não encontrado")
			return
		}

		w.WriteHeader(http.StatusNoContent) //204 = sem corpo

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Método não permitido")

	}

}
