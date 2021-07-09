# First-assignment-
import "fmt"
func main(){
    i := 1
    for i <= 2 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 6; j <= 8; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
