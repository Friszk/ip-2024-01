package main

import "fmt"

func main() {
	array := []float64{1, 2, 3, 4, 5.1, 6, 7, 8.888888, 9, 10.2}
	fmt.Println(somaReais(array))
}

func somaReais(a []float64) []float64 {
	i := len(a) - 1
	if i == 0 {
		return a
	}
	a[0] += a[i]
	a = append(a[:i], a[i+1:]...)
	return somaReais(a)
}
