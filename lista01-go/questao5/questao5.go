package main

import "fmt"

var m3gasto, conta float32
var numConta int
var tipo string

func main() {
	fmt.Println("Digite o número da conta, os metros cúbicos de água gastos e o tipo de conta;\nC = COMERCIAL, I = INDUSTRIAL, R = RESIDENCIAL")
	fmt.Scan(&numConta, &m3gasto, &tipo)

	if m3gasto < 0 {
		fmt.Println("METROS INVÁLIDOS")
		return
	}
	if tipo == "R" {
		conta = 5 + (m3gasto * 0.05)
	} else if tipo == "C" {
		if m3gasto > 80 {
			conta = 500 + ((m3gasto - 80) * 0.25)
		} else {
			conta = 500
		}
	} else if tipo == "I" {
		if m3gasto > 100 {
			conta = 800 + ((m3gasto - 100) * 0.04)
		} else {
			conta = 800
		}
	} else {
		fmt.Println("TIPO DE CONTA INVÁLIDA")
		return
	}

	fmt.Printf("CONTA = %d\nVALOR DA CONTA = %.2f\n", numConta, conta)

}
