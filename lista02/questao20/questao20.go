package main

import "fmt"

func main() {
  var n, soma int
  divisores := []int{}
  fmt.Println("Digite um número para descobrir se ele é perfeito.")
  fmt.Scan(&n)
  
  for i := 1; i < n; i++ {
    if n % i == 0 {
      divisores = append(divisores, i)
    }
  }
  
  for i := 0; i < len(divisores); i++ {
    soma += divisores[i]
  }
  
  fmt.Printf("%v = ", n)
  for i, num := range divisores {=
    fmt.Print(num)
    if i < len(divisores)-1 {
      fmt.Print(" + ")
    }
  }
  if soma == n {
  fmt.Printf(" = %v (NÚMERO PERFEITO)\n", soma)
  } else {
    fmt.Printf(" = %v (NÚMERO NÃO É PERFEITO)\n", soma)
  }
}