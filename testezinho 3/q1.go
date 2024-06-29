package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	resultado := []string{}
	for {
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		linha := strings.TrimSpace(input)
		if linha == "#" {
			break
		}
		palavras := strings.Split(linha, " ")
		for i := range palavras {
			palavras[i] = inverterPalavra(palavras[i])
		}
		inverterPalavras(palavras)
		frase := ""
		for i := range palavras {
			frase += palavras[i]
			frase += " "
		}
		resultado = append(resultado, frase)
	}
	for i := range resultado {
		fmt.Println(resultado[i])
	}

}

func inverterPalavra(s string) string {
	palavraRune := []rune(s)
	for i := 0; i < len(palavraRune)/2; i++ {
		palavraRune[i], palavraRune[len(palavraRune)-1-i] = palavraRune[len(palavraRune)-1-i], palavraRune[i]
	}
	return string(palavraRune)
}

func inverterPalavras(s []string) []string {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return s
}
