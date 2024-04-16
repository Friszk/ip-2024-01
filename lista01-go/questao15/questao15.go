package main

import (
	"fmt"
	"math"
)

var numQ, resultado, numero float64

func main() {

	fmt.Println("Digite o número a ser gerado seus quadrados pares")
	fmt.Scan(&numero)

	numQ = 2
	if numero <= 5 || numero >= 2000 {
		fmt.Println("Número fora do campo limite")
		return
	} else {
		for i := 2.0; i <= numero; i += 2 {
			resultado = math.Pow(i, numQ)
			fmt.Printf("%v ^ %v = %v\n", i, numQ, resultado)
		}
	}
}
