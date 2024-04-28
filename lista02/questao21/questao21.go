package main

import "fmt"

func main() {
  
  var n, num float64
  var ordenado bool
  numeros := []float64{}
  ordens := []string{}
  
  fmt.Println("Digite o tamanho da lista e em seguida a lista")
  
  for {
    numeros = nil
    ordenado = false
    fmt.Scan(&n)
    for i := 0; i < int(n); i++ {
      fmt.Scan(&num)
      numeros = append(numeros, num)
    }
    for i := 0; i < int(n-1); i++ {
      if numeros[i] < numeros[i+1] {
        ordenado = true
      } else {
        ordenado = false
        break
      }
    }
    if ordenado {
    ordens = append(ordens, "ORDENADA")
    } else {ordens = append(ordens, "DESORDENADA")}
    if n == 0 {break}
  }
  
  for i:=0; i < len(ordens) - 1; i++ {
    fmt.Println(ordens[i])
  }
}