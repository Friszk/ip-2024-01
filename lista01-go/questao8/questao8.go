package main

import "fmt"

var raio, altura, areaTotal, compCir, areaB float32

func main() {

	fmt.Println("Entre com a medida do raio e da altura da lata, em metros")
	fmt.Scan(&raio, &altura)

	compCir = 2 * 3.14159 * raio
	areaB = 3.14159 * raio * raio
	areaTotal = 2*areaB + (compCir * altura)

	fmt.Printf("O VALOR DO CUSTO É = %.2f\n", areaTotal*100)
}
