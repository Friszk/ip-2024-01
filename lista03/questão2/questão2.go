package main

import "fmt"

func main() {
	var n, maior, maiores int

	for {
		fmt.Scan(&n)
		if n >= 1 && n <= 1000 {
			break
		} else {
			fmt.Println("O número precisa ser > 1 e < 1000")
		}
	}

	numeros := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&numeros[i])
	}

	fmt.Scan(&maior)

	for i := 0; i < n; i++ {
		if numeros[i] >= maior {
			maiores++
		}
	}

	fmt.Println(maiores)
}
