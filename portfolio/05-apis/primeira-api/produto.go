package main

type Produto struct {
	ID      int     `json:"id" gorm:"primaryKey"`
	Nome    string  `json:"nome"`
	Preco   float64 `json:"preco"`
	Ativo   bool    `json:"ativo"`
	Estoque int     `json:"estoque"`
}
