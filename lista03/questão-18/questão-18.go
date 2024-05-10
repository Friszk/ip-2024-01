package main

import "fmt"

func main() {
  var T int
  fmt.Scan(&T)
  
  CPF := make([]int, 11)
  resultados := []string{}
  
  for i := 0; i < T; i++ {
    for i := range CPF {
      fmt.Scan(&CPF[i])
    }
    
    if verificarDigitoX(CPF) && verificarDigitoY(CPF) {
      resultados = append(resultados, "CPF válido")
    } else {
      resultados = append(resultados, "CPF inválido")
    }
  }
  for i := range resultados {
    fmt.Println(resultados[i])
  }
}
func verificarDigitoX(s[] int) bool {
  soma := 0
  for i := 0; i < 9; i++ {
    soma += s[i] * (i+1)
  }
  if soma % 11 == 10 {
    return s[9] == 0
  }
  return soma % 11 == s[9]
}
func verificarDigitoY(s[] int)bool {
  soma := 0
  a := 9
  for i := 0; i < 9; i++ {
    soma += s[i] * (a)
    a--
  }
  if soma % 11 == 10 {
    return s[10] == 0
  }
  return soma % 11 == s[10]
}