package main

import (
  "fmt"
  "math"
  )

func main() {
  var n int
  var dist float64
  
  fmt.Scan(&n)
  pontos := make([]float64, n*3)
  
  for i := 0; i < n*3; i++ {
    fmt.Scan(&pontos[i])
  }
  
  for i := 0; i < len(pontos) - 4; i+=3 {
  dist = math.Sqrt(math.Pow(pontos[i] - pontos[i+3], 2) + math.Pow(pontos[i+1] - pontos[i+4], 2) + math.Pow(pontos[i+2] - pontos[i+5], 2))
  fmt.Printf("%.2f\n", dist)
  }
}