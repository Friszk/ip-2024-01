package main

import (
	"fmt"
)

func main() {
  var codMercadoria, nVMercadorias, maiorLucro, maiorVenda, maiorValorVenda int
  var pCMercadoria, pVMercadoria, valorTotalCompra, ValorTotalVenda, LucroTotal float64
  var tudo []interface{}
  quantLucroMenor10 := 0
  quantLucroMenor20 := 0
  quantLucroMaior20 := 0
  
  
  for{
    fmt.Scan(&codMercadoria)
    if codMercadoria == 0 {
      break
    }
    fmt.Scan(&pCMercadoria, &pVMercadoria, &nVMercadorias)
    tudo = append(tudo, codMercadoria, pCMercadoria, pVMercadoria, nVMercadorias)
    tudo = append(tudo, ((pVMercadoria*float64(nVMercadorias)*100)/(pCMercadoria*float64(nVMercadorias))-100))
  }
  
  //Maior prctg de lucro
  for i := 4; i < len(tudo); i += 5{
    num := tudo[i].(float64)
      if num*100+100*float64(tudo[i-1].(int)) > float64(maiorLucro) {
        maiorLucro = tudo[i-4].(int)
      }
    if num < 10.0{
      quantLucroMenor10++
    } else if num <= 20 {
      quantLucroMenor20++
    } else {
      quantLucroMaior20++
    }
  }
  //Maior lucro mercadoria
  for i:= 3; i < len(tudo); i +=5{
    if tudo[i].(int) > maiorValorVenda {
      maiorValorVenda = tudo[i].(int)
      maiorVenda = tudo[i-3].(int)
    }
  }
  
  for i:= 1; i < len(tudo); i +=5 {
    valorTotalCompra += tudo[i].(float64)*float64(tudo[i+2].(int))
  }
  
  for i:= 2; i < len(tudo); i +=5 {
    ValorTotalVenda += tudo[i].(float64)*float64(tudo[i+1].(int))
  }
  
  LucroTotal = (ValorTotalVenda*100/valorTotalCompra)-100
  
  fmt.Printf("Quantidade de mercadorias que geraram lucro menor que 10%%: %v\n", quantLucroMenor10)
  fmt.Printf("Quantidade de mercadorias que geraram lucro maior ou igual a 10%% e menor ou igual a 20%%: %v\n", quantLucroMenor20)
  fmt.Printf("Quantidade de mercadorias que geraram lucro maior do que 20%%: %v\n", quantLucroMaior20)
  fmt.Printf("Código da mercadoria que gerou maior lucro: %v\n", maiorLucro)
  fmt.Printf("Codigo da mercadoria mais vendida: %v\n", maiorVenda)
  fmt.Printf("Valor total de compras: %.2f, valor total de vendas: %.2f e percentual de lucro total: %.2f\n", valorTotalCompra, ValorTotalVenda, LucroTotal)
}
