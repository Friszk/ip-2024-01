package main

import (
	"fmt"
)

func main() {

	var n, i int64
	var eprimo bool

	fmt.Scan(&n)
	if n == 2 {
		fmt.Println("O número 2 é primo")
		return
	} else if n == 1 {
		fmt.Println("O número não é considerado nem primo nem composto")
		return
	} else if n < 1 {
		fmt.Println("Número inválido")
		return
	}

	for i = 2; i < n; i++ {
		if n%i == 0 {
			eprimo = false
			break
		} else {
			eprimo = true
		}
	}

	if eprimo {
		fmt.Printf("O número %v é primo\n", n)
	} else {
		fmt.Printf("O número %v não é primo\n", n)
	}
}
