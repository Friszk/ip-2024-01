package main

import "fmt"
import "math"

func main() {
  var x, N, somatorio float64
  
  fmt.Scan(&x, &N)
  
  for i := 0.0; i <= N; i++ {
    denominador := 1.0
    for a := 1.0; a <= i; a++ {
      denominador *= a
    }
    somatorio += (math.Pow(x, i)/denominador)
  }
  fmt.Printf("eˆ%.2f = %.6f\n", x, somatorio)
}