package main

import (
	"fmt"
	"sort"
)

type Pessoa struct {
	nome  string
	idade int
}

type PorIdade []Pessoa

func (p PorIdade) Len() int {
	return len(p)
}

func (p PorIdade) Less(i, j int) bool {
	return p[i].idade < p[j].idade
}

func (p PorIdade) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	pessoas := []Pessoa{
		{nome: "Tiago", idade: 23},
		{nome: "Jaime", idade: 75},
		{nome: "Ronaldo", idade: 42},
		{nome: "Anderson", idade: 14},
	}

	fmt.Println("Pessoas:", pessoas)

	sort.Sort(PorIdade(pessoas))

	fmt.Println("Pessoas por ordem de idade:", pessoas)
}
