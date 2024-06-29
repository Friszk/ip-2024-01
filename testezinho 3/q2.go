package main

import "fmt"

func main() {
	var nT, tV int
	resultados := [][]int{}

	fmt.Scan(&nT)
	for i := 0; i < nT; i++ {
		fmt.Scan(&tV)
		v := make([]int, tV)

		for i := range v {
			fmt.Scan(&v[i])
		}
		resultados = append(resultados, trocaOpostosSeMenor(v, len(v)))
	}
	for i := range resultados {
		fmt.Println(resultados[i])
	}
}

func trocaOpostosSeMenor(s []int, n int) []int {
	for i := 0; i < n/2; i++ {
		x := i
		y := len(s) - 1 - i
		if s[x] < s[y] {
			troca(s, x, y)
		}
	}
	return s
}

func troca(s []int, x, y int) []int {
	s[x], s[y] = s[y], s[x]
	return s
}
