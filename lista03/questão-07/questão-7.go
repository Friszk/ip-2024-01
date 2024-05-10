package main

import "fmt"

func main() {
	var n, maior, contador int
	resultados := []int{}
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		if n > 10000 || n < 1 {
			fmt.Println("O número deve estar no intervalo de 1 a 10000")
			return
		}
		numeros := make([]int, n)
		for i := range numeros {
			fmt.Scan(&numeros[i])
			if numeros[i] > maior {
				maior = numeros[i]
			}
		}
		encontrarMaior(maior, numeros, &resultados, contador)
		maior = 0
	}
	for i := 0; i < len(resultados); i += 2 {
		fmt.Printf("(%v) %v\n", resultados[i], resultados[i+1])
	}
}

func encontrarMaior(maior int, lista []int, result *[]int, cont int) {
	for i := 0; i <= maior; i++ {
		for in := range lista {
			if i >= lista[in] {
				cont++
			}
		}
		*result = append(*result, i)
		*result = append(*result, cont)
		cont = 0
	}
}
