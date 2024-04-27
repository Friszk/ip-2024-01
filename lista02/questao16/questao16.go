package main

import "fmt"

func main() {
  var n float64
  
  fmt.Scan(&n)
  
  for i := 1.0; i <= n; i += 0.1 {
    numH1 := 5*i
    numH2 := float64(int(5*i))
    varDeH1H2 := numH1 - numH2
    numCC1 := 4*i
    numCC2 := float64(int(4*i))
    varDeCC1CC2 := numCC1 - numCC2
    numC1 := 3*i
    numC2 := float64(int(3*i))
    varDeC1C2 := numC1 - numC2
    
    fmt.Println(varDeH1H2)
   fmt.Println(varDeCC1CC2)
   fmt.Println(varDeC1C2)
    fmt.Println(" ")
    
    
    
    
    if (4*i)*(4*i) + (3*i)*(3*i) == (5*i)*(5*i) && varDeCC1CC2 < 0.01 && varDeC1C2 < 0.001 && varDeH1H2 < 0.001{
      fmt.Printf("hipotenusa = %v, catetos %v e %v\n", int(5*i), int(3*i), int(4*i))
    }
  }
}