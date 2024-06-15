package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(inverterArray(array))
}

func inverterArray(a []int) []int {
	var i int
	var e *int
	e = &i
	if *e == len(a) {
		return a
	}
	a = append(a[1:], a[:1]...)
	*e++
	return inverterArray(a)
}
