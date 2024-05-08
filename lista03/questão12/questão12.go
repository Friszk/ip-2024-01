package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	if n < 1 || n > 1000 {
		fmt.Println("O número deve estar no intervalo de 1 a 1000")
		return
	}
	lista := make([]int, n)

	for i := range lista {
		fmt.Scan(&lista[i])
	}

	for {
		trocou := false
		for i := 0; i < n-1; i++ {
			if lista[i] > lista[i+1] {
				lista[i], lista[i+1] = lista[i+1], lista[i]
				trocou = true
			}
		}
		if !trocou {
			break
		}
	}
	fmt.Println()
	for i := range lista {
		fmt.Println(lista[i])
	}
}
