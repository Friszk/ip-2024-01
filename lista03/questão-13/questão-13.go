package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	if n < 1 || n > 1000000 {
		fmt.Println("O número deve estar no intervalo de 1 a 1000000")
		return
	}

	lista := make([]int, n)

	for i := range lista {
		fmt.Scan(&lista[i])
	}

	maiorEnc := 0
	numMaiorEnc := 0

	for i := range lista {
		maiorAtual := 0
		for c := range lista {
			if lista[i] == lista[c] {
				maiorAtual++
			}
		}
		if maiorAtual > maiorEnc {
			maiorEnc = maiorAtual
			numMaiorEnc = lista[i]
		} else if maiorAtual == maiorEnc {
			if lista[i] < numMaiorEnc {
				numMaiorEnc = lista[i]
			}
		}
	}

	fmt.Printf("%v\n%v\n", numMaiorEnc, maiorEnc)

}
