package main

import (
	"fmt"
	"time"
)

func gerarFibonacci(n int) func() {
	return func() {
		a, b := 0, 1

		fib := func() int {
			a, b = b, a+b
			return a
		}

		for range n {
			fmt.Printf("%d ", fib())
		}
	}
}

func cronometrar(funcao func()) {
	inicio := time.Now()
	funcao()
	fmt.Printf("\nTempo de execução: %s\n\n", time.Since(inicio))
}

func main() {
	cronometrar(gerarFibonacci(8))
	cronometrar(gerarFibonacci(48))
	cronometrar(gerarFibonacci(88))
}
