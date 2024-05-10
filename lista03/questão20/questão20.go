package main

import (
  "fmt"
  "math"
  )
  
func main() {
  var N, c int
  fmt.Scan(&N)
  if N < 2 || N > 1000 {
    fmt.Println("O número de pontos deve estar no intervalo de 2 a 1000")
    return
  }
  
  pontos := make([]float64, N*3)
  vetoresMaior := make([]float64, N-1)
  
  for i := 0; i < N*3; i++ {
    fmt.Scan(&pontos[i])
  }
  
  for i := 0; i < len(pontos)-3; i++ {
    vetor := math.Abs(pontos[i] - pontos[i+3])
    if i % 3 == 0 && i != 0 {
      c++
    }
    if vetor > vetoresMaior[c] {
      vetoresMaior[c] = vetor
    }
  }
  for i := range vetoresMaior{
    fmt.Printf("%.2f\n", vetoresMaior[i])
  }
}