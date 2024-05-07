package main

import (
  "fmt"
  "strconv"
)

func main() {
  var num int
  fmt.Scan(&num)
  numString := strconv.Itoa(num)
  fmt.Println(string(numString[0]))
}