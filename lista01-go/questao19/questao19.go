package main

import (
	"fmt"
)

var soma, num float64

func main() {

	fmt.Println("Digite o numero n, inteiro e positivo, a ser calculado o somatório")
	fmt.Scan(&num)

	if num < 1.0 {
		fmt.Println("Número inválido")
	} else {

		for i := 1.0; i <= num; i++ {
			soma = soma + (1 / i)
		}
	}

	fmt.Printf("%.6f\n", soma)
}
