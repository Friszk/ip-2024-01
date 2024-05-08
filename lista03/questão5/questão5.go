package main

import "fmt"

func main() {
	var N, maior, tempI int
	resultados := []int{}

	for {
		fmt.Scan(&N)
		if N < 1 || N > 10000 {
			fmt.Println("O número deve estar no intervalo de 1 a 10000")
			return
		}
		if N == 0 {
			break
		}
		numeros := make([]int, N)
		for i := range numeros {
			fmt.Scan(&numeros[i])
		}
		for i, v := range numeros {
			if v > maior {
				maior = v
				tempI = i
			}
		}
		resultados = append(resultados, tempI, maior)
		maior = 0
		tempI = 0
	}

	for i := 0; i < len(resultados); i += 2 {
		fmt.Println(resultados[i], resultados[i+1])
	}
}
