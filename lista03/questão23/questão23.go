package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "math"
)

func main() {
  var soma float64
  vogais1:= map[rune]float64{'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0}
  vogais2:= map[rune]float64{'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0}
  input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
  frase := strings.Split(input, ";")
  if len(frase) <= 1 {
    fmt.Println("FORMATO INVÁLIDO")
    return
  }
  for _, char := range strings.ToLower(strings.TrimSpace(frase[0])) {
    if _, ok := vogais1[char]; ok {
      vogais1[char]++
    }
  }
  for _, char := range strings.ToLower(strings.TrimSpace(frase[1])) {
    if _, ok := vogais2[char]; ok {
      vogais2[char]++
    }
  }
  soma = math.Sqrt(math.Pow((vogais1['a']-vogais2['a']), 2) + math.Pow((vogais1['e']-vogais2['e']), 2) + math.Pow((vogais1['i']-vogais2['i']), 2) + math.Pow((vogais1['o']-vogais2['o']), 2) + math.Pow((vogais1['u']-vogais2['u']), 2))
  fmt.Printf("(%v, %v, %v, %v, %v)\n", vogais1['a'], vogais1['e'], vogais1['i'], vogais1['o'], vogais1['u'])
  fmt.Printf("(%v, %v, %v, %v, %v)\n", vogais2['a'], vogais2['e'], vogais2['i'], vogais2['o'], vogais2['u'])
  fmt.Printf("%.2f\n", soma)
}
