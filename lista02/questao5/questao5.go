package main

import (
	"fmt"
)

var n int

func main() {

	fmt.Println("Digite o comprimento da lista e em seguida a lista dos números")
	fmt.Scan(&n)
	if n < 1 {
		fmt.Println("O comprimento deve ser maior que 0")
		return
	}
	Lista := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&Lista[i])
	}

	maiorCompEncontrado := make([]int, 0)
	maiorCompAtual := []int{Lista[0]}

	for i := 1; i < n; i++ {

		if Lista[i] > maiorCompAtual[len(maiorCompAtual)-1] {

			maiorCompAtual = append(maiorCompAtual, Lista[i])
		} else {
			maiorCompAtual = []int{Lista[i]}
		}

		if len(maiorCompAtual) > len(maiorCompEncontrado) {
			maiorCompEncontrado = maiorCompAtual
		}

	}

	fmt.Printf("O comprimento do segmento crescente máximo é: %v\n", len(maiorCompEncontrado)-1)
}
