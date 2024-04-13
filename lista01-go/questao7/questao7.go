package main

import "fmt"

var tempF, chuvaP, tempC, chuvaM float32

func main() {

	fmt.Println("Digite a temperatura em Fahrenheit e a quantidade de chuva em polegadas a serem calculadas")
	fmt.Scan(&tempF, &chuvaP)

	tempC = (5*tempF - 160) / 9
	chuvaM = (chuvaP * 25.4)

	fmt.Printf("VALOR EM CELSIUS = %.2f\nA QUANTIDADE DE CHUVA É = %.2f\n", tempC, chuvaM)

}
