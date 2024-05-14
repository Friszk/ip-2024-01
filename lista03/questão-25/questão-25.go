package main

import (
  "fmt"
)

func main() {
  var acertos, sena, quina, quadra int
  sorteado := make([]int, 6)
  for i := range sorteado {
    fmt.Scan(&sorteado[i])
  }

  n := 0
  fmt.Scan(&n)
  for i := 0; i < n; i++ {
    aposta := make([]int, 6)
    for i := range aposta {
      fmt.Scan(&aposta[i])
    }
    for i := range aposta {
      for j := range sorteado {
        if aposta[i] == sorteado[j] {
          acertos++
          break
        }
      }
    }
    switch {
      case acertos == 6: sena++
      case acertos == 5: quina++
      case acertos == 4: quadra++
    }
    acertos = 0
  }
  if sena == 0 {
    fmt.Println("Não houve acertador para sena")
  } else {
    fmt.Printf("Houve %v acertadores para sena\n", sena)
  }
  if quina == 0 {
    fmt.Println("Não houve acertador para quina")
  } else {
    fmt.Printf("Houve %v acertadores para quina\n", quina)
  }
  if quadra == 0 {
    fmt.Println("Não houve acertador para quadra")
  } else {
    fmt.Printf("Houve %v acertadores para quadra\n", quadra)
  }
}