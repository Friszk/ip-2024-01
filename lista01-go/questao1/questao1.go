package main

import (
	"fmt"
)

var nota1, nota2, nota3, m float64

func main() {
	fmt.Println("Escreva as três notas")
	fmt.Scan(&nota1, &nota2, &nota3)
	if nota1 < 0 || nota2 < 0 || nota3 < 0 {
		fmt.Println("Nota digitada inválida")
		return
	}
	m = ((nota1 + nota2 + nota3) / 3)
	fmt.Printf("A média do aluno é %.2f\n", m)
	if m < 6 {
		fmt.Print("REPROVADO\n")
	} else {
		fmt.Print("APROVADO\n")
	}
}
