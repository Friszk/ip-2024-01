package main

import (
	"fmt"
)

var salario, novoSalario float32

func main() {

	fmt.Println("Digite seu salário para descobrir o valor com reajuste")
	fmt.Scan(&salario)

	if salario > 350 {
		novoSalario = salario * 1.3
	} else {
		novoSalario = salario * 1.5
	}
	fmt.Printf("SALÁRIO COM REAJUSTE = %.2f\n", novoSalario)
}
