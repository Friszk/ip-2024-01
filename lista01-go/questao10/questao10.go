package main

import "fmt"

var a, b, c, d, detMatriz float32

func main() {

	fmt.Println("Digite os quatro elementos a, b, c, d da matriz quadrada bidimensional")
	fmt.Scan(&a, &b, &c, &d)

	detMatriz = a*d - b*c
	fmt.Printf("O VALOR DO DETERMINANTE É = %.2f\n", detMatriz)
}
