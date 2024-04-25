package main

import (
	"fmt"
)

func main() {

	fmt.Println("Digite um número e precisão da raiz quadrada a ser calculada utilizando o método\nBabilônico ")
	var n, e, r float64
	fmt.Scan(&n, &e)

	r = 1
	a := n - (r * r)
	for a > e {
		r = ((n / r) + r) / 2
		a = n - (r * r)
		if a < 0 {
			a = (n - (r * r)) * -1
		}
		fmt.Printf("%.9f %.9f\n", r, a)
	}

}
