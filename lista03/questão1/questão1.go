package main

import "fmt"

func main() {
	var n, M, num int

	fmt.Scan(&n)
	numeros := make([]int, n)
	resultados := []string{}
	for i := 0; i < n; i++ {
		fmt.Scan(&numeros[i])
	}

	fmt.Scan(&M)
	for i := 0; i < M; i++ {
		fmt.Scan(&num)
		achou := encontrar(num, n, numeros)
		if achou {
			resultados = append(resultados, "ACHEI")
		} else {
			resultados = append(resultados, "NAO ACHEI")
		}
	}
	for i := range resultados {
		fmt.Println(resultados[i])
	}
}

func encontrar(n int, t int, s []int) bool {
	achou := false
	for i := 0; i < t; i++ {
		if n == s[i] {
			achou = true
			break
		} else {
			achou = false
		}
	}
	return achou
}
