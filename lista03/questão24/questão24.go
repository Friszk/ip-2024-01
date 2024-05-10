package main

import "fmt"

func main() {
  var N int
  resultados := [][]int{}
  for {
    fmt.Scan(&N)
    if N == 0 {
      break
    }
    v := make([]int, N)
    for i := range v {
      fmt.Scan(&v[i])
    }
    resultados = append(resultados, (CountingSort(v)))
  }
  for i := range resultados {
    fmt.Println(resultados[i])
  }
}

func CountingSort(s []int) []int {
  M := 0
  N := len(s)
  vOrd := make([]int, N)
  for i := range s {
    if s[i] > M {
      M = s[i]
    }
  }
  vCount := make([]int, M+1)
  for i := range s {
    vCount[s[i]]++
  }
  for i := 1; i <= M; i++ {
    vCount[i] += vCount[i-1]
  }
  for i := range s {
    vOrd[vCount[s[i]]-1] = s[i]
    vCount[s[i]]--
  }
  return vOrd
}