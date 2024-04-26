package main
import "fmt"

func main() {
    
    var n float64
    fmt.Scan(&n)
    
    for i:= 1.0; i > 0; i++ {
        if n*i == float64(int(n*i)) {
            fmt.Printf("%.f/%.f", n*i, i)
            break
        }
    }
}
