package main

import (
	"fmt"
)

var num int
var quantN int

func main() {
	fmt.Println("Digite um número par e a quantidade de pares a serem exibidos")
	fmt.Scan(&num)

	if num%2 != 0 {
		fmt.Println("O PRIMEIRO NÚMERO NÃO É PAR")
	} else {
		fmt.Scan(&quantN)

		for i := 0; i < quantN; i++ {
			fmt.Printf("%v ", num)
			num = num + 2
		}
		fmt.Println()
	}

}
