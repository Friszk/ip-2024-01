package main

import "fmt"

func main() {
  var n int
  resultados := []int{}
  for{
    _, erro := fmt.Scanln(&n)
    if erro != nil {
      break
    }
    resultados = append(resultados, n)
  }
  for i := range resultados {
    fmt.Printf("%b\n", resultados[i])
  }
}