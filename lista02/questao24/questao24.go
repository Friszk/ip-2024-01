package main

import "fmt"
import "math"

func main() {
  var x, N, somatorio float64
  
  fmt.Scan(&x, &N)
  
  for i := 0.0; i <= N; i++ {
    denominador := 1.0
    for a := 1.0; a <= (2*i); a++ {
      denominador *= a
    }
    somatorio += (math.Pow(-1, i)*math.Pow(x, 2*i))/denominador
  }
  fmt.Printf("cos(%.2f) = %.6f\n", x, somatorio)
}