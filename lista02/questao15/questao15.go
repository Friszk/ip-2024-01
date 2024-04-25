package main

import (
	"fmt"
)

func main() {

	var (
		T, n, u, d, c, u2, d2, c2, p int
	)

	fmt.Println("Digite o número de testes a serem feitos e, em seguida, os pares de números a serem comparados")
	fmt.Scan(&T)
	numeros := []int{}
	for i := 1; i <= T*2; i++ {
		fmt.Scan(&n)
		numeros = append(numeros, n)
	}
	fmt.Println("")

	for cont := 1; cont <= T; cont++ {
		u = numeros[p] % 10
		d = numeros[p] % 100 / 10
		c = numeros[p] / 100
		u2 = numeros[p+1] % 10
		d2 = numeros[p+1] % 100 / 10
		c2 = numeros[p+1] / 100
		p += 2

		if u > u2 {
			fmt.Println(u*100 + d*10 + c)
		} else if u == u2 {
			if d > d2 {
				fmt.Println(u*100 + d*10 + c)
			} else if d == d2 {
				if c > c2 {
					fmt.Println(u*100 + d*10 + c)
				}
			} else {
				fmt.Println(u2*100 + d2*10 + c2)
			}
		}
	}
}
