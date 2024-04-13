package main

import (
	"fmt"
)

var nJogos, ingQ, valorTotal, pP, pG, pA, pC float64

func main() {
	fmt.Print("Digite o número de jogos a serem testados\n")
	fmt.Scan(&nJogos)

	for n := 0; n < int(nJogos); n++ {
		fmt.Println("Digite a quantidade de ingressos e as porcentagens equivalentes a cada categoria")
		fmt.Scan(&ingQ, &pP, &pG, &pA, &pC)
		valorTotal = ingQ * (pP/100*1 + pG/100*5 + pA/100*10 + pC/100*20)
		fmt.Printf("A RENDA DO JOGO %v É = %.2f\n", n+1, valorTotal)
	}
}
