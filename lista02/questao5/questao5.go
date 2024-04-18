package main

import (
	"fmt"
)

var tamanhoLista int

func main() {

	fmt.Println("Digite o comprimento da lista e em seguida a lista dos números")
	fmt.Scan(&tamanhoLista)

	Lista := make([]int, tamanhoLista)

	for i := 0; i < tamanhoLista; i++ {
		fmt.Scan(&Lista[i])
	}

	maiorCompEncontrado := make([]int, 0, 1)
	maiorCompAtual := []int{Lista[0]}

	for i := 1; i < tamanhoLista; i++ {
		if Lista[i] > maiorCompAtual[len(maiorCompAtual)-1] {

			maiorCompAtual = append(maiorCompAtual, Lista[i])

		} else {
			maiorCompAtual = []int{Lista[i]}
		}

		if len(maiorCompAtual) > len(maiorCompEncontrado) {
			maiorCompEncontrado = maiorCompAtual
		}

	}

	if maiorCompEncontrado[0] == 0 && len(maiorCompEncontrado) == 1 {
		fmt.Println("O comprimento do segmento crescente máximo é: 0")
	} else {
		fmt.Printf("O comprimento do segmento crescente máximo é: %v\n", len(maiorCompEncontrado)-1)
	}

}
