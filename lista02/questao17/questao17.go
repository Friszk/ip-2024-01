package main

import (
	"fmt"
)

func main() {
  var codMercadoria, nVMercadorias int
  var pCMercadoria, pVMercadoria float64
  var tudo []interface{}
  quantLucroMenor10 := 0
  quantLucroMenor20 := 0
  quantLucroMaior20 := 0
  
  for {
    fmt.Scan(&codMercadoria)
    if codMercadoria == 0 {
      break
    }
    fmt.Scan(&pCMercadoria, &pVMercadoria, &nVMercadorias)
    tudo = append(tudo, codMercadoria, pCMercadoria, pVMercadoria, nVMercadorias)
    //Lucro
    tudo = append(tudo, ((pVMercadoria*float64(nVMercadorias)*100)/(pCMercadoria*float64(nVMercadorias))-100))
  }
  
  for i := 3; i < len(tudo); i =+ 3{
    num := tudo[i].(int)
    if float64(num) < 10.0{
      quantLucroMenor10++
    } else if float64(num) < 20 {
      quantLucroMenor20++
    } else {
      quantLucroMaior20++
    }
  }
  
  fmt.Printf("Quantidade de mercadorias que geraram lucro menor que 10%: %v\n", quantLucroMenor10)
  fmt.Printf("Quantidade de mercadorias que geraram lucro maior ou igual a 10% e menor ou igual a 20%: %v\n", quantLucroMenor20)
  fmt.Printf("Quantidade de mercadorias que geraram lucro maior do que 20%: %v\n", quantLucroMaior20)
//  "Quantidade de mercadorias que geraram lucro maior ou igual a 10% e menor ou igual a 20%: %v\nQuantidade de mercadorias que geraram lucro maior do que 20%: %v\nCodigo da mercadoria que gerou maior lucro: %vCodigo da mercadoria mais vendida: %v\nValor total de compras: %v, valor total de vendas: %v e percentual de lucro total: %v",)
}
