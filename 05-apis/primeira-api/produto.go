package main

type Produto struct {
	ID      int     `json:"id"`
	Nome    string  `json:"nome"`
	Preco   float64 `json:"preco"`
	Ativo   bool    `json:"ativo"`
	Estoque int     `json:"estoque"`
}
