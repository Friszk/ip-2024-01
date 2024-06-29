package main

import "fmt"

func main() {
	nums := []int{9, 7, 6, 5, 4, 3, 2, 1, 1, 0}
	fmt.Println(Ordenar(nums))
}

func Ordenar(s []int) []int {
	for {
		trocou := false
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				trocou = true
			}
		}
		if !trocou {
			break
		}
	}
	return s
}
