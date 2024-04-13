package main

import "fmt"

var quantTemp int
var tempF, tempC float32

func main() {

	fmt.Println("Por favor, digite a quantidade de temperaturas a serem calculadas e, em seguida,\nas temperaturas em Fahrenheit")
	fmt.Scan(&quantTemp)

	for i := 0; i < quantTemp; i++ {
		fmt.Scan(&tempF)
		tempC = (5 * (tempF - 32)) / 9
		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", tempF, tempC)
	}

}
