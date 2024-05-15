package main

import "fmt"

func main() {
  var tA, tB int
  var encontrou bool
	for {
	  fmt.Scan(&tA)
	  if tA >= 1 && tA <= 100 {
	    break
	  }
	}
	for {
	  fmt.Scan(&tB)
	  if tB >= 1 && tB <= 100 {
	    break
	  }
	}
	
	A := []int{}
	B := []int{}
	I := []int{}
	
	for len(A) != tA{
	  temp := 0
	  fmt.Scan(&temp)
	  for l := range A {
	    if temp == A[l] {
	      encontrou = true
	    }
	  }
	  if !encontrou {
	    A = append(A, temp)
	  }
	  encontrou = false
	}
	
	for len(B) != tB{
	  temp := 0
	  fmt.Scan(&temp)
	  for l := range B {
	    if temp == B[l] {
	      encontrou = true
	    }
	  }
	  if !encontrou {
	    B = append(B, temp)
	  }
	  encontrou = false
	}
	
	U := A
	for i := range B {
	  encontrou = false
	  for j := range A {
	    if B[i] == A[j] {
	      encontrou = true
	    }
	  }
	  if encontrou {
	    I = append(I, B[i])
	  }
	  if !encontrou {
	    U = append(U, B[i])
	  }
	}
	fmt.Println(U)
	fmt.Println(I)
}