package main

import (
	"fmt"
)

var n, f int
var sequencia []int

func main() {
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		sequencia = append(sequencia, i)
	}
	permutar(sequencia, 0)

}

func trocar(x, y int) {
	var temp int

	temp = sequencia[x]
	sequencia[x] = sequencia[y]
	sequencia[y] = temp
}

func permutar(s []int, inicio int) {
	if inicio == len(s)-1 {
		fmt.Println(s)
		return
	}
	for i := inicio; i < len(s); i++ {
		trocar(inicio, i)
		permutar(s, inicio+1)
		trocar(inicio, i)
	}
}
