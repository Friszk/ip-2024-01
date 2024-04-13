package main

import "fmt"

var numero int

func main() {

	fmt.Println("Digite um número e descubra se ele é divisível por 3 e por 5 ao mesmo tempo")
	fmt.Scan(&numero)

	//POR 3
	if numero%3 != 0 || numero%5 != 0 {
		fmt.Println("O número não é divisível")
	} else {
		fmt.Println("O número é divisível")
	}

}
