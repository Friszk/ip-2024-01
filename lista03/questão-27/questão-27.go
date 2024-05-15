package main

import (
	"fmt"
)

func main() {
	var q1, q2 int
	fmt.Scan(&q1, &q2)
	if q1 < 0 || q1 > 500000 || q2 < 0 || q2 > 500000 {
	  fmt.Println("q1 e q2 devem estar no intervalo de 1 a 500000")
	  return
	}
	V1 := make([]int, q1)
	V2 := make([]int, q2)
	for i := range V1 {
		fmt.Scan(&V1[i])
	}
	for i := range V2 {
		fmt.Scan(&V2[i])
	}
	Vr := intercalarOrdenar(V1, V2)
	for i := range Vr {
		fmt.Println(Vr[i])
	}
}

func intercalarOrdenar(s1, s2 []int) []int {
	Vr := []int{}
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			Vr = append(Vr, s1[i])
			i++
		} else {
			Vr = append(Vr, s2[j])
			j++
		}
	}
	for i < len(s1) {
		Vr = append(Vr, s1[i])
		i++
	}
	for j < len(s2) {
		Vr = append(Vr, s2[j])
		j++
	}
	return Vr
}