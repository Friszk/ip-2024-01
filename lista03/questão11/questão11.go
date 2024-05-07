package main

import "fmt"

func main() {
  var n, maior int
  fmt.Scan(&n)
  
  vetor := make([]int, n)
  
  for i := range vetor {
    fmt.Scan(&vetor[i])
    if vetor[i] > maior {
      maior = vetor[i]
    }
  }
  menor := maior
  fmt.Println()
  for i := range vetor {
    fmt.Printf("%v ",vetor[i])}; fmt.Println()
  for i := n-1; i >= 0; i-- {
    if vetor[i] < menor {
      menor = vetor[i]
    }
    fmt.Printf("%v ", vetor[i])} ; fmt.Println()
    fmt.Println(maior)
    fmt.Println(menor)
}