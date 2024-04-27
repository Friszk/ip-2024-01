package main
import "fmt"

func main() {
    
    var n float64
    fmt.Println("Digite o número real para descobrir sua fração originária")
    fmt.Scan(&n)
    
    for i:= 1.0; i > 0; i++ {
      
      num1 := n*i
      num2 := float64(int(n*i))
      erro := num1 - num2
      
        if erro < 0.000000001 {
            fmt.Printf("%.f/%.f", n*i, i)
            break
        }
    }
}
