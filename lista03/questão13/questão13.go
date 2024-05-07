package main

import "fmt"

func main() {
  var n int
  
  fmt.Scan(&n)
  lista := make([]int, n)
  
  for i := range lista {
    fmt.Scan(&lista[i])
  }
  
  maiorEnc := 0
  numMaiorEnc := 0
  
  for i := range lista {
  maiorAtual := 0
  for c := range lista {
    if lista[i] == lista[c] {
      maiorAtual++
      }
    }
    if maiorAtual > maiorEnc {
      maiorEnc = maiorAtual
      numMaiorEnc = lista[i]
    } else if maiorAtual == maiorEnc {
      if lista[i] < numMaiorEnc {
        numMaiorEnc = lista[i]
      }
    }
  }
  
  fmt.Printf("%v\n%v\n", numMaiorEnc, maiorEnc)
  
}