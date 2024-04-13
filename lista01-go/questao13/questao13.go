package main

import "fmt"

var nota float32
var conceito string

func main() {
	fmt.Println("Digite sua nota")
	fmt.Scan(&nota)

	if nota < 0 || nota > 10 {
		fmt.Println("Nota inválida")
		return
	}
	if nota >= 9 {
		conceito = "A"
	} else if nota >= 7.5 {
		conceito = "B"
	} else if nota >= 6 {
		conceito = "C"
	} else {
		conceito = "D"
	}

	fmt.Printf("NOTA = %v\nCONCEITO = %v\n", nota, conceito)

}
