package main

import "fmt"

func main() {
  var n, maior, tempI int
  resultados := []int{}
  
  for {
    fmt.Scan(&n)
    if n == 0 {break}
    numeros := make([]int, n)
    for i := range numeros{
      fmt.Scan(&numeros[i])
    }
    for i, v := range numeros {
      if v > maior {
        maior = v
        tempI = i
      }
    }
    resultados = append(resultados, tempI, maior)
      maior = 0
      tempI = 0
  }
  
  for i := 0; i < len(resultados); i += 2 {
    fmt.Println(resultados[i], resultados[i+1])
  }
}