package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(inverterArray(array, len(array)))
}

func inverterArray(a []int, i int) []int {
	if i == int(len(a)/2) {
	    return a
	}
	a[len(a)-i], a[i-1] = a[i-1], a[len(a) -i]
	i--
	return inverterArray(a, i)
}
