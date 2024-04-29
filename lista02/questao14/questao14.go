package main

import (
	"fmt"
)

func main() {
	var linhas, colunas int

	fmt.Println("Digite o número de linhas e o número de colunas")
	fmt.Scan(&linhas, &colunas)

	matriz := make([][]string, linhas)
	for i := range matriz {
		matriz[i] = make([]string, colunas)
	}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			matriz[i][j] = fmt.Sprintf("(%d,%d)", i+1, j+1)
		}
	}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			if i > j {
				if j > 0 {
					fmt.Print(" - ")
				}
				fmt.Printf("%s", matriz[i][j])
			}
		}
		fmt.Println()
	}
}
