package main

import "fmt"

func main() {
	var x, n int
	fmt.Scan(&x, &n)
	fmt.Println(elevar(x, n))
}

func elevar(x, n int) int {
	if n == 1 {
		return x
	} else {
		return x * elevar(x, n-1)
	}
}
