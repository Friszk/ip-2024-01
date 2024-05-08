package main

import "fmt"

func main() {
	var n, soma int
	fmt.Scan(&n)
	if n > 5000 {
		fmt.Println("O número deve ser menor que 5000")
		return
	}

	numeros := make([]int, n)

	for i := range numeros {
		fmt.Scan(&numeros[i])
	}

	for _, v := range numeros {
		soma += v
	}

	fmt.Println(soma)
}
