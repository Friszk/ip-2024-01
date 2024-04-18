package main

import (
	"fmt"
)

var sp, si float64

func main() {

	fmt.Println("Digite a sequência de números inteiros diferentes de zero para ser calculada a média dos números pares e ímpares.")
	Seq := []int{}
	pares := []float64{}
	impares := []float64{}

	for i := 0; i > -1; i++ {
		var num int
		fmt.Scan(&num)
		Seq = append(Seq, num)
		if Seq[i] == 0 {
			break
		}

		if Seq[i]%2 == 0 {
			pares = append(pares, float64(Seq[i]))
		} else {
			impares = append(impares, float64(Seq[i]))
		}
	}

	for i := 0; i < len(pares); i++ {
		sp += (pares[i])
	}

	for i := 0; i < len(impares); i++ {
		si += (impares[i])
	}

	fmt.Printf("MEDIA PAR: %.2f\nMEDIA IMPAR: %.2f\n", sp/float64(len(pares)), si/float64(len(impares)))
}
