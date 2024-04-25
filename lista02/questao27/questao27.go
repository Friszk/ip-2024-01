package main

import (
	"fmt"
)

func main() {

	var n, i, primo int
	var eprimo bool
	fatDeN := []int{}
	primosAteN := []int{2}

	fmt.Println("Digite o número a ser fatorado")
	for {

		fmt.Scan(&n)
		if n <= 1 {
			fmt.Println("Fatoração não é possível para o número", n)
		} else {
			break
		}
	}

	nI := n
	for nAtual := 3; nAtual <= n; nAtual++ {
		for i = 2; i < nAtual; i++ {
			if nAtual%i == 0 {
				eprimo = false
				break
			} else {
				eprimo = true
			}
		}
		if eprimo {
			primo = i
			primosAteN = append(primosAteN, primo)
		}

	}

	for i := 0; n >= 2; {
		if n%primosAteN[i] == 0 {
			fatDeN = append(fatDeN, primosAteN[i])
			n = n / primosAteN[i]
		} else {
			i++
		}
	}

	fmt.Printf("%v = %v", nI, fatDeN[0])

	for i := 1; i <= len(fatDeN)-1; i++ {
		fmt.Printf(" x %v", fatDeN[i])
	}
	fmt.Println("")

}
