package main

import (
	"fmt"
)

var n, cont int
var times []int

func main() {
	fmt.Println("Digite a quantidade de times para descobrir as possíveis finais do campeonato.")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		times = append(times, i)
	}
	if n <= 1 {
		fmt.Println("Campeonato inválido!")
	}
	permutarJogos(1, 1)
}

func permutarJogos(i int, anterior int) {
	for ; i <= len(times)-anterior; i++ {
		cont++
		fmt.Printf("Final %v: Time%v X Time%v\n", cont, times[anterior-1], times[len(times)-i])
		if len(times)-i == anterior {

			if anterior-1 == (len(times) - 1/2) {
				return
			}
			permutarJogos(1, anterior+1)
		}
	}
}
