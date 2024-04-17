package main

import (
	"fmt"
)

var A, B float64
var a = 1

func main() {

	fmt.Println("Digite a população dos países\nO primeiro deve ter uma quantidade de habitantes menor que o segundo")
	fmt.Scan(&A, &B)
	if A >= B {
		fmt.Println("O primeiro país possuí uma quantidade maior ou igual")
		return
	}

	for A <= B {
		A += A * 0.015
		a++
	}

	fmt.Printf("ANOS = %v\n", a)

}
