package main

import (
	"fmt"
)

func main() {
	var acertos int
	sorteado := make([]int, 6)
	for i := range sorteado {
		fmt.Scan(&sorteado[i])
	}

	n := 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		aposta := make([]int, 6)
		for i := range aposta {
			fmt.Scan(&aposta[i])
		}
		for i := range aposta {
			for j := range sorteado {
				if aposta[i] == sorteado[j] {
					acertos++
					break
				}
			}
		}
		fmt.Println(acertos)
		acertos = 0
	}

}
