package main
import "fmt"

func main() {
    var x, y, z, maior int
    var mudou bool
    fmt.Scan(&x, &y, &z)
    
    if x > y && x > z {
        maior = x
    } else if y > x && y > z {
        maior = y
    } else {maior = z}
    
    for i:= 2; i < maior; {
        mudou = false
        
        if x%i == 0 {
            x /= i
            mudou = true
        }
        if y%i == 0 {
            y /= i
            mudou = true
        }
        if z%i == 0 {
            z /= i
            mudou = true
        }
        
        if mudou {
            fmt.Println(i)
        } else {i++}
    }
}
