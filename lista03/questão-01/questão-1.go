package main

import "fmt"

func main() {
	var n, M, num int

	for {
		fmt.Scan(&n)
		if n > 1 && n < 10000 {
			break
		} else {
			fmt.Println("O número precisa ser > 1 e < 10000")
		}
	}
	numeros := make([]int, n)
	resultados := []string{}
	for i := 0; i < n; i++ {
		fmt.Scan(&numeros[i])
	}

	for {
		fmt.Scan(&M)
		if M > 1 && M < 1000 {
			break
		} else {
			fmt.Println("O número precisa ser > 1 e < 1000")
		}
	}
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
		}
	}
	return achou
}
