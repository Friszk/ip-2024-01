package main

import "fmt"

func main() {
	var N, K, aTempo int
	fmt.Scan(&N, &K)
	if N < 0 || N > 1000 {
		fmt.Println("O número de alunos deve estar no intervalo de 0 a 1000")
		return
	}

	tempoAlunos := make([]int, N)
	adiantados := []int{}

	for i := range tempoAlunos {
		fmt.Scan(&tempoAlunos[i])
		if tempoAlunos[i] <= 0 {
			aTempo++
			adiantados = append(adiantados, i+1)
		}
	}

	if aTempo >= K {
		fmt.Println("NAO")
		for i := len(adiantados) - 1; i >= 0; i-- {
			fmt.Println(adiantados[i])
		}
	} else {
		fmt.Println("SIM")
	}

}
