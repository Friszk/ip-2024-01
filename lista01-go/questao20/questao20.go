package main

import (
	"fmt"
)

var horas, minutos, segundos, total int

func main() {

	fmt.Println("Digite as horas, os minutos e os segundos, respectivamente")
	fmt.Scan(&horas, &minutos, &segundos)
	total = (horas * 60 * 60) + (minutos * 60) + segundos
	fmt.Printf("O TEMPO EM SEGUNDOS É %v\n", total)
}
