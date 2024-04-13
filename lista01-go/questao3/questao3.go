package main

import (
	"fmt"
	"strconv"
)

var n1, n2, n3 int

func main() {
	fmt.Println("Digite os 3 números")
	fmt.Scan(&n1, &n2, &n3)
	if n1 > 0 && n1 < 10 || n2 > 0 && n2 < 10 || n3 > 0 && n3 < 10 {
		s := fmt.Sprintf("%v%v%v", n1, n2, n3)
		n, _ := strconv.Atoi(s)
		fmt.Println(n, n*n)
	} else {
	fmt.Println("O primeiro dígito não pode ser 0")
	return
}
}
