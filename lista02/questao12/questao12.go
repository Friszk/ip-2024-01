package main

import (
	"fmt"
)

var (
	ValorFinal,
	ValorIngresso,
	ValorInicial,
	Lucro,
	nIngressos float64

	maiorAtual,
	maiorEnc [3]float64
)

func main() {
	fmt.Println("Digiteo valor do ingresso, o valor inicial e o valor final")
	fmt.Scan(&ValorIngresso, &ValorInicial, &ValorFinal)

	for ValorAtual := ValorInicial; ValorAtual <= ValorFinal; ValorAtual++ {
		if ValorAtual <= ValorIngresso {
			nIngressos = (ValorIngresso-ValorAtual)*2*25 + 120

		} else {
			nIngressos = (ValorIngresso-ValorAtual)*2*30 + 120
		}

		Lucro = ValorAtual*nIngressos - (200 + (nIngressos * 0.05))

		fmt.Printf("V:%.2f, N:%v, L:%.2f\n", ValorAtual, nIngressos, Lucro)

		maiorAtual[0] = ValorAtual
		maiorAtual[1] = Lucro
		maiorAtual[2] = nIngressos

		if maiorAtual[1] > maiorEnc[1] {
			maiorEnc[0] = maiorAtual[0]
			maiorEnc[1] = maiorAtual[1]
			maiorEnc[2] = maiorAtual[2]
		}
	}

	fmt.Printf("Melhor valor final: %.2f\nLucro: %.2f\nNúmero de ingressos: %v\n", maiorEnc[0], maiorEnc[1], maiorEnc[2])

}
