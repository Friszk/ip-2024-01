package main

import "fmt"

var horasDiv, horasRes, horas, valorTotal int

func main() {

	fmt.Println("Digite por quantas horas você usou a charrete")
	fmt.Scan(&horas)

	horasDiv = horas / 3
	horasRes = horas - (horasDiv * 3)
	valorTotal = (10 * horasDiv) + (5 * horasRes)

	fmt.Printf("O VALOR A PAGAR É = %v\n", valorTotal)

}
