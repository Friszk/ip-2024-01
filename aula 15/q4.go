package main

import (
	"fmt"
	"strconv"
)

func main() {
  var n int
	fmt.Scan(&n)
	fmt.Println(toBinary(n))
}

func toBinary(n int) string {
	if n == 0 {
		return "0"
	}
	if n == 1 {
		return "1"
	}
	return toBinary(n/2) + strconv.Itoa(n%2)
}