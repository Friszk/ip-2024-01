package main

import (
	"fmt"
)

var (
	p1, p2, p3, p4, p5, p6, p7, p8, e1, e2, e3, e4, e5, t, p, nF float64
	sF                                                           string
	matricula                                                    int
)

func main() {

	fmt.Println("Digite a matrícula do aluno, as 8 notas das provas, as 5 notas das listas de exercícios, a nota do trabalho final e por fim o valor da presença do aluno")

	for i := 0; i < 1; i-- {
		fmt.Scan(&matricula)

		if matricula < 0 {
			return
		}

		fmt.Scan(&p1, &p2, &p3, &p4, &p5, &p6, &p7, &p8, &e1, &e2, &e3, &e4, &e5, &t, &p)

		nF = (0.7*((p1+p2+p3+p4+p5+p6+p7+p8)/8) + 0.15*((e1+e2+e3+e4+e5)/5) + 0.15*(t))

		if nF >= 6 && p > (128*0.75) {
			sF = "APROVADO"
		} else if nF < 6 && p > (128*0.75) {
			sF = "REPROVADO POR NOTA"
		} else if nF >= 6 && p < (128*0.75) {
			sF = "REPROVADO POR FREQUÊNCIA"
		} else {
			sF = "REPROVADO POR NOTA E FREQUÊNCIA"
		}

		fmt.Printf("Matrícula: %v, Nota final: %.2f, Situação final: %v\n", matricula, nF, sF)
	}

}
