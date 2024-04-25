package main

import (
	"fmt"
)

func main() {
	var n, r int
	fmt.Println("Digite o número de grãos de milho a ser colocado no primeiro quadro e descubra a quantidade final.")
	fmt.Scan(&n)
	r = 33*n + 31*n*2
	fmt.Println(r)
}
