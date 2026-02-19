package main

import (
	"errors"
	"fmt"
)

func main() {
	pilha := Pilha{}
	fmt.Println("pilha criada com tamanho", pilha.Tamanho())
	fmt.Println("Vazia?", pilha.Vazia())

	pilha.Empilhar("Go")
	pilha.Empilhar(2009)
	pilha.Empilhar(3.14)
	pilha.Empilhar("Fim")

	fmt.Println("tamanho após empilhar 4 valores:", pilha.Tamanho())
	fmt.Println("vazia?", pilha.Vazia())

	for !pilha.Vazia() {
		v, _ := pilha.Desempilhar()
		fmt.Println("desempilhando", v)
		fmt.Println("tamanho:", pilha.Tamanho())
		fmt.Println("vazia?", pilha.Vazia())
	}

	_, err := pilha.Desempilhar()
	if err != nil {
		fmt.Println("error:", err)
	}
}

type Pilha struct {
	valores []any
}

func (pilha Pilha) Tamanho() int {
	return len(pilha.valores)
}

func (pilha Pilha) Vazia() bool {
	return pilha.Tamanho() == 0
}

func (pilha *Pilha) Empilhar(valor any) {
	pilha.valores = append(pilha.valores, valor)
}

func (pilha *Pilha) Desempilhar() (any, error) {
	if pilha.Vazia() {
		return nil, errors.New("pilha vazia!")
	}
	valor := pilha.valores[pilha.Tamanho()-1]
	pilha.valores = pilha.valores[:pilha.Tamanho()-1]

	return valor, nil
}
