package main

import (
	"fmt"
	"math"
)

func main() {

	var n int

	fmt.Println("Digite o valor da hipotenusa para descobrir quais são os catetos inteiros que podem formá-la")
	for {
		fmt.Scan(&n)
		if n > 0 {
			break
		} else {
			fmt.Println("O número deve ser maior que 0")
		}
	}

	for c := 1; c <= n; c++ {
		for b := 1; b < c; b++ {
			a := int(math.Sqrt(float64(c*c - b*b)))
			if b*b+a*a == c*c && a <= n && a > b {
				fmt.Printf("hipotenusa = %d, catetos %d e %d\n", c, b, a)
			}
		}
	}
}
