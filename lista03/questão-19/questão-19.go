package main

import "fmt"

func main() {
  var N int
  fmt.Scan(&N)
  if N < 1 || N > 1000 {
    fmt.Println("O número deve estar entre 1 e 1000")
    return
  }
  
  sequencia := make([]int, N)
  sequenciaP := []int{}
  
  for i := range sequencia {
    igual := false
    fmt.Scan(&sequencia[i])
    for a := range sequenciaP {
      if sequencia[i] == sequenciaP[a] {
        igual = true
        continue
      }
    }
    sequenciaP = append(sequenciaP, sequencia[i])
    if !igual {
      fmt.Println(sequencia[i])
    }
  }
}