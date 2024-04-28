package main

import "fmt"

func main() {
  var n, quantN, soma1, soma2 int
  fmt.Println("Digite a quantidade de números amigos a serem mostrados (max. 9).")
  
  for {
    fmt.Scan(&n)
    if n <= 9 && n >= 1{
    break
    } else {
      fmt.Println("O número deve ser menor ou igual a 9 e maior que 0.")
    }
  }
  
  for i := 2; ; i++ {
    for x := 1; x < i; x++ {
      if x == i && x == 1 {
        break
      }
      if i % x == 0 {
        soma1 += x
      }
    }
    for y := 1; y < soma1; y++ {
      if soma1 % y == 0 && soma1 != i {
        soma2 += y
      }
    }
    if soma2 == i && soma1 > i {
      quantN++
      fmt.Println(i, soma1)
    } else {
      soma1 = 0
      soma2 = 0
    }
    if quantN == n {
      break
    }
  }
}