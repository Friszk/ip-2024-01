package main

import (
	"fmt"
	"math"
)

var altura, aresta, areaB, areaTotal float64

func main() {

	fmt.Println("Digite a altura da pirâmide e o tamanho da aresta da base, em metros:")
	fmt.Scan(&altura, &aresta)

	areaB = (3 * aresta * aresta * math.Sqrt(3)) / 2
	areaTotal = (areaB * altura) / 3

	fmt.Printf("O VOLUME DA PIRÂMIDE É %.2f METROS CÚBICOS\n", areaTotal)
}
