package main

var produtos = []Produto{
	{ID: 1, Nome: "Colar prata", Preco: 49.90, Ativo: true, Estoque: 10},
	{ID: 2, Nome: "Pulseira dourada", Preco: 39.90, Ativo: true, Estoque: 5},
}

var proxID = 3

func adicionaProduto(p Produto) Produto {
	p.ID = proxID
	proxID++
	produtos = append(produtos, p)
	return p
}

func buscaProdutoPorID(id int) *Produto {
	for i := range produtos {
		if produtos[i].ID == id {
			return &produtos[i]
		}
	}
	return nil
}

func removeProdutoPorID(id int) bool {
	for i, p := range produtos {
		if p.ID == id {
			produtos = append(produtos[:1], produtos[i+1:]...)
			return true
		}
	}
	return false
}
