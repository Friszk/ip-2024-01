package main

import (
	"fmt"
)

var n, r, a, termoPA, soma int

func main() {

	fmt.Println("Digite o valor inicial da P.A, a razão e o número de elementos")
	fmt.Scan(&a, &r, &n)

	termoPA = a

	for i := 0; i < n; i++ {
		termoPA = a + r*i
		soma = termoPA + soma
	}
	fmt.Println(soma)
}
