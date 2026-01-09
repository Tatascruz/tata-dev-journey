API de produtos em Go (REST)

Este projeto é uma **API REST simples desenvolvida em Go**, criada como parte do meu plano de estudos para 
praticar conceitos fundamentais de backend.
A aplicação  simula o gerenciamento de produtos, permitindo criar, listar, buscar, atualizar e remover registros via HTTP.

## Objetivos de estudo

Neste projeto foram praticados 

- Criação de API REST com Go
- Manipulção de rotas HTTP
- Métodos REST (GET, POST, PUT, DELETE)
- Manipulação de JSON 
- Organizar códigos em camadas
- handlers
- models
- database / storage
- Testes de API com Postman

---

## Tecnologias utilizadas

- Go (Golang)
- net/http
- enconding/json
- SQLite (para persistência simples)
- Postman (para testes das requisições)

---

## Rotas da API

Método    |         Rota        |     Descrição
--------------------------------------------------------
GET       | /produtos           | Lista todos os produtos
POST      | /produtos           | Cria um novo produto
GET       | /produtos/{id}      | Busca produto por ID
PUT       | /produtos/{id}      | Atualiza produto por ID
DELETE    | /produtos/{id}      | Remove produto por ID

---

## Exemplo de produto (JSON)

```JSON
{
  "nome": "Colar prata",
  "preco": 49.9,
  "estoque": 10,
  "ativo": true
}



**Como executar o projeto**

### 1. Clone o repositório
```bash
git clone https://github.com/Tatascruz/tata-dev-journey.git


