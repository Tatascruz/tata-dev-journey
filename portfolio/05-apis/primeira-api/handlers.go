package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"tata-dev-journey/05-apis/primeira-api/database"
	"tata-dev-journey/05-apis/primeira-api/models"
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

	case http.MethodPut:
		// 1) pegar ID da URL
		idStr := strings.TrimPrefix(r.URL.Path, "/produtos/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			respondError(w, http.StatusBadRequest, "ID inválido")
			return
		}

		// 2) Ler JSON do body
		var dados models.Produto
		if err := json.NewDecoder(r.Body).Decode(&dados); err != nil {
			respondError(w, http.StatusBadRequest, "JSON inválido")
			return
		}

		// 3) Buscar no banco
		var produto models.Produto
		result := database.DB.First(&produto, id)
		if result.Error != nil {
			respondError(w, http.StatusNotFound, "Produto não encontrado")
			return
		}

		// 4) Atualizar campos
		produto.Nome = dados.Nome
		produto.Preco = dados.Preco
		produto.Estoque = dados.Estoque
		produto.Ativo = dados.Ativo

		// 5) Salvar
		if err := database.DB.Save(&produto).Error; err != nil {
			respondError(w, http.StatusInternalServerError, "Erro ao atualizar produto")
			return
		}

		// 6) Responder
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(produto)

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
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	switch r.Method {

	case http.MethodGet:
		p := buscaProdutoPorID(id)
		if p == nil {
			respondError(w, http.StatusNotFound, "Produto não encontrado")
			return

		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)

	case http.MethodPut:
		p := buscaProdutoPorID(id)
		if p == nil {
			respondError(w, http.StatusNotFound, "Produto não encontrado")
			return
		}

		var dados Produto
		if err := json.NewDecoder(r.Body).Decode(&dados); err != nil {
			respondError(w, http.StatusBadRequest, "JSON inválido")
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
			respondError(w, http.StatusNotFound, "Produto não encontrado")
			return
		}

		w.WriteHeader(http.StatusNoContent) //204 = sem corpo

	default:
		respondError(w, http.StatusMethodNotAllowed, "Método não permitido")

	}

}

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})

}
