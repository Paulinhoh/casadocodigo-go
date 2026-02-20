package main

import (
	"fmt"
	"sort"
)

type Pessoa struct {
	nome  string
	idade int
}

func main() {
	pessoas := []Pessoa{
		{nome: "Tiago", idade: 23},
		{nome: "Jaime", idade: 75},
		{nome: "Ronaldo", idade: 42},
		{nome: "Anderson", idade: 14},
	}

	fmt.Println("Pessoas:", pessoas)

	sort.Slice(pessoas, func(i, j int) bool {
		return pessoas[i].idade < pessoas[j].idade
	})

	fmt.Println("Pessoas por ordem de idade:", pessoas)
}
