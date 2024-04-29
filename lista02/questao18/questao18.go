package main
import "fmt"

func main() {
    var x, y, z, maior, MMC int
    var mudoux, mudouy, mudouz bool
    fmt.Scan(&x, &y, &z)
    
    if x > y && x > z {
        maior = x
    } else if y > x && y > z {
        maior = y
    } else {maior = z}
    
    MMC = 1
    
    for i:= 2; i <= maior; {
        mudoux = false
        mudouy = false
        mudouz = false
        
        if x%i == 0 {
            mudoux = true
        }
        if y%i == 0 {
            mudouy = true
        }
        if z%i == 0 {
            mudouz = true
        }
        if mudoux || mudouy || mudouz{
            fmt.Printf("%v %v %v : %v\n", x, y, z, i)
            MMC *= i
        } else {i++}
        if mudoux {
            x /= i
        }
        if mudouy {
            y /= i
        }
        if mudouz {
            z /= i
        }
    }
    fmt.Println("MMC:", MMC)
}
