package main

import "fmt"

type agregadora func(n, m int) int

func agregar(valores []int, valorInicial int, fn agregadora) int {
	agregado := valorInicial

	for _, v := range valores {
		agregado = fn(v, agregado)
	}

	return agregado
}

func calcularSoma(valores []int) int {
	soma := func(n, m int) int {
		return n + m
	}
	return agregar(valores, 0, soma)
}

func calcularProduto(valores []int) int {
	multiplicacao := func(n, m int) int {
		return n * m
	}
	return agregar(valores, 1, multiplicacao)
}

func main() {
	valores := []int{3, -2, 5, 7, 8, 22, 32, -1}

	fmt.Println(calcularSoma(valores))
	fmt.Println(calcularProduto(valores))
}
