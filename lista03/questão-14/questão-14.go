package main

import "fmt"

func main() {
	var n int
	var mediana float64

	fmt.Scan(&n)
	if n < 0 || n > 1000000 {
		fmt.Println("O número deve estar no intervalo de 0 a 1000000")
		return
	}

	lista := make([]int, n)

	for i := range lista {
		fmt.Scan(&lista[i])
	}

	for {
		trocou := false
		for i := 0; i < n-1; i++ {
			if lista[i] > lista[i+1] {
				lista[i], lista[i+1] = lista[i+1], lista[i]
				trocou = true
			}
		}
		if !trocou {
			break
		}
	}
	if n%2 == 0 {
		mediana = float64((lista[(n/2)-1] + lista[(n/2)])) / 2
	} else {
		mediana = float64(lista[int(n/2)])
	}
	fmt.Printf("%.2f\n", mediana)
}
