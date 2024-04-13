package main

import "fmt"

var salario, kW, prctg, custoporKw, custodConsumo, custocDesconto float32

func main() {
	fmt.Println("Digite o salário mínimo e a quantidade de kW gastos")
	fmt.Scan(&salario, &kW)
	prctg = (kW * 70) / 10000
	custodConsumo = prctg * salario
	custoporKw = custodConsumo / kW
	custocDesconto = custodConsumo * 0.9
	fmt.Printf("Custo por kW: R$%.2f\nCusto do consumo: R$%.2f\nCusto com desconto: R$%.2f\n", custoporKw, custodConsumo, custocDesconto)
}
