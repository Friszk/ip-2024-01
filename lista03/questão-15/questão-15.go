package main

import (
	"fmt"
	"math"
)

func main() {
	var T, N int
	menorDis := 2001
	resultados := []int{}

	fmt.Scan(&T)
	if T < 1 || T > 10 {
		fmt.Println("A quantidade de testes deve estar no intervalo de 1 a 10")
		return
	}

	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		if N < 2 || N > 1000 {
			fmt.Println("O número deve estar no intervalo de 2 a 1000")
			return
		}

		lista := make([]int, N)
		for i := range lista {
			fmt.Scan(&lista[i])
		}

		e := 0
		a := e + 1
		testes := 0
		for e < N {
			for ; a < N; a++ {
				if menorDis > int(math.Abs(float64(lista[e]-lista[a]))) {
					menorDis = int(math.Abs(float64(lista[e] - lista[a])))
				}
				testes++
			}
			e++
			a = e + 1
		}
		resultados = append(resultados, menorDis, testes)
		menorDis = 2001
		testes = 0
	}
	fmt.Println()
	for i := 0; i < len(resultados)-1; i += 2 {
		fmt.Println(resultados[i], resultados[i+1])
	}
}
