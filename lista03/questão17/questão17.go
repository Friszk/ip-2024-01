package main

import "fmt"

func main() {
  var n, unicos int
  fmt.Scan(&n)
  
  if n > 5000 {
    fmt.Println("O número deve ser menor que 5000")
    return
  }
  
  lista := make([]int, n)
  
  for i := range lista {
    fmt.Scan(&lista[i])
  }
  
  for _, v := range lista {
    cont := 0
    for _, j := range lista {
      if v == j {
        cont++
      }
    }
    if cont == 1 {
      unicos ++
    }
    cont = 0
  }
  fmt.Println(unicos)
}