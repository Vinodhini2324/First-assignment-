# First-assignment-
package main
import "fmt"

func plus(A int, B int) int {

    return A + B
}

func plusPlus(A, B, C int) int {
    return A + B + C
}

func main() {


    res := plus(2, 2)
    fmt.Println("2+2 =", res)

    res = plusPlus(2, 2, 1)
    fmt.Println("2+2+1 =", res)
}
