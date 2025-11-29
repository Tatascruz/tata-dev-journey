# Capítulo 1 - Pesquisa Binária

Este capítulo apresenta os conceitos fundamentais de buscas em lista e mostra como a **pesquisa binária** é uma solução muito mais eficientedo que a busca
linear, principalmente quando trabalhamos com listas grandes.

---

## O que aprendi neste capítulo

- Diferença entre **busca linear** e **pesquisa binária**
- Como a binária elimina sempre **metade** da lista a cada tentativa
- Que pesquisa binária **só funciona** com lista ordenadas
- Implementar os algoritmos em Go
- Comparar desempenho entre os dois tipos de busca 
- Aplicar pesquisa binária em lista de **string** (palavras)

## Exercícios realizados

### **Exercício 1 - Busca Linear**
Implementação da busca simples, analisando um item por vez.
Resultado: funciona, mas é lenta para listas muito grandes.

Arquivo: exercicio_busca_linear.go

---

### **Exercício 2 - Pesquisa Binária (números)**
Implementação clássica de pesquisa binária para números inteiros ordenados.

Arquivo: exercicio2_pesquisa_binaria.go

---
### **Exercício 3 - Comparando eficiência**
Comparação entre busca linear e binária usando uma lista de **5000 números**.

Resultado:
- "Tata" encontrada em **4 tentativas**
- Prova de que funciona também para string.

Arquivo: exercicio4_palavras.go

---

## Conceitos importantes

### Pesquisa Binária
- Divide a lista ao meio repetidamente
- Muito mais rápida
- Precisa de lista **ordenada**
- Complexidade: **0(log n)**

---

## Observações finais 

Este capítulo foi essencial para entender como algoritmos de busca funcionam.
Aprendi a importancia de ordenar listas e como podemos otimizar uma tarefa simples usando
lógica e matemática. 

