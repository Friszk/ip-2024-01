package main

import "fmt"

func main() {
	var n int

	for {
		fmt.Scan(&n)
		if n < 5000 {
			break
		} else {
			fmt.Println("O tamanho da lista deve ser menor que 5000")
		}
	}
	if n < 0 {
		return
	}

	numeros := make([]int, n)

	for i := range numeros {
		fmt.Scan(&numeros[i])
	}

	for i := range numeros {
		fmt.Printf("%v ", numeros[i])
	}
}
