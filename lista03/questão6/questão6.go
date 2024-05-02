package main

import "fmt"

func main() {
  var n, soma int
  fmt.Scan(&n)
  
  numeros := make([]int, n)
  
  for i := range numeros{
    fmt.Scan(&numeros[i])
  }
  
  for _, v := range numeros {
    soma += v
  }
  
  fmt.Println(soma)
}