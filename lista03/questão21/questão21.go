package main

import "fmt"

func main() {
  var N int
  
  fmt.Scan(&N)
  listaP := []int{}
  listaI := []int{}
  lista := make([]int, N)
  
  for i := range lista {
    fmt.Scan(&lista[i])
  }
  for {
    trocou := false
    for i := 0; i < N - 1; i++ {
      if lista[i] > lista [i+1] {
        lista[i], lista[i+1] = lista[i+1], lista[i]
        trocou = true
      }
    }
    if !trocou {
      break
    }
  }
    
    for i := range lista {
      if lista[i] % 2 == 0 {
        listaP = append(listaP, lista[i])
      } else {
        listaI = append(listaI, lista[i])
      }
    }
    for i := range listaP {
      fmt.Println(listaP[i])
    }
    for i := len(listaI) - 1; i >= 0; i-- {
      fmt.Println(listaI[i])
    }
}