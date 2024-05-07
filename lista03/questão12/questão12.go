package main

import "fmt"

func main() {
  var n int
  
  fmt.Scan(&n)
  lista := make([]int, n)
  
  for i := range lista {
    fmt.Scan(&lista[i])
  }
  
  for {
    trocou := false
    for i := 0; i < n - 1; i++ {
      if lista[i] > lista[i+1]{
        lista[i], lista[i+1] = lista[i+1], lista[i]
        trocou = true
      }
    }
    if !trocou {
      break
    }
  }
  fmt.Println()
  for i := range lista {
    fmt.Println(lista[i])
  }
}