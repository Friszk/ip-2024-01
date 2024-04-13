package main

import "fmt"

var a, b, c, resultado float32

func main() {

	fmt.Println("Digite os coeficientes A, B e C da equação do 2° grau a ser calculada")
	fmt.Scan(&a, &b, &c)

	resultado = (b * b) - (4 * a * c)
	fmt.Printf("O VALOR DE DELTA E = %.2f\n", resultado)
}
