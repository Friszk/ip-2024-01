package main

import (
	"fmt"
)

var n, i, K, s, r float64

func main() {

	fmt.Println("Digite o número a ser apresentada sua tabuada, onde será iniciado, sua quantidade de valores e seu incremento")
	fmt.Scan(&n, &i, &K, &s)

	fmt.Println("Tabuada de soma")

	for o, in := 0.0, 0.0; o < K; o++ {
		r = n + (i + in)
		in += s
		fmt.Printf("%.2f + %.2f = %.2f\n", n, in, r)
	}

	fmt.Println("Tabuada de subtração")

	for o, in := 0.0, 0.0; o < K; o++ {
		r = n - (i + in)
		in += s
		fmt.Printf("%.2f - %.2f = %.2f\n", n, in, r)
	}

	fmt.Println("Tabuada de multiplicação")

	for o, in := 0.0, 0.0; o < K; o++ {
		r = n * (i + in)
		in += s
		fmt.Printf("%.2f x %.2f = %.2f\n", n, in, r)
	}

	fmt.Println("Tabuada de divisão")

	for o, in := 0.0, 0.0; o < K; o++ {
		r = n / (i + in)
		in += s
		fmt.Printf("%.2f / %.2f = %.2f\n", n, in, r)
	}
}
