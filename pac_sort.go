package main

import (
	"fmt"
	"sort"
)

type Dados struct {
	Nome  string
	Idade int
}

type ParaNome []Dados

func (ps ParaNome) Len() int {
	return len(ps)
}

func (ps ParaNome) Less(i, j int) bool { // ver se o item na posição i é menor que o item na posição j
	return ps[i].Nome < ps[j].Nome
}

func (ps ParaNome) Swap(i, j int) { // swap troca os itens
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	criancas := []Dados{
		{"Marina", 5},
		{"João", 10},
	}
	sort.Sort(ParaNome(criancas))
	fmt.Println(criancas)
}
