package main

import (
	"fmt"
)

var n, f float64

func main() {

	fmt.Println("Digite o número a ser calculado seu fatorial")
	fmt.Scan(&n)

	f = n

	if n <= 0 {
		fmt.Printf("%v! = 1\n", n)
		return
	}

	for n := n - 1; n > 0; n-- {
		f = f * n
	}

	fmt.Printf("%v! = %.f\n", n, f)
}