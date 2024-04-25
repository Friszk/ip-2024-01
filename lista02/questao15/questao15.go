package main

import (
	"fmt"
)

func main() {

	var T, n int

	fmt.Scan(&T)
	numeros := []int{}
	for i := 0; i < T*2; i++ {
		fmt.Scan(&n)
		numeros = append(numeros, n)
	}

	fmt.Println(numeros)
}
