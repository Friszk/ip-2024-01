package main

import (
	"fmt"
)

var nota1, nota2, nota3, m float64

func main() {
	fmt.Println("Escreva as três notas")
	fmt.Scan(&nota1, &nota2, &nota3)
	if nota1 < 0 || nota2 < 0 || nota3 < 0 {
		fmt.Print("Nota digitada inválida\n")
		return
	}
	m = ((nota1 + nota2 + nota3) / 3)
	fmt.Printf("A média do aluno é %.2f\n", m)
}
