package main

import "fmt"

func main() {
  var T int
  anoes := make([]int, 9)
  anoesCertos := []int{}
  somatotal := 0
  
  fmt.Scan(&T)
  
  for i := 0; i < T; i++ {
    for i := range anoes {
      fmt.Scan(&anoes[i])
      somatotal += anoes[i]
    }
    BubbleSort(anoes)
    out:
    for a := range anoes {
      for b := 1; b < len(anoes); b++ {
        if somatotal - (anoes[a]+anoes[b]) == 100 && anoes[a] != anoes[b]{
          anoesCertos = append(anoesCertos, append(anoes[:a], append(anoes[a+1:b], anoes[b+1:]...)...)...)
          somatotal = 0
          break out
        }
      }
    }
  }
  for i := range anoesCertos {
    fmt.Println(anoesCertos[i])
  }
}

func BubbleSort(s []int)[]int {
  for {
    trocou := false
    for i := 0; i < len(s)-1; i++ {
      if s[i] > s[i+1] {
        s[i], s[i+1] = s[i+1], s[i]
        trocou = true
      }
    }
    if !trocou {
      break
    }
  }
  return s
}