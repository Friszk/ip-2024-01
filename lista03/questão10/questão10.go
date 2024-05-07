package main

import "fmt"

func main() {
  var n, ultimo, c, maior, indMaior int
  fmt.Scan(&n)
  
  numeros := make([]int, n)
  for i := 0; i < n; i ++ {
    fmt.Scan(&numeros[i])
    if numeros[i] > maior {
      maior = numeros[i]
      indMaior = i
    }
    ultimo = numeros[i]
  }
  for i := range numeros {
    if ultimo == numeros[i] {
      c ++
    }
  }
  fmt.Printf("Nota %v, %v vezes\nNota %v, índice %v\n", ultimo, c, maior, indMaior)
}