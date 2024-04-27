package main

import "fmt"

func main() {
	var n, soma int
	impares := []int{}

	fmt.Scan(&n)

	for i := n + 1; i < n*n*n; i++ {
		if i%2 != 0 {
			impares = append(impares, i)
		}
	}

	fmt.Println(impares)
}
