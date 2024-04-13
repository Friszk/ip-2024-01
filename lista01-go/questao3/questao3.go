package main

import (
	"fmt"
	"strconv"
)

var n1, n2, n3 int

func main() {
	fmt.Println("Digite os 3 números")
	fmt.Scan(&n1, &n2, &n3)
	if n1 == 0 {
		fmt.Println("O primeiro dígito não pode ser 0")
		return
	}
	s := fmt.Sprintf("%v%v%v", n1, n2, n3)
	n, _ := strconv.Atoi(s)
	fmt.Println(n, n*n)
}
