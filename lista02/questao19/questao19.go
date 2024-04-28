package main

import "fmt"

func main() {
	var n, soma, i int
	impares := []int{}
	somaslice := []int{}
	
	fmt.Println("Digite o número a ser mostrada sua soma em ímpares igual ao seu cubo")
	fmt.Scan(&n)

	for i := 0; i < n*n*n; i++ {
		if i%2 != 0 {
			impares = append(impares, i)
		}
	}


for n2 := 1; n2 < n+1; n2++ {	
	for i = 0; ; i++{
	  for a := i; a < n2+i; a++ {
	    somaslice = append(somaslice, impares[a])
	    soma += impares[a]
	  }
	  if soma == n2*n2*n2{
	    fmt.Printf("%v*%v*%v = ", n2, n2, n2)
	    for i, num := range somaslice {
	      fmt.Print(num)
	      if i < len(somaslice)-1 {
	        fmt.Print(" + ")
	      }
	    }
	    fmt.Println("")
	    break
	  } else {
	    somaslice = nil
	    soma = 0
	  }
}
}
}