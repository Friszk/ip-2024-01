package main

import (
	"fmt"
)

func main() {
	var matricula, horas int
	var porhora, total float64
	resultados := []float64{}

	fmt.Println("Digite o número da matrícula do funcionário, suas horas trabalhadas e o valor recebido por hora")
	for i := 0; i < 1; {
		fmt.Scan(&matricula)
		if matricula == 0 {
			break
		}
		fmt.Scan(&horas, &porhora)
		total = float64(horas) * porhora
		resultados = append(resultados, float64(matricula))
		resultados = append(resultados, float64(total))
	}
	if matricula == 0 {
		for i := 0; i < len(resultados)-1; {
			fmt.Printf("%.f %.2f\n", resultados[i], resultados[i+1])
			i += 2
		}
		return
	}

}
