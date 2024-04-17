package main

import (
	"fmt"
)

var n, i, K, s, r float64

func main() {

	fmt.Println("Digite o número a ser apresentada sua tabuada, onde será iniciado, sua quantidade de valores e seu incremento")
	fmt.Scan(&n, &i, &K, &s)

	fmt.Println("Tabuada de soma")
	r = n + i
	fmt.Printf("%.2f + %.2f = %.2f\n", n, i, r)

	for o := 0.0; o < i; o++ {
		r = n + (i + s)
		fmt.Printf("%.2f + %.2f = %.2f\n", n, i+s, r)
	}

	fmt.Println("Tabuada de subtração")
	r = n - i
	fmt.Printf("%.2f - %.2f = %.2f\n", n, i, r)

	for o := 0.0; o < i; o++ {
		r = n - (i + s)
		fmt.Printf("%.2f - %.2f = %.2f\n", n, i+s, r)
	}

	fmt.Println("Tabuada de multiplicação")
	r = n * i
	fmt.Printf("%.2f x %.2f = %.2f\n", n, i+s, r)

	for o := 0.0; o < i; o++ {
		r = n * (i + s)
		fmt.Printf("%.2f x %.2f = %.2f\n", n, i, r)
	}

	fmt.Println("Tabuada de divisão")
	r = n / i
	fmt.Printf("%.2f / %.2f = %.2f\n", n, i+s, r)

	for o := 0.0; o < i; o++ {
		r = n / (i + s)
		fmt.Printf("%.2f / %.2f = %.2f\n", n, i, r)
	}
}
