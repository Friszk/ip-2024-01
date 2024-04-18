package main

import (
	"fmt"
)

var n int

func main() {

	fmt.Println("Digite o comprimento da lista e em seguida a lista dos números")
	fmt.Scan(&n)
	Lista := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&Lista[i])
	}

	maiorSeqConsAtual := []int{Lista[0]}
	maiorSeqConsEncontrada := make([]int, 0)

	for i := 1; i < n; i++ {
		if Lista[i] == maiorSeqConsAtual[len(maiorSeqConsAtual)-1] {

			maiorSeqConsAtual = append(maiorSeqConsAtual, Lista[i])
		} else {
			maiorSeqConsAtual = []int{Lista[i]}
		}

		if len(maiorSeqConsAtual) > len(maiorSeqConsEncontrada) {
			maiorSeqConsEncontrada = maiorSeqConsAtual
		}
	}

	fmt.Printf("A maior subsequência de elementos iguais possui %v elementos\n.", len(maiorSeqConsEncontrada))
}
